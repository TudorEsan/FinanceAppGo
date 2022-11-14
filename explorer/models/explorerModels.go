package models

type AddressOverview struct {
	Blockchain string `json:"blockchain"`
	Amount float64 `json:"amount"`
	USDAmount float64 `json:"usdAmount"`
}