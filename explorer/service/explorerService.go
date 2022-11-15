package service

import (
	"context"

	"github.com/TudorEsan/FinanceAppGo/explorer/blockchains"
	explorer "github.com/TudorEsan/FinanceAppGo/explorer/proto"
	"github.com/hashicorp/go-hclog"
)

type AddressExplorerServer struct {
	l hclog.Logger
	explorer.UnimplementedAddressExplorerServer
}

type IAddressExplorerServer interface {
	explorer.AddressExplorerServer
}

func NewAddressExplorerServer(l hclog.Logger) IAddressExplorerServer{
	return &AddressExplorerServer{l, explorer.UnimplementedAddressExplorerServer{}}
}

func (s *AddressExplorerServer) GetBtcBallance(ctx context.Context, address *explorer.Address) (*explorer.AddressOverview, error) {
	s.l.Info("GetBtcBallance")
	res, err := blockchains.GetBtcFromAddress(address.Address)
	if err != nil {
		return nil, err
	}

	return &explorer.AddressOverview{
		Blockchain: "BTC",
		Balance: float32(res.Amount),
		UsdBallance: float32(res.USDAmount),
	}, nil

}

func (s *AddressExplorerServer) GetEthBallance(ctx context.Context, address *explorer.Address) (*explorer.AddressOverview, error) {
	s.l.Info("GetEthBallance")
	res, err := blockchains.GetEthFromAddress(address.Address)
	if err != nil {
		return nil, err
	}

	return &explorer.AddressOverview{
		Blockchain: "ETH",
		Balance: float32(res.Amount),
		UsdBallance: float32(res.USDAmount),
	}, nil

}

