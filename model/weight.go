package model

import "time"

type WeightEntry struct {
	Id            int       `json:"id"`
	UserAccountId int       `json:"userId"`
	Weight        int       `json:"weight"`
	EntryDate     time.Time `json:"entryDate"`
	Category      string    `json:"category"`
}
