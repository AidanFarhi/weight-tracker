package model

type WeightEntry struct {
	Id            int    `json:"id"`
	UserAccountId int    `json:"userId"`
	Weight        int    `json:"weight"`
	EntryDate     string `json:"entryDate"`
	Category      string `json:"category"`
}
