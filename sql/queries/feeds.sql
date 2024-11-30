-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;


-- name: GetFeed :one
SELECT * FROM feeds
WHERE url = $1 LIMIT 1;

-- name: ListFeeds :many
SELECT feeds.name as feed_name, feeds.url as feed_url, users.name as user_name FROM feeds
LEFT JOIN users
ON users.id = feeds.user_id;
