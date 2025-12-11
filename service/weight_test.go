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

func TestGetDailyWeightEntryForUserGivenNoIssues(t *testing.T) {
	wr := &weightRepoStub{}
	ws := NewWeightService(wr)
	_, err := ws.GetLatestDailyWeightEntryForUser(1)
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestGetDailyWeightEntryForUserGivenDBError(t *testing.T) {
	wr := &weightRepoStub{err: errors.New("DB error")}
	ws := NewWeightService(wr)
	_, err := ws.GetLatestDailyWeightEntryForUser(1)
	if err != ErrDBIssue {
		t.Fatalf("expected ErrDBIssue, got %v", err)
	}
}

func TestGetTargetWeightForUserGivenNoIssues(t *testing.T) {
	wr := &weightRepoStub{}
	ws := NewWeightService(wr)
	_, err := ws.GetLatestTargetWeightForUser(1)
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestGetTargetWeightForUserGivenDBError(t *testing.T) {
	wr := &weightRepoStub{err: errors.New("DB error")}
	ws := NewWeightService(wr)
	_, err := ws.GetLatestTargetWeightForUser(1)
	if err != ErrDBIssue {
		t.Fatalf("expected ErrDBIssue, got %v", err)
	}
}

func TestCreateDailyWeightEntryGivenNoIssues(t *testing.T) {
	wr := &weightRepoStub{}
	ws := NewWeightService(wr)
	err := ws.CreateDailyWeightEntry(1, 100, "2025-01-01")
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestCreateDailyWeightEntryGivenBadDate(t *testing.T) {
	wr := &weightRepoStub{}
	ws := NewWeightService(wr)
	err := ws.CreateDailyWeightEntry(1, 100, "BAD_DATE")
	if err != ErrDateParseIssue {
		t.Fatalf("expected ErrDateParseIssue, got %v", err)
	}
}

func TestCreateDailyWeightEntryGivenDBError(t *testing.T) {
	wr := &weightRepoStub{err: errors.New("DB Error")}
	ws := NewWeightService(wr)
	err := ws.CreateDailyWeightEntry(1, 100, "2025-01-01")
	if err != ErrDBIssue {
		t.Fatalf("expected ErrDBIssue, got %v", err)
	}
}

func TestCreateTargetWeightEntryGivenNoIssues(t *testing.T) {
	wr := &weightRepoStub{}
	ws := NewWeightService(wr)
	err := ws.CreateTargetWeightEntry(1, 100)
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestCreateTargetWeightEntryGivenDBError(t *testing.T) {
	wr := &weightRepoStub{err: errors.New("DB Error")}
	ws := NewWeightService(wr)
	err := ws.CreateTargetWeightEntry(1, 100)
	if err != ErrDBIssue {
		t.Fatalf("expected ErrDBIssue, got %v", err)
	}
}
