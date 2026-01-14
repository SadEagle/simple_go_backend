-- COALESCE and LEFT JOIN because rating exist if only at least one person rate movie
-- TODO: check will 2 subquery be faster than one

-- name: GetMovieByID :one
SELECT id, title, COALESCE(amount_rates, 0), COALESCE(rating, 0), created_at
FROM (
  select * from movie where id = $1
  ) m
LEFT JOIN ( 
  select * from movie_rating_view where movie_id = $1
) mrv ON m.id = mrv.movie_id;

-- TODO: is it optimal (same as above query)

-- name: GetMovieByTitle :one
SELECT id, title, COALESCE(amount_rates, 0), COALESCE(rating, 0), created_at
FROM (
  select * from movie where title = $1
  ) m
LEFT JOIN movie_rating_view ON m.id = mrv.movie_id;

-- name: GetMovieList :many
SELECT id, title, created_at
FROM movie;

-- name: CreateMovie :one
INSERT INTO movie(title)
VALUES ($1)
RETURNING *;

-- name: UpdateMovie :one
UPDATE movie SET
  title = COALESCE($2, title)
WHERE id = $1
RETURNING *;

-- name: DeleteMovie :execrows
DELETE FROM movie
WHERE id = $1;
