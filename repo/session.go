package repo

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionRepo struct {
	db *pgxpool.Pool
}

func NewSessionRepo(db *pgxpool.Pool) *SessionRepo {
	return &SessionRepo{db}
}

func (sr *SessionRepo) CreateSession(sessionId string, userId int) error {
	c := context.Background()
	q := "INSERT INTO user_session (id, user_account_id) VALUES ($1, $2)"
	_, err := sr.db.Exec(c, q, sessionId, userId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PGUniqueConstraintViolationErrorCode {
				return ErrDuplicateSessionId
			}
		}
	}
	return err
}

func (sr *SessionRepo) DeleteSession(sessionId string) error {
	c := context.Background()
	q := "DELETE FROM user_session WHERE sessionId = $1"
	res, err := sr.db.Exec(c, q, sessionId)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrNoRecordsFound
	}
	return nil
}

func (sr *SessionRepo) GetUserIdForSession(sessionId string) (int, error) {
	var userId int
	c := context.Background()
	q := "SELECT user_account_id FROM user_session WHERE id = $1"
	err := sr.db.QueryRow(c, q, sessionId).Scan(&userId)
	if err == pgx.ErrNoRows {
		return 0, ErrNoRecordsFound
	}
	return userId, err
}
