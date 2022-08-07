package models

import "time"

type Overview struct {
	MonthsOverview []MonthOverview `json:"monthsOverview"`
}

type MonthOverview struct {
	Date  time.Time `json:"date" bson:"date"`
	Total float64   `json:"total" bson:"total"`
}
