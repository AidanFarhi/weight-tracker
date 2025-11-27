package service

import (
	"time"
	"weight-tracker/model"
	"weight-tracker/repo"
)

type WeightService struct {
	wr *repo.WeightRepo
}

func NewWeightService(wr *repo.WeightRepo) *WeightService {
	return &WeightService{wr}
}

func (ws *WeightService) GetDailyWeightEntriesForUser(userId int) ([]model.WeightEntry, error) {
	var weightEntries []model.WeightEntry
	weightEntries, err := ws.wr.GetWeightEntriesForUser(userId, CategoryDaily)
	if err != nil {
		return weightEntries, ErrDBIssue
	}
	return weightEntries, nil
}

func (ws *WeightService) GetLatestDailyWeightEntryForUser(userId int) (model.WeightEntry, error) {
	var weightEntry model.WeightEntry
	weightEntry, err := ws.wr.GetLatestWeightEntryForUser(userId, CategoryDaily)
	if err != nil {
		return weightEntry, ErrDBIssue
	}
	return weightEntry, nil
}

func (ws *WeightService) GetLatestTargetWeightForUser(userId int) (model.WeightEntry, error) {
	var weightEntry model.WeightEntry
	weightEntry, err := ws.wr.GetLatestWeightEntryForUser(userId, CategoryTarget)
	if err != nil {
		return weightEntry, ErrDBIssue
	}
	return weightEntry, nil
}

func (ws *WeightService) CreateDailyWeightEntry(userId int, weight int, date string) error {
	weightEntry := model.WeightEntry{
		UserAccountId: userId,
		Weight:        weight,
		Category:      CategoryDaily,
		EntryDate:     date,
	}
	err := ws.wr.CreateWeightEntry(weightEntry)
	if err != nil {
		return ErrDBIssue
	}
	return nil
}

func (ws *WeightService) CreateTargetWeightEntry(userId int, weight int) error {
	weightEntry := model.WeightEntry{
		UserAccountId: userId,
		Weight:        weight,
		Category:      CategoryTarget,
		EntryDate:     time.Now().Format("2006-01-02"),
	}
	err := ws.wr.CreateWeightEntry(weightEntry)
	if err != nil {
		return ErrDBIssue
	}
	return nil
}
