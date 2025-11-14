package repo

import (
	"errors"
	"weight-tracker/model"
)

type UserRepo struct {
	Db            []model.User
	userIdCounter int
}

func NewUserRepo() *UserRepo {
	// create a simple in-memory db for now
	return &UserRepo{
		userIdCounter: 2,
		Db: []model.User{
			{
				Id:       1,
				Email:    "example.email@gmail.com",
				Password: "secretpassword123",
			},
		},
	}
}

func (ur *UserRepo) CreateUser(email, password string) (int, error) {
	newUser := model.User{
		Id:       ur.userIdCounter,
		Email:    email,
		Password: password,
	}
	ur.Db = append(ur.Db, newUser)
	ur.userIdCounter = ur.userIdCounter + 1
	return newUser.Id, nil
}

// TODO: change the return to only throw an error if there is a SQL DB error.
// otherwise, it should be a boolean?
func (ur *UserRepo) CheckIfUserExists(email string) error {
	for _, u := range ur.Db {
		if u.Email == email {
			return nil
		}
	}
	return errors.New("no matching user")
}

func (ur *UserRepo) GetIdForUser(email, password string) (int, error) {
	for _, u := range ur.Db {
		if u.Email == email && u.Password == password {
			return u.Id, nil
		}
	}
	return -1, errors.New("no matching credentials")
}
