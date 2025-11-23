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
				Weight:    198,
				EntryDate: "2025-11-01",
				Category:  "daily",
			},
			{
				Id:        2,
				UserId:    1,
				Weight:    195,
				EntryDate: "2025-11-02",
				Category:  "daily",
			},
			{
				Id:        3,
				UserId:    1,
				Weight:    196,
				EntryDate: "2025-11-03",
				Category:  "daily",
			},
			{
				Id:        4,
				UserId:    1,
				Weight:    192,
				EntryDate: "2025-11-04",
				Category:  "daily",
			},
			{
				Id:        5,
				UserId:    1,
				Weight:    189,
				EntryDate: "2025-11-05",
				Category:  "daily",
			},
			{
				Id:        6,
				UserId:    1,
				Weight:    175,
				EntryDate: "2025-10-01",
				Category:  "target",
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

func (wr *WeightRepo) GetLatestDailyWeightForUser(userId int) (model.WeightEntry, error) {
	// if there is a DB error, return that error
	var latestWeightEntry model.WeightEntry
	maxDate := "1900-01-01"
	for _, we := range wr.Db {
		if we.EntryDate > maxDate && we.Category == "daily" {
			latestWeightEntry = we
		}
	}
	return latestWeightEntry, nil
}

func (wr *WeightRepo) GetLatestTargetWeightForUser(userId int) (model.WeightEntry, error) {
	// if there is a DB error, return that error
	var latestTargetWeight model.WeightEntry
	maxDate := "1900-01-01"
	for _, we := range wr.Db {
		if we.EntryDate > maxDate && we.Category == "target" {
			latestTargetWeight = we
		}
	}
	return latestTargetWeight, nil
}
