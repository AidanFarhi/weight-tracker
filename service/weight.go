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
	return ws.wr.GetDailyWeightEntriesForUser(userId)
}

func (ws *WeightService) GetLatestDailyWeightEntryForUser(userId int) (model.WeightEntry, error) {
	return ws.wr.GetLatestDailyWeightForUser(userId)
}

func (ws *WeightService) GetLatestTargetWeightForUser(userId int) (model.WeightEntry, error) {
	return ws.wr.GetLatestTargetWeightForUser(userId)
}

func (ws *WeightService) CreateWeightEntry(userId int, weight int, category string, date string) error {
	if category == repo.CATEGORY_TARGET {
		date = time.Now().Format("2006-01-02")
	}
	weightEntry := model.WeightEntry{
		UserId:    userId,
		Weight:    weight,
		Category:  category,
		EntryDate: date,
	}
	return ws.wr.CreateWeightEntry(weightEntry)
}
