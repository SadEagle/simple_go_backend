package crudl

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	db "movie_backend_go/db/sqlc"
)

func CreateMovieFavorite(ctx context.Context, querier db.Querier, favMovieCreate db.CreateMovieFavoriteParams) (db.FavoriteMovie, error) {
	favMovie, err := querier.CreateMovieFavorite(ctx, favMovieCreate)
	if err != nil {
		return db.FavoriteMovie{}, err
	}
	return favMovie, nil
}

func DeleteMovieFavorite(ctx context.Context, querier db.Querier, favMovieDelete db.DeleteMovieFavoriteParams) error {
	numDel, err := querier.DeleteMovieFavorite(ctx, favMovieDelete)
	if err != nil {
		return err
	}
	if numDel == 0 {
		return EmptyDeletionError
	}
	return nil
}

func GetMovieFavoriteList(ctx context.Context, querier db.Querier, userID pgtype.UUID) ([]pgtype.UUID, error) {
	favMovieList, err := querier.GetMovieFavoriteList(ctx, userID)
	if err != nil {
		return []pgtype.UUID{}, err
	}
	return favMovieList, nil
}
