package reqmodel

import (
	"github.com/jackc/pgx/v5/pgtype"
	db "movie_backend_go/db/sqlc"
)

type FavoriteMovieListResponse struct {
	UserID           pgtype.UUID   `json:"user_id"`
	FavoriteMovieIDs []pgtype.UUID `json:"favorite_movie_ids"`
}

// Create and Update requests contain same possible parameters
type MovieRequest struct {
	Title string `json:"title"`
}

type MovieListResponse struct {
	MovieList []db.Movie `json:"movie_list"`
}

type RatedMovieRequest struct {
	MovieID pgtype.UUID `json:"movie_id"`
	Rating  pgtype.Int4 `json:"rating"`
}

type RatedMovieListResponse struct {
	UserID         pgtype.UUID                `json:"user_id"`
	RatedMovieList []db.GetMovieRatingListRow `json:"rated_movie_list"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
type UserListResponse struct {
	UserList []db.UserDatum `json:"user_list"`
}
