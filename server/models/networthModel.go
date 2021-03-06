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

type Record struct {
	Id                    primitive.ObjectID `json:"id" bson:"id"`
	Date                  time.Time          `json:"date" bson:"date" binding:"required"`
	InvestedAmount        float32            `json:"investedAmount,omitempty" bson:"investedAmount"`
	Stocks                Stocks             `json:"stocks" bson:"stocks" binding:"required,dive"`
	Cryptos               Cryptos            `json:"cryptos" bson:"cryptos" binding:"required"`
	StockDiversification  []Diversification  `json:"stockDiversification" `
	CryptoDiversification []Diversification  `json:"cryptoDiversification"`
	CryptosValue          float32            `json:"cryptosValue,omitempty" bson:"cryptosValue"`
	StocksValue           float32            `json:"stocksValue,omitempty" bson:"stocksValue"`
	Liquidity             *float32           `json:"liquidity" bson:"liquidity" binding:"required,min=0"`
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

func (record Record) GetStockDiversification(stocksValue float32) (diversification []Diversification) {
	diversification = make([]Diversification, 0, len(record.Stocks))
	for _, stock := range record.Stocks {
		div := Diversification{stock.Symbol, stock.GetValue() / stocksValue}
		diversification = append(diversification, div)
	}
	return
}

func (record Record) GetCryptoDiversification(cryptosValue float32) (diversification []Diversification) {
	diversification = make([]Diversification, 0, len(record.Cryptos))
	for _, crypto := range record.Cryptos {
		div := Diversification{crypto.Symbol, crypto.GetValue() / cryptosValue}
		diversification = append(diversification, div)
	}
	return
}

func (record *Record) GenerateStatistics() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		record.StocksValue = record.Stocks.GetValue()
		record.StockDiversification = record.GetStockDiversification(record.StocksValue)
		wg.Done()
	}()
	go func() {
		record.CryptosValue = record.Cryptos.GetValue()
		record.CryptoDiversification = record.GetCryptoDiversification(record.CryptosValue)
		wg.Done()
	}()
	wg.Wait()
}
