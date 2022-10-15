// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: sub.sql

package models

import (
	"context"
	"database/sql"
	"time"
)

const deleteSub = `-- name: DeleteSub :exec
DELETE FROM sub
WHERE user_id = $1 and creator_id = $2
`

type DeleteSubParams struct {
	UserID    int64 `db:"user_id"`
	CreatorID int64 `db:"creator_id"`
}

func (q *Queries) DeleteSub(ctx context.Context, arg DeleteSubParams) error {
	_, err := q.db.ExecContext(ctx, deleteSub, arg.UserID, arg.CreatorID)
	return err
}

const deleteSubCreator = `-- name: DeleteSubCreator :exec
DELETE FROM sub
WHERE creator_id = $1
`

func (q *Queries) DeleteSubCreator(ctx context.Context, creatorID int64) error {
	_, err := q.db.ExecContext(ctx, deleteSubCreator, creatorID)
	return err
}

const deleteSubUser = `-- name: DeleteSubUser :exec
DELETE FROM sub
WHERE user_id = $1
`

func (q *Queries) DeleteSubUser(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteSubUser, userID)
	return err
}

const findSubID = `-- name: FindSubID :one
SELECT user_id, creator_id, activated_at, expires_at, status, price FROM sub
WHERE user_id = $1 and creator_id = $2 LIMIT 1
`

type FindSubIDParams struct {
	UserID    int64 `db:"user_id"`
	CreatorID int64 `db:"creator_id"`
}

func (q *Queries) FindSubID(ctx context.Context, arg FindSubIDParams) (*Sub, error) {
	row := q.db.QueryRowContext(ctx, findSubID, arg.UserID, arg.CreatorID)
	var i Sub
	err := row.Scan(
		&i.UserID,
		&i.CreatorID,
		&i.ActivatedAt,
		&i.ExpiresAt,
		&i.Status,
		&i.Price,
	)
	return &i, err
}

const insertSub = `-- name: InsertSub :one
INSERT INTO sub (
       user_id,
       creator_id,
       activated_at,
       expires_at,
       status,
       price
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING user_id, creator_id, activated_at, expires_at, status, price
`

type InsertSubParams struct {
	UserID      int64         `db:"user_id"`
	CreatorID   int64         `db:"creator_id"`
	ActivatedAt time.Time     `db:"activated_at"`
	ExpiresAt   time.Time     `db:"expires_at"`
	Status      SubStatus     `db:"status"`
	Price       sql.NullInt32 `db:"price"`
}

func (q *Queries) InsertSub(ctx context.Context, arg InsertSubParams) (*Sub, error) {
	row := q.db.QueryRowContext(ctx, insertSub,
		arg.UserID,
		arg.CreatorID,
		arg.ActivatedAt,
		arg.ExpiresAt,
		arg.Status,
		arg.Price,
	)
	var i Sub
	err := row.Scan(
		&i.UserID,
		&i.CreatorID,
		&i.ActivatedAt,
		&i.ExpiresAt,
		&i.Status,
		&i.Price,
	)
	return &i, err
}

const isExistSub = `-- name: IsExistSub :one
SELECT EXISTS (SELECT user_id, creator_id, activated_at, expires_at, status, price FROM sub
WHERE user_id = $1 and creator_id = $2)
`

type IsExistSubParams struct {
	UserID    int64 `db:"user_id"`
	CreatorID int64 `db:"creator_id"`
}

