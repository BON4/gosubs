-- name: FindTguserID :one
SELECT * FROM tguser
WHERE user_id = $1 LIMIT 1;

-- name: FindTguserTelegramID :one
SELECT * FROM tguser
WHERE telegram_id = $1 LIMIT 1;


-- name: IsExistTguser :one
SELECT EXISTS (SELECT * FROM tguser
WHERE telegram_id = $1);


-- name: InsertTguser :one
INSERT INTO tguser (
       Telegram_id,
       Username,
       Status
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: DeleteTguser :exec
DELETE FROM tguser
WHERE user_id == $1;

-- name: UpdateTguser :one
UPDATE tguser
SET
	Telegram_id = COALESCE(sqlc.narg(telegram_id), telegram_id),
	Username = COALESCE(sqlc.narg(username), username),
	Status = COALESCE(sqlc.narg(status), status)
WHERE
	user_id = sqlc.arg(user_id)
RETURNING *;

-- name: ListTguser :many
SELECT * FROM tguser OFFSET $1 LIMIT $2;
