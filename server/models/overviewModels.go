package models

import "time"

type Overview struct {
	NetworthOverview  []NetworthOverview  `json:"networthOverview"`
	LiquidityOverview []LiquidityOverview `json:"liquidityOverview"`
	CurrentRecord     Record              `json:"currentRecord"`
	LastRecord        Record              `json:"lastRecord"`
}

type NetworthOverview struct {
	Date  time.Time `json:"date" bson:"date"`
	Total float64   `json:"total" bson:"total"`
}

type LiquidityOverview struct {
	Date      time.Time `json:"date" bson:"date"`
	Liquidity float64   `json:"liquidity" bson:"liquidity"`
}
