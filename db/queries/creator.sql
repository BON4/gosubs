-- name: FindCreatorID :one
SELECT * FROM creator
WHERE creator_id = $1 LIMIT 1;

-- name: FindCreatorTelegramID :one
SELECT * FROM creator
WHERE creator_id = $1 LIMIT 1;

-- name: InsertCreator :one
INSERT INTO Creator (
       	TelegramID,
	Username,
       	Password,
	Email,
	ChanName
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: DeleteCreator :exec
DELETE FROM creator
WHERE creator_id == $1;

-- name: UpdateCreator :one
UPDATE creator
SET
	TelegramID = COALESCE(sqlc.narg(telegramid), telegramid),
	Username = COALESCE(sqlc.narg(username), username),
	Password = COALESCE(sqlc.narg(password), password),
	Email = COALESCE(sqlc.narg(email), email),
	ChanName = COALESCE(sqlc.narg(channame), channame)
WHERE
	creator_id = sqlc.arg(creator_id)
RETURNING *;
