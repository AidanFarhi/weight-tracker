package repo

import (
	"context"
	"weight-tracker/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WeightRepo struct {
	db *pgxpool.Pool
}

func NewWeightRepo(db *pgxpool.Pool) *WeightRepo {
	return &WeightRepo{db}
}

func (wr *WeightRepo) GetWeightEntriesForUser(userId int, category string) ([]model.WeightEntry, error) {
	var weightEntries []model.WeightEntry
	c := context.Background()
	q := `
		SELECT id, user_account_id, weight, entry_date, category
		FROM weight_entry
		WHERE user_account_id = $1 AND category = $2
	`
	rows, err := wr.db.Query(c, q, userId, category)
	if err != nil {
		return weightEntries, ErrSQLDBIssue
	}
	defer rows.Close()
	for rows.Next() {
		var we model.WeightEntry
		err := rows.Scan(
			&we.Id,
			&we.UserAccountId,
			&we.Weight,
			&we.EntryDate,
			&we.Category,
		)
		if err != nil {
			return weightEntries, ErrSQLDBIssue
		}
		weightEntries = append(weightEntries, we)
	}
	return weightEntries, nil
}

func (wr *WeightRepo) GetLatestWeightEntryForUser(userId int, category string) (model.WeightEntry, error) {
	var latestWeightEntry model.WeightEntry
	c := context.Background()
	q := `
		SELECT id, user_account_id, weight, entry_date, category
		FROM weight_entry
		WHERE user_account_id = $1 AND category = $2
		ORDER BY entry_date DESC
		LIMIT 1
	`
	err := wr.db.QueryRow(c, q, userId, category).Scan(
		&latestWeightEntry.Id,
		&latestWeightEntry.UserAccountId,
		&latestWeightEntry.Weight,
		&latestWeightEntry.EntryDate,
		&latestWeightEntry.Category,
	)
	if err != nil && err != pgx.ErrNoRows {
		return latestWeightEntry, ErrSQLDBIssue
	}
	return latestWeightEntry, nil
}

func (wr *WeightRepo) CreateWeightEntry(weightEntry model.WeightEntry) error {
	c := context.Background()
	q := `
		INSERT INTO weight_entry (user_account_id, weight, entry_date, category)
		VALUES ($1, $2, $3, $4)
	`
	_, err := wr.db.Exec(
		c,
		q,
		weightEntry.UserAccountId,
		weightEntry.Weight,
		weightEntry.EntryDate,
		weightEntry.Category,
	)
	if err != nil {
		return ErrSQLDBIssue
	}
	return nil
}
