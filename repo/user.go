package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db}
}

func (ur *UserRepo) CreateUser(email, password string) (int, error) {
	var userId int
	c := context.Background()
	q := "INSERT INTO user_account (email, password) VALUES ($1, $2) RETURNING id"
	err := ur.db.QueryRow(c, q, email, password).Scan(&userId)
	if err != nil {
		return -1, err
	}
	return userId, nil
}

func (ur *UserRepo) CheckIfUserExists(email string) (bool, error) {
	var userExists bool
	c := context.Background()
	q := "SELECT EXISTS(SELECT 1 FROM user_account WHERE email = $1)"
	err := ur.db.QueryRow(c, q, email).Scan(&userExists)
	if err != nil {
		return false, err
	}
	return userExists, nil
}

func (ur *UserRepo) GetIdForUser(email, password string) (int, error) {
	var userId int
	c := context.Background()
	q := "SELECT id FROM user_account WHERE email = $1 AND password = $2"
	err := ur.db.QueryRow(c, q, email, password).Scan(&userId)
	if err != nil {
		return -1, err
	}
	return userId, nil
}