func (q *Queries) IsExistSub(ctx context.Context, arg IsExistSubParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, isExistSub, arg.UserID, arg.CreatorID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const listSub = `-- name: ListSub :many
SELECT user_id, creator_id, activated_at, expires_at, status, price FROM sub
WHERE TRUE
	AND (CASE WHEN $1::bool THEN price = $2 ELSE TRUE END)
	AND (CASE WHEN NOT $1::bool AND $3::bool THEN price >= $4 ELSE TRUE END)
	AND (CASE WHEN NOT $1::bool AND $5::bool THEN price <= $6 ELSE TRUE END)
	AND (CASE WHEN $7::bool THEN status = $8 ELSE TRUE END)
	AND (CASE WHEN $9::bool THEN creator_id = $10 ELSE TRUE END)
	AND (CASE WHEN $11::bool THEN user_id = $12 ELSE TRUE END)
OFFSET $13 LIMIT $14
`

type ListSubParams struct {
	IsPriceEq     bool          `db:"is_price_eq"`
	PriceEq       sql.NullInt32 `db:"price_eq"`
	IsPriceFrom   bool          `db:"is_price_from"`
	PriceFrom     sql.NullInt32 `db:"price_from"`
	IsPriceTo     bool          `db:"is_price_to"`
	PriceTo       sql.NullInt32 `db:"price_to"`
	IsStatusEq    bool          `db:"is_status_eq"`
	StatusEq      SubStatus     `db:"status_eq"`
	IsCreatorIDEq bool          `db:"is_creator_id_eq"`
	CreatorIDEq   int64         `db:"creator_id_eq"`
	IsUserIDEq    bool          `db:"is_user_id_eq"`
	UserIDEq      int64         `db:"user_id_eq"`
	PageNumber    int32         `db:"page_number"`
	PageSize      int32         `db:"page_size"`
}

func (q *Queries) ListSub(ctx context.Context, arg ListSubParams) ([]*Sub, error) {
	rows, err := q.db.QueryContext(ctx, listSub,
		arg.IsPriceEq,
		arg.PriceEq,
		arg.IsPriceFrom,
		arg.PriceFrom,
		arg.IsPriceTo,
		arg.PriceTo,
		arg.IsStatusEq,
		arg.StatusEq,
		arg.IsCreatorIDEq,
		arg.CreatorIDEq,
		arg.IsUserIDEq,
		arg.UserIDEq,
		arg.PageNumber,
		arg.PageSize,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Sub{}
	for rows.Next() {
		var i Sub
		if err := rows.Scan(
			&i.UserID,
			&i.CreatorID,
			&i.ActivatedAt,
			&i.ExpiresAt,
			&i.Status,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const saveSub = `-- name: SaveSub :one
INSERT INTO sub_history (
       user_id,
       creator_id,
       activated_at,
       expires_at,
       status,
       price
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING user_id, creator_id, activated_at, expires_at, status, price, sub_hist_id
`

type SaveSubParams struct {
	UserID      int64         `db:"user_id"`
	CreatorID   int64         `db:"creator_id"`
	ActivatedAt time.Time     `db:"activated_at"`
	ExpiresAt   time.Time     `db:"expires_at"`
	Status      SubStatus     `db:"status"`
	Price       sql.NullInt32 `db:"price"`
}

func (q *Queries) SaveSub(ctx context.Context, arg SaveSubParams) (*SubHistory, error) {
	row := q.db.QueryRowContext(ctx, saveSub,
		arg.UserID,
		arg.CreatorID,
		arg.ActivatedAt,
		arg.ExpiresAt,
		arg.Status,
		arg.Price,
	)
	var i SubHistory
	err := row.Scan(
		&i.UserID,
		&i.CreatorID,
		&i.ActivatedAt,
		&i.ExpiresAt,
		&i.Status,
		&i.Price,
		&i.SubHistID,
	)
	return &i, err
}

const updateSub = `-- name: UpdateSub :one
UPDATE sub
SET
	activated_at = COALESCE($1, activated_at),
	expires_at = COALESCE($2, expires_at),
	status = COALESCE($3, status),
	price = COALESCE($4, price)
WHERE
	user_id = $5 and creator_id = $6
RETURNING user_id, creator_id, activated_at, expires_at, status, price
`

type UpdateSubParams struct {
	ActivatedAt sql.NullTime  `db:"activated_at"`
	ExpiresAt   sql.NullTime  `db:"expires_at"`
	Status      NullSubStatus `db:"status"`
	Price       sql.NullInt32 `db:"price"`
	UserID      int64         `db:"user_id"`
	CreatorID   int64         `db:"creator_id"`
}

func (q *Queries) UpdateSub(ctx context.Context, arg UpdateSubParams) (*Sub, error) {
	row := q.db.QueryRowContext(ctx, updateSub,
		arg.ActivatedAt,
		arg.ExpiresAt,
		arg.Status,
		arg.Price,
		arg.UserID,
		arg.CreatorID,
	)
	var i Sub
	err := row.Scan(
		&i.UserID,
		&i.CreatorID,
		&i.ActivatedAt,
		&i.ExpiresAt,
		&i.Status,
		&i.Price,
	)
	return &i, err
}
