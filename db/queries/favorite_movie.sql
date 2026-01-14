-- name: CreateMovieFavorite :one
INSERT INTO favorite_movie(user_id, movie_id)
VALUES ($1, $2)
RETURNING user_id, movie_id;

-- name: GetMovieFavoriteList :many
SELECT movie_id
FROM favorite_movie
WHERE user_id = $1;

-- name: DeleteMovieFavorite :execrows
DELETE FROM favorite_movie
WHERE user_id = $1 AND movie_id = $2;
