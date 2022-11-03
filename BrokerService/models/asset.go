package models

import "fmt"


type Asset struct {
	Name   string
	Amount float64
	Price  float64
	Worth  float64
}

func (a Asset) String() string {
	return fmt.Sprintf("%s: %f (%f)", a.Name, a.Amount, a.Amount*a.Price)
}