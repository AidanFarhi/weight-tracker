package repo

import (
	"context"
	"weight-tracker/model"

	"github.com/jackc/pgx/v5"
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
		return 0, ErrSQLDBIssue
	}
	return userId, nil
}

func (ur *UserRepo) CheckIfUserExists(email string) (bool, error) {
	var userExists bool
	c := context.Background()
	q := "SELECT EXISTS(SELECT 1 FROM user_account WHERE email = $1)"
	err := ur.db.QueryRow(c, q, email).Scan(&userExists)
	if err != nil {
		return false, ErrSQLDBIssue
	}
	return userExists, nil
}

func (ur *UserRepo) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	c := context.Background()
	q := "SELECT id, password FROM user_account WHERE email = $1"
	err := ur.db.QueryRow(c, q, email).Scan(&user.Id, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user, ErrNoRecordsFound
		}
		return user, ErrSQLDBIssue
	}
	return user, nil
}
