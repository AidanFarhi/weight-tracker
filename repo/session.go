package repo

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"weight-tracker/model"
)

type SessionRepo struct {
	Db []model.Session
}

func NewSessionRepo() *SessionRepo {
	return &SessionRepo{
		Db: []model.Session{},
	}
}

func (sr *SessionRepo) CreateSession(userId int) (string, error) {
	sessionId := generateSessionId()
	newSession := model.Session{
		Id:     sessionId,
		UserId: userId,
	}
	sr.Db = append(sr.Db, newSession)
	return sessionId, nil
}

func (sr *SessionRepo) DeleteSession(sessionId string) error {
	sessionsToKeep := []model.Session{}
	for _, s := range sr.Db {
		if s.Id != sessionId {
			sessionsToKeep = append(sessionsToKeep, s)
		}
	}
	if len(sessionsToKeep) == len(sr.Db) {
		return errors.New("no session found")
	}
	sr.Db = sessionsToKeep
	return nil
}

func (sr *SessionRepo) GetUserIdForSession(sessionId string) (int, error) {
	// TODO: throw an actual DB error if there is one while querying
	for _, s := range sr.Db {
		if s.Id == sessionId {
			return s.UserId, nil
		}
	}
	return -1, nil
}

func generateSessionId() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}
