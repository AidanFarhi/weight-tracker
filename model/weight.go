package model

type WeightEntry struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Weight    int    `json:"weight"`
	EntryDate string `json:"entryDate"`
	Category  string `json:"category"`
}
