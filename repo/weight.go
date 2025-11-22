package repo

import (
	"weight-tracker/model"
)

type WeightRepo struct {
	Db []model.WeightEntry
}

func NewWeightRepo() *WeightRepo {
	// in memory-db for now
	return &WeightRepo{
		Db: []model.WeightEntry{
			{
				Id:        1,
				UserId:    1,
				Weight:    198.23,
				EntryDate: "2025-11-01",
				Category:  "daily",
			},
			{
				Id:        2,
				UserId:    1,
				Weight:    195.5,
				EntryDate: "2025-11-02",
				Category:  "daily",
			},
			{
				Id:        3,
				UserId:    1,
				Weight:    196.7,
				EntryDate: "2025-11-03",
				Category:  "daily",
			},
			{
				Id:        4,
				UserId:    1,
				Weight:    192.1,
				EntryDate: "2025-11-04",
				Category:  "daily",
			},
			{
				Id:        5,
				UserId:    1,
				Weight:    189.75,
				EntryDate: "2025-11-05",
				Category:  "daily",
			},
		},
	}
}

func (wr *WeightRepo) GetDailyWeightEntriesForUser(userId int) ([]model.WeightEntry, error) {
	// if there is a DB error, return that error
	weightEntries := []model.WeightEntry{}
	for _, we := range wr.Db {
		if we.UserId == userId && we.Category == "daily" {
			weightEntries = append(weightEntries, we)
		}
	}
	return weightEntries, nil
}
