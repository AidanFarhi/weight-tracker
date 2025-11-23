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
				Category:  CategoryDaily,
			},
			{
				Id:        2,
				UserId:    1,
				Weight:    195,
				EntryDate: "2025-11-02",
				Category:  CategoryDaily,
			},
			{
				Id:        3,
				UserId:    1,
				Weight:    196,
				EntryDate: "2025-11-03",
				Category:  CategoryDaily,
			},
			{
				Id:        4,
				UserId:    1,
				Weight:    192,
				EntryDate: "2025-11-04",
				Category:  CategoryDaily,
			},
			{
				Id:        5,
				UserId:    1,
				Weight:    189,
				EntryDate: "2025-11-05",
				Category:  CategoryDaily,
			},
			{
				Id:        6,
				UserId:    1,
				Weight:    175,
				EntryDate: "2025-10-01",
				Category:  CategoryTarget,
			},
		},
	}
}

func (wr *WeightRepo) GetDailyWeightEntriesForUser(userId int) ([]model.WeightEntry, error) {
	// if there is a DB error, return that error
	weightEntries := []model.WeightEntry{}
	for _, we := range wr.Db {
		if we.UserId == userId && we.Category == CategoryDaily {
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
		if we.EntryDate > maxDate && we.Category == CategoryDaily {
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
		if we.EntryDate > maxDate && we.Category == CategoryTarget {
			latestTargetWeight = we
		}
	}
	return latestTargetWeight, nil
}

func (wr *WeightRepo) CreateWeightEntry(weightEntry model.WeightEntry) error {
	weightEntry.Id = 1 // hardcode as one for now
	wr.Db = append(wr.Db, weightEntry)
	return nil
}
