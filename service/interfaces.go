package service

import "weight-tracker/model"

type UserRepo interface {
	CheckIfUserExists(email string) (bool, error)
	CreateUser(email, hashedPassword string) (int, error)
}

type SessionRepo interface {
	CreateSession(sessionId string, userId int) error
}

type WeightRepo interface {
	GetNWeightEntriesByUserId(userId int, category string, n int) ([]model.WeightEntry, error)
	GetLatestWeightEntryForUser(userId int, category string) (model.WeightEntry, error)
	CreateWeightEntry(weightEntry model.WeightEntry) error
}
