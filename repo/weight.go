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
				Category:  CATEGORY_DAILY,
			},
			{
				Id:        2,
				UserId:    1,
				Weight:    195,
				EntryDate: "2025-11-02",
				Category:  CATEGORY_DAILY,
			},
			{
				Id:        3,
				UserId:    1,
				Weight:    196,
				EntryDate: "2025-11-03",
				Category:  CATEGORY_DAILY,
			},
			{
				Id:        4,
				UserId:    1,
				Weight:    192,
				EntryDate: "2025-11-04",
				Category:  CATEGORY_DAILY,
			},
			{
				Id:        5,
				UserId:    1,
				Weight:    189,
				EntryDate: "2025-11-05",
				Category:  CATEGORY_DAILY,
			},
			{
				Id:        6,
				UserId:    1,
				Weight:    175,
				EntryDate: "2025-10-01",
				Category:  CATEGORY_TARGET,
			},
		},
	}
}

func (wr *WeightRepo) GetDailyWeightEntriesForUser(userId int) ([]model.WeightEntry, error) {
	// if there is a DB error, return that error
	weightEntries := []model.WeightEntry{}
	for _, we := range wr.Db {
		if we.UserId == userId && we.Category == CATEGORY_DAILY {
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
		if we.EntryDate > maxDate && we.Category == CATEGORY_DAILY {
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
		if we.EntryDate > maxDate && we.Category == CATEGORY_TARGET {
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
