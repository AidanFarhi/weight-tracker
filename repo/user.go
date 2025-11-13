package repo

import (
	"errors"
	"weight-tracker/model"
)

type UserRepo struct {
	Db []model.User
}

func NewUserRepo() *UserRepo {
	// create a simple in-memory db for now
	return &UserRepo{
		Db: []model.User{
			{
				Id:       1,
				Email:    "example.email@gmail.com",
				Password: "secretpassword123",
			},
		},
	}
}

func (ur *UserRepo) ValidateCredentials(email, password string) error {
	for _, u := range ur.Db {
		if u.Email == email && u.Password == password {
			return nil
		}
	}
	return errors.New("no matching credentials")
}
