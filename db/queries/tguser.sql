-- name: FindTguserID :one
SELECT * FROM tguser
WHERE user_id = $1 LIMIT 1;

-- name: FindTguserTelegramID :one
SELECT * FROM tguser
WHERE telegram_id = $1 LIMIT 1;

-- name: InsertTguser :one
INSERT INTO tguser (
       TelegramID,
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
	TelegramID = COALESCE(sqlc.narg(telegramid), telegramid),
	Username = COALESCE(sqlc.narg(username), username),
	Status = COALESCE(sqlc.narg(status), status)
WHERE
	user_id = sqlc.arg(user_id)
RETURNING *;
