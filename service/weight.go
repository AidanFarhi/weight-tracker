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

func (ws *WeightService) GetNDailyWeightEntriesByUserId(userId, n int) ([]model.WeightEntry, error) {
	var weightEntries []model.WeightEntry
	// figure out a better way?
	if n <= 0 {
		n = 100_000
	}
	weightEntries, err := ws.wr.GetNWeightEntriesByUserId(userId, CategoryDaily, n)
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
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ErrDateParseIssue
	}
	now := time.Now()
	entryDate := time.Date(
		parsedDate.Year(),
		parsedDate.Month(),
		parsedDate.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Nanosecond(),
		time.Local,
	)
	weightEntry := model.WeightEntry{
		UserAccountId: userId,
		Weight:        weight,
		Category:      CategoryDaily,
		EntryDate:     entryDate,
	}
	err = ws.wr.CreateWeightEntry(weightEntry)
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
		EntryDate:     time.Now(),
	}
	err := ws.wr.CreateWeightEntry(weightEntry)
	if err != nil {
		return ErrDBIssue
	}
	return nil
}
