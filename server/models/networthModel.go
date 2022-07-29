package models

import (
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Diversificationable interface {
	GetValue() float32
	GetAssets() []Diversification
}

type NetWorth struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	UserId  primitive.ObjectID `json:"userId" bson:"userId" binding:"required"`
	Records []Record           `json:"records" bson:"records"`
}

type Info struct {
	Id                    primitive.ObjectID `json:"id" bson:"_id"`
	InvestedAmount        float32            `json:"investedAmount,omitempty" bson:"investedAmount"`
	Stocks                Stocks             `json:"stocks" bson:"stocks" binding:"required,dive"`
	Cryptos               Cryptos            `json:"cryptos" bson:"cryptos" binding:"required"`
	StockDiversification  []Diversification  `json:"stockDiversification" `
	CryptoDiversification []Diversification  `json:"cryptoDiversification"`
	CryptosValue          float32            `json:"cryptosValue,omitempty" bson:"cryptosValue"`
	StocksValue           float32            `json:"stocksValue,omitempty" bson:"stocksValue"`
	Liquidity             *float32           `json:"liquidity" bson:"liquidity" binding:"required,min=0"`
}

type Record struct {
	Id     primitive.ObjectID `json:"id" bson:"id"`
	Date   time.Time          `json:"date" bson:"date" binding:"required"`
	InfoId primitive.ObjectID `json:"infoId" bson:"infoId" binding:"required"`
}

type Diversification struct {
	Symbol  string  `json:"symbol"`
	Percent float32 `json:"percent"`
}

type Stock struct {
	Name     string   `json:"name" bson:"name" binding:"required"`
	Symbol   string   `json:"symbol" bson:"symbol" binding:"required"`
	Shares   float32  `json:"shares" bson:"shares" binding:"required,min=0"`
	ValuedAt *float32 `json:"valuedAt" bson:"valuedAt" binding:"required,min=0"`
}

type Crypto struct {
	Name     string   `json:"name" bson:"name" binding:"required"`
	Symbol   string   `json:"symbol" bson:"symbol" binding:"required"`
	Coins    float32  `json:"coins" bson:"coins" binding:"required,min=0"`
	ValuedAt *float32 `json:"valuedAt" bson:"valuedAt" binding:"required,min=0"`
}

type Stocks []Stock
type Cryptos []Crypto

type DeleteRecordBody struct {
	Id string `json:"id" binding:"required"`
}

type RecordBody struct {
	Id                    string            `json:"id,omitempty"`
	Date                  time.Time         `json:"date" binding:"required"`
	InvestedAmount        float32           `json:"investedAmount,omitempty"`
	Stocks                Stocks            `json:"stocks" binding:"required,dive"`
	Cryptos               Cryptos           `json:"cryptos" binding:"required"`
	StockDiversification  []Diversification `json:"stockDiversification" `
	CryptoDiversification []Diversification `json:"cryptoDiversification"`
	CryptosValue          float32           `json:"cryptosValue,omitempty"`
	StocksValue           float32           `json:"stocksValue,omitempty"`
	Liquidity             *float32          `json:"liquidity" binding:"required,min=0"`
}

func (stocks Stocks) GetValue() (sum float32) {
	for _, stock := range stocks {
		sum += *stock.ValuedAt * stock.Shares
	}
	return
}
func (cryptos Cryptos) GetValue() (sum float32) {
	for _, crypto := range cryptos {
		sum += *crypto.ValuedAt * crypto.Coins
	}
	return
}
func (stock Stock) GetValue() (sum float32) {
	return *stock.ValuedAt * stock.Shares
}
func (crypto Crypto) GetValue() (sum float32) {
	return *crypto.ValuedAt * crypto.Coins
}

func (info Info) GetStockDiversification(stocksValue float32) (diversification []Diversification) {
	diversification = make([]Diversification, 0, len(info.Stocks))
	for _, stock := range info.Stocks {
		div := Diversification{stock.Symbol, stock.GetValue() / stocksValue}
		diversification = append(diversification, div)
	}
	return
}

func (info Info) GetCryptoDiversification(cryptosValue float32) (diversification []Diversification) {
	diversification = make([]Diversification, 0, len(info.Cryptos))
	for _, crypto := range info.Cryptos {
		div := Diversification{crypto.Symbol, crypto.GetValue() / cryptosValue}
		diversification = append(diversification, div)
	}
	return
}

func (info *Info) GenerateStatistics() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		info.StocksValue = info.Stocks.GetValue()
		info.StockDiversification = info.GetStockDiversification(info.StocksValue)
		wg.Done()
	}()
	go func() {
		info.CryptosValue = info.Cryptos.GetValue()
		info.CryptoDiversification = info.GetCryptoDiversification(info.CryptosValue)
		wg.Done()
	}()
	wg.Wait()
}

func (recordBody RecordBody) Split() (record Record, info Info) {
	info.Id = primitive.NewObjectID()
	record.Date = recordBody.Date
	record.Id = primitive.NewObjectID()
	record.InfoId = info.Id
	info.InvestedAmount = recordBody.InvestedAmount
	info.Stocks = recordBody.Stocks
	info.Cryptos = recordBody.Cryptos
	info.GenerateStatistics()
	info.Liquidity = recordBody.Liquidity
	return
}

func ConcatRecord(record Record, info Info) RecordBody {
	recordBody := RecordBody{
		Id:                    record.Id.Hex(),
		Date:                  record.Date,
		InvestedAmount:        info.InvestedAmount,
		Stocks:                info.Stocks,
		Cryptos:               info.Cryptos,
		StockDiversification:  info.StockDiversification,
		CryptoDiversification: info.CryptoDiversification,
		CryptosValue:          info.CryptosValue,
		StocksValue:           info.StocksValue,
		Liquidity:             info.Liquidity,
	}
	return recordBody
}