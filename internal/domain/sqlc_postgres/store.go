package models

import (
	"context"
	"database/sql"
	"fmt"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
	DeleteCreatorTx(ctx context.Context, arg DeleteCreatorTxParams) error
	DeleteUserTx(ctx context.Context, arg DeleteUserTxParams) error
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type DeleteUserTxParams struct {
	UserID int64 `json:"user_id"`
}

func (store *SQLStore) DeleteUserTx(ctx context.Context, arg DeleteUserTxParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		if err := q.DeleteSubUser(ctx, arg.UserID); err != nil {
			return err
		}
		if err := q.DeleteTguser(ctx, arg.UserID); err != nil {
			return err
		}

		return nil
	})

	return err
}

type DeleteCreatorTxParams struct {
	CreatorID int64 `json:"creator_id"`
}

func (store *SQLStore) DeleteCreatorTx(ctx context.Context, arg DeleteCreatorTxParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		if err := q.DeleteSubCreator(ctx, arg.CreatorID); err != nil {
			return err
		}
		if err := q.DeleteCreator(ctx, arg.CreatorID); err != nil {
			return err
		}

		return nil
	})

	return err
}
