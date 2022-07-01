package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Diversificationable interface {
	GetValue() float32
	GetAssets() []Diversification
}

type NetWorth struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	UserId  primitive.ObjectID `json:"userId" bson:"userId" validate:"required"`
	Records []Record           `json:"records" bson:"records"`
}

type Record struct {
	Id                    primitive.ObjectID `json:"id" bson:"id"`
	Date                  time.Time          `json:"date" bson:"date" validate:"required"`
	CurrentMoney          float32            `json:"currentMoney" bson:"currentMoney" validate:"required"`
	Stocks                Stocks             `json:"stocks" bson:"stocks"`
	Cryptos               Cryptos            `json:"cryptos" bson:"cryptos"`
	StockDiversification  []Diversification  `json:"stockDiversification"`
	CryptoDiversification []Diversification  `json:"cryptoDiversification"`
	CryptoWorth           float32            `json:"cryptoWorth,omitempty"`
	StockWorth            float32            `json:"stockWorth,omitempty"`
	Liquidity             float32            `json:"liquidity" bson:"liquidity" validate:"min=0"`
}

type Diversification struct {
	Symbol  string  `json:"symbol"`
	Percent float32 `json:"percent"`
}

type Stock struct {
	Name     string  `json:"name" bson:"name"`
	Symbol   string  `json:"symbol" bson:"symbol"`
	Shares   float32 `json:"shares" bson:"shares"`
	ValuedAt float32 `json:"valuedAt" bson:"valuedAt"`
}

type Crypto struct {
	Name     string  `json:"name" bson:"name"`
	Symbol   string  `json:"symbol" bson:"symbol"`
	Coins    float32 `json:"coins" bson:"coins"`
	ValuedAt float32 `json:"valuedAt" bson:"valuedAt"`
}

type Stocks []Stock
type Cryptos []Crypto

func (stocks Stocks) GetStocksValue() (sum float32) {
	for _, stock := range stocks {
		sum += stock.ValuedAt
	}
	return
}
func (cryptos Cryptos) GetStocksValue() (sum float32) {
	for _, stock := range cryptos {
		sum += stock.ValuedAt
	}
	return
}

func (record Record) GetStockDiversification() (diversification []Diversification) {
	diversification = make([]Diversification, len(record.Stocks))
	stocksWorth := record.Cryptos.GetStocksValue()
	for _, stock := range record.Stocks {
		div := Diversification{stock.Symbol, stock.ValuedAt / stocksWorth}
		diversification = append(diversification, div)
	}
	return
}

func (record Record) GetCryptoDiversification() (diversification []Diversification) {
	diversification = make([]Diversification, len(record.Stocks))
	cyptoWorth := record.Cryptos.GetStocksValue()
	for _, crypto := range record.Stocks {
		div := Diversification{crypto.Symbol, crypto.ValuedAt / cyptoWorth}
		diversification = append(diversification, div)
	}
	return
}
