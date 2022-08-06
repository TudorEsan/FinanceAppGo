package models

type Overview struct {
	MonthsOverview []MonthOverview `json:"monthsOverview"`
}

type MonthOverview struct {
	Month int
	Total float64
}
