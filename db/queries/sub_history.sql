-- name: FindSubHistory :one
SELECT * FROM sub_history
WHERE sub_hist_id = $1 LIMIT 1;

-- name: InsertSubHistory :one
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
