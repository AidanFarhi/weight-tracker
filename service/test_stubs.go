package service

import "weight-tracker/model"

type userRepoStub struct {
	user   model.User
	exists bool
	err    error
}

func (u *userRepoStub) CheckIfUserExists(email string) (bool, error) {
	return u.exists, u.err
}
func (u *userRepoStub) CreateUser(email, hashed string) (int, error) {
	return 123, nil
}

func (u *userRepoStub) GetUserByEmail(email string) (model.User, error) {
	return u.user, u.err
}

type sessionRepoStub struct {
	err error
}

func (s *sessionRepoStub) CreateSession(id string, uid int) error {
	return s.err
}

func (s *sessionRepoStub) GetUserIdForSession(sessionId string) (int, error) {
	return 1, s.err
}
