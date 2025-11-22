package service

import (
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
