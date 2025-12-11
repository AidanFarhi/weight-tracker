package service

import (
	"errors"
	"testing"
	"weight-tracker/repo"
)

func TestRegisterGivenNoIssues(t *testing.T) {
	ur := &userRepoStub{}
	sr := &sessionRepoStub{}
	svc := NewRegisterService(ur, sr)
	_, err := svc.Register("x@test.com", "password")
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestRegisterGivenUserAlreadyExists(t *testing.T) {
	ur := &userRepoStub{exists: true}
	sr := &sessionRepoStub{}
	svc := NewRegisterService(ur, sr)
	_, err := svc.Register("x@test.com", "password")
	if err != ErrUserAlreadyExists {
		t.Fatalf("expected ErrUserAlreadyExists, got %v", err)
	}
}

func TestRegisterUserGivenDBIssueDuringUserCreation(t *testing.T) {
	ur := &userRepoStub{
		exists: false,
		err:    errors.New("DB Error"),
	}
	sr := &sessionRepoStub{}
	svc := NewRegisterService(ur, sr)
	_, err := svc.Register("x@test.com", "password")
	if err != ErrDBIssue {
		t.Fatalf("expected ErrDBIssue, got %v", err)
	}
}

func TestRegisterUserGivenDBIssueDuringCreatingSession(t *testing.T) {
	ur := &userRepoStub{exists: false}
	sr := &sessionRepoStub{err: repo.ErrDuplicateSessionId}
	svc := NewRegisterService(ur, sr)
	_, err := svc.Register("x@test.com", "password")
	if err != ErrDBIssue {
		t.Fatalf("expected ErrDBIssue, got %v", err)
	}
}
