-- name: FindCreatorID :one
SELECT * FROM creator
WHERE creator_id = $1 LIMIT 1;

-- name: FindCreatorTelegramID :one
SELECT * FROM creator
WHERE creator_id = $1 LIMIT 1;

-- name: IsExistCreator :one
SELECT EXISTS (SELECT * FROM creator
WHERE telegram_id = $1);

-- name: InsertCreator :one
INSERT INTO Creator (
       	Telegram_id,
	Username,
       	Password,
	Email,
	Chan_Name
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: DeleteCreator :exec
DELETE FROM creator
WHERE creator_id == $1;

-- name: UpdateCreator :one
UPDATE creator
SET
	Telegram_id = COALESCE(sqlc.narg(telegram_id), telegram_id),
	Username = COALESCE(sqlc.narg(username), username),
	Password = COALESCE(sqlc.narg(password), password),
	Email = COALESCE(sqlc.narg(email), email),
	Chan_Name = COALESCE(sqlc.narg(chan_name), chan_name)
WHERE
	creator_id = sqlc.arg(creator_id)
RETURNING *;

-- name: ListCreator :many
SELECT * FROM creator OFFSET $1 LIMIT $2;
