package service

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/hashicorp/go-hclog"
)

type IBinanceService interface {
	GetWalletAssets() ([]Asset, error)
	GetPrice(ticker string) (price float64, err error)
	GetStakingAssets() (assets map[string]*Asset, err error)
	GetAssets() ([]Asset, error)
}

type BinanceService struct {
	client *binance.Client
	l      hclog.Logger
}

func NewBinanceService(apiKey, secretKey string) IBinanceService {
	binanceClient := binance.NewClient(apiKey, secretKey)
	return &BinanceService{
		client: binanceClient,
		l:      hclog.L().Named("BinanceService"),
	}
}

type Asset struct {
	Name   string
	Amount float64
	Price  float64
	Worth  float64
}

func (a Asset) String() string {
	return fmt.Sprintf("%s: %f (%f)", a.Name, a.Amount, a.Amount*a.Price)
}

func (s *BinanceService) GetPrice(ticker string) (price float64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tickerPrice, err := s.client.NewListPricesService().Symbol(ticker).Do(ctx)
	if err != nil {
		return 0, err
	}
	price, err = strconv.ParseFloat(tickerPrice[0].Price, 64)
	return
}

func (s *Asset) AddPrice(wg *sync.WaitGroup) {

}

func (s *BinanceService) GetStakingAssets() (assets map[string]*Asset, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	staking, err := s.client.NewStakingProductPositionService().Product("STAKING").Do(ctx)
	if err != nil {
		panic(err)
	}
	assets = make(map[string]*Asset)
	for _, asset := range *staking {
		assetAmount, err := strconv.ParseFloat(asset.Amount, 64)
		if err != nil {
			s.l.Error("Error parsing asset amount", "error", err)
			return nil, err
		}
		if _, ok := assets[asset.Asset]; ok {
			if err != nil {
				s.l.Error("Error parsing asset amount", "error", err, "asset", asset.Asset)
				return nil, err
			}
			s.l.Info("Adding to existing asset", "asset", asset.Asset, "amount", assetAmount)
			assets[asset.Asset].Amount += assetAmount
		} else {
			s.l.Info("Adding new asset", "asset", asset.Asset, "amount", assetAmount)
			assets[asset.Asset] = &Asset{
				Name:   asset.Asset,
				Amount: assetAmount,
				Price:  0,
			}
		}
	}
	for _, asset := range assets {
		asset.Price, err = s.GetPrice(asset.Name + "USDT")
		if err != nil {
			return nil, err
		}
	}
	s.l.Info("Staking assets", "assets", assets)

	return
}

func (s *BinanceService) GetWalletAssets() ([]Asset, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	snapshot, err := s.client.NewGetAccountSnapshotService().Limit(1).Type("SPOT").Do(ctx)
	if err != nil {
		s.l.Error("Error getting account", "error", err)
		return nil, err
	}
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	assets := make([]Asset, 0)
	snp := snapshot.Snapshot[0]
	for _, balance := range snp.Data.Balances {
		wg.Add(1)
		go func(balance *binance.SnapshotBalances) {
			free, err := strconv.ParseFloat(balance.Free, 64)
			if err != nil {
				s.l.Error("Error parsing free balance", "error", err)
				wg.Done()
				return
			}

			locked, err := strconv.ParseFloat(balance.Locked, 64)
			if err != nil {
				s.l.Error("Error parsing locked balance", "error", err)
				wg.Done()
				return
			}

			if free+locked > 0 {
				s.l.Info("Getting price for", "ticker", balance.Asset, "amount", balance.Free)
				price, err := s.GetPrice(balance.Asset + "USDT")
				if err != nil {
					s.l.Error("Error getting price", "error", err, "ticker", balance.Asset)
					wg.Done()
					return
				}
				mutex.Lock()
				assets = append(assets, Asset{
					Name:   balance.Asset,
					Amount: free + locked,
					Price:  price,
				})
				mutex.Unlock()
			}
			wg.Done()
		}(balance)
	}
	wg.Wait()
	s.l.Info("Wallet assets", "assets", assets)
	return assets, nil
}

func (s *BinanceService) GetAssets() ([]Asset, error) {
	walletAssets, err := s.GetWalletAssets()
	if err != nil {
		return nil, err
	}
	stakingAssets, err := s.GetStakingAssets()
	if err != nil {
		return nil, err
	}
	assets := make([]Asset, 0, len(walletAssets)+len(stakingAssets))
	for _, asset := range walletAssets {
		if a, ok := stakingAssets[asset.Name]; ok {
			asset.Amount += a.Amount
			assets = append(assets, Asset{
				Name:   asset.Name,
				Amount: asset.Amount + a.Amount,
				Price:  asset.Price,
				Worth:  (asset.Amount + a.Amount) * asset.Price,
			})
		} else {
			assets = append(assets, Asset{
				Name:   asset.Name,
				Amount: asset.Amount,
				Price:  asset.Price,
				Worth:  asset.Amount * asset.Price,
			})
		}
	}
	return assets, nil
}
