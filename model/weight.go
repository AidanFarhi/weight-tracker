package model

type WeightEntry struct {
	Id        int     `json:"id"`
	UserId    int     `json:"userId"`
	Weight    float64 `json:"weight"`
	EntryDate string  `json:"entryDate"`
	Category  string  `json:"category"`
}
