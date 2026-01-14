-- name: GetMovieRatingList :many
SELECT movie_id, rating
FROM rated_movie
WHERE user_id = $1;

-- name: CreateMovieRating :one
INSERT INTO rated_movie (user_id, movie_id, rating)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateMoveRating :one
UPDATE rated_movie SET
rating = COALESCE($3, rating)
WHERE user_id = $1 
  AND movie_id = $2
RETURNING *;


-- name: DeleteMovieRating :execrows
DELETE FROM rated_movie
WHERE user_id = $1 AND movie_id = $2;

