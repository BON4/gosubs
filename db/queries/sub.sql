-- name: FindSubID :one
SELECT * FROM sub
WHERE user_id = $1 and creator_id = $2 LIMIT 1;

-- name: InsertSub :one
INSERT INTO sub (
       user_id,
       creator_id,
       activated_at,
       expires_at,
       status,
       price
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: SaveSub :one
INSERT INTO sub_history (
       user_id,
       creator_id,
       activated_at,
       expires_at,
       status,
       price
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: IsExistSub :one
SELECT EXISTS (SELECT * FROM sub
WHERE user_id = $1 and creator_id = $2);

-- name: DeleteSubUser :exec
DELETE FROM sub
WHERE user_id = $1;

-- name: DeleteSubCreator :exec
DELETE FROM sub
WHERE creator_id = $1;

-- name: DeleteSub :exec
DELETE FROM sub
WHERE user_id = $1 and creator_id = $2;

-- name: UpdateSub :one
UPDATE sub
SET
	activated_at = COALESCE(sqlc.narg(activated_at), activated_at),
	expires_at = COALESCE(sqlc.narg(expires_at), expires_at),
	status = @status,
	price = COALESCE(sqlc.narg(price), price)
WHERE
	user_id = sqlc.arg(user_id) and creator_id = sqlc.arg(creator_id)
RETURNING *;


-- name: ListSub :many
SELECT * FROM sub
WHERE TRUE
	AND (CASE WHEN @is_price_eq::bool THEN price = @price_eq ELSE TRUE END)
	AND (CASE WHEN NOT @is_price_eq::bool AND @is_price_from::bool THEN price >= @price_from ELSE TRUE END)
	AND (CASE WHEN NOT @is_price_eq::bool AND @is_price_to::bool THEN price <= @price_to ELSE TRUE END)
	AND (CASE WHEN @is_status_eq::bool THEN status = @status_eq ELSE TRUE END)
	AND (CASE WHEN @is_creator_id_eq::bool THEN creator_id = @creator_id_eq ELSE TRUE END)
	AND (CASE WHEN @is_user_id_eq::bool THEN user_id = @user_id_eq ELSE TRUE END)
OFFSET @page_number LIMIT @page_size;


-- name: DeleteAll :exec
TRUNCATE TABLE sub_history, sub, creator, tguser;
