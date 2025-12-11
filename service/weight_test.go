package service

import (
	"errors"
	"testing"
	"weight-tracker/model"
)

type weightRepoStub struct {
	weightEntry   model.WeightEntry
	weightEntries []model.WeightEntry
	err           error
}

func (w *weightRepoStub) GetNWeightEntriesByUserId(userId int, category string, n int) ([]model.WeightEntry, error) {
	return w.weightEntries, w.err
}

func (w *weightRepoStub) CreateWeightEntry(model.WeightEntry) error {
	return w.err
}

func (w *weightRepoStub) GetLatestWeightEntryForUser(userId int, category string) (model.WeightEntry, error) {
	return w.weightEntry, w.err
}

func TestGetNDailyWeightEntriesByUserIdGivenNoIssues(t *testing.T) {
	wr := &weightRepoStub{}
	ws := NewWeightService(wr)
	_, err := ws.GetNDailyWeightEntriesByUserId(1, 10)
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestGetNDailyWeightEntriesByUserIdGivenDBError(t *testing.T) {
	wr := &weightRepoStub{err: errors.New("DB error")}
	ws := NewWeightService(wr)
	_, err := ws.GetNDailyWeightEntriesByUserId(1, 10)
	if err != ErrDBIssue {
		t.Fatalf("expected ErrDBIssue, got %v", err)
	}
}
