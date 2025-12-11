package service

import (
	"errors"
	"testing"
	"weight-tracker/model"
	"weight-tracker/repo"
)

func TestLoginGivenNoIssues(t *testing.T) {
	password := "password"
	hashedPassword, _ := HashPassword(password)
	ur := &userRepoStub{user: model.User{Password: hashedPassword}}
	sr := &sessionRepoStub{}
	as := NewAuthService(ur, sr)
	_, err := as.Login("x@test.com", password)
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestLoginGivenNoRecordsFound(t *testing.T) {
	ur := &userRepoStub{err: repo.ErrNoRecordsFound}
	sr := &sessionRepoStub{}
	as := NewAuthService(ur, sr)
	_, err := as.Login("x@test.com", "password")
	if err != ErrEmailNotFound {
		t.Fatalf("expected ErrEmailNotFound, got %v", err)
	}
}

func TestLoginGivenInvalidPassword(t *testing.T) {
	ur := &userRepoStub{}
	sr := &sessionRepoStub{}
	as := NewAuthService(ur, sr)
	_, err := as.Login("x@test.com", "password")
	if err != ErrIncorrectPassword {
		t.Fatalf("expected ErrIncorrectPassword, got %v", err)
	}
}

func TestLoginGivenIssuCreatingSession(t *testing.T) {
	password := "password"
	hashedPassword, _ := HashPassword(password)
	ur := &userRepoStub{user: model.User{Password: hashedPassword}}
	sr := &sessionRepoStub{err: repo.ErrDuplicateSessionId}
	as := NewAuthService(ur, sr)
	_, err := as.Login("x@test.com", "password")
	if err != ErrDBIssue {
		t.Fatalf("expected ErrDBIssue, got %v", err)
	}
}

func TestValidateSessionGivenNoIssues(t *testing.T) {
	ur := &userRepoStub{}
	sr := &sessionRepoStub{}
	as := NewAuthService(ur, sr)
	_, err := as.ValidateSession("sessionid")
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestValidateSessionGivenNoRecordsFound(t *testing.T) {
	ur := &userRepoStub{}
	sr := &sessionRepoStub{err: repo.ErrNoRecordsFound}
	as := NewAuthService(ur, sr)
	_, err := as.ValidateSession("sessionid")
	if err != ErrNoSessionFound {
		t.Fatalf("expected ErrNoSessionFound, got %v", err)
	}
}

func TestValidateSessionGivenDBError(t *testing.T) {
	ur := &userRepoStub{}
	sr := &sessionRepoStub{err: errors.New("DB Error")}
	as := NewAuthService(ur, sr)
	_, err := as.ValidateSession("sessionid")
	if err != ErrDBIssue {
		t.Fatalf("expected ErrDBIssue, got %v", err)
	}
}
