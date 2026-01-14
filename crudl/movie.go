package crudl

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	db "movie_backend_go/db/sqlc"
)

func CreateMovie(ctx context.Context, querier db.Querier, title string) (db.Movie, error) {
	movie, err := querier.CreateMovie(ctx, title)
	if err != nil {
		return db.Movie{}, err
	}
	return movie, err
}

func DeleteMovie(ctx context.Context, querier db.Querier, movieID pgtype.UUID) error {
	numDel, err := querier.DeleteMovie(ctx, movieID)
	if err != nil {
		return err
	}
	if numDel == 0 {
		return EmptyDeletionError
	}
	return nil
}

func GetMovieByID(ctx context.Context, querier db.Querier, movieID pgtype.UUID) (db.GetMovieByIDRow, error) {
	movie, err := querier.GetMovieByID(ctx, movieID)
	if err != nil {
		return db.GetMovieByIDRow{}, err
	}
	return movie, nil
}

func GetMovieByTitle(ctx context.Context, querier db.Querier, movieTitle string) (db.GetMovieByTitleRow, error) {
	movie, err := querier.GetMovieByTitle(ctx, movieTitle)
	if err != nil {
		return db.GetMovieByTitleRow{}, err
	}
	return movie, nil
}

func GetMovieList(ctx context.Context, querier db.Querier) ([]db.Movie, error) {
	movieList, err := querier.GetMovieList(ctx)
	if err != nil {
		return []db.Movie{}, err
	}
	return movieList, nil
}

func UpdateMovie(ctx context.Context, querier db.Querier, movieUpdate db.UpdateMovieParams) (db.Movie, error) {
	movie, err := querier.UpdateMovie(ctx, movieUpdate)
	if err != nil {
		return db.Movie{}, err
	}
	return movie, nil
}
