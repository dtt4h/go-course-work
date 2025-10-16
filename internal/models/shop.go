package models

type Shop struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	Specialization string `json:"specialization"`
}

type ShopSummary struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Specialization string `json:"specialization"`
}
