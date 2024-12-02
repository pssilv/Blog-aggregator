-- name: CreateFeedFollow :one
WITH insert_feed_follow AS (
  INSERT INTO feeds_follows (id, created_at, updated_at, user_id, feed_id) 
  VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
  )
  RETURNING *
)
SELECT 
  insert_feed_follow.*,
  feeds.name AS feeds_name,
  users.name AS user_name
FROM insert_feed_follow

INNER JOIN users
ON insert_feed_follow.user_id = users.id

INNER JOIN feeds
ON insert_feed_follow.feed_id = feeds.id;
--

-- name: GetFeedFollowsForUser :many
SELECT feeds.url FROM feeds_follows
INNER JOIN feeds
ON feeds_follows.feed_id = feeds.id
WHERE feeds_follows.user_id = $1;
--

-- name: DeleteFeedFollow :exec
DELETE FROM feeds_follows
WHERE feeds_follows.user_id = $1
AND feeds_follows.feed_id = $2;
--


