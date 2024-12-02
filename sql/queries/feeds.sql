-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id, last_fetched_at)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
)
RETURNING *;
--

-- name: GetFeed :one
SELECT * FROM feeds
WHERE url = $1 LIMIT 1;
--

-- name: ListFeeds :many
SELECT 
  feeds.name as feed_name, 
  feeds.url as feed_url, 
  users.name as user_name 
FROM feeds

INNER JOIN users
ON users.id = feeds.user_id;
--

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = NOW(), 
updated_at = NOW()
WHERE id = $1
RETURNING *;
--

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY feeds.last_fetched_at ASC NULLS FIRST
LIMIT 1;
--
