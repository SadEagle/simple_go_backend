package crudl

import (
	"context"
	db "movie_backend_go/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func CreateMovieRating(ctx context.Context, querier db.Querier, movieRatingCreate db.CreateMovieRatingParams) (db.RatedMovie, error) {
	movieRating, err := querier.CreateMovieRating(ctx, movieRatingCreate)
	if err != nil {
		return db.RatedMovie{}, err
	}
	return movieRating, err
}

func DeleteMovieRating(ctx context.Context, querier db.Querier, movieRatingDelete db.DeleteMovieRatingParams) error {
	numDel, err := querier.DeleteMovieRating(ctx, movieRatingDelete)
	if err != nil {
		return err
	}
	if numDel == 0 {
		return EmptyDeletionError
	}
	return nil
}

func GetMovieRatingList(ctx context.Context, querier db.Querier, userID pgtype.UUID) ([]db.GetMovieRatingListRow, error) {
	movieRatingList, err := querier.GetMovieRatingList(ctx, userID)
	if err != nil {
		return []db.GetMovieRatingListRow{}, err
	}
	return movieRatingList, nil
}

func UpdateMovieRating(ctx context.Context, querier db.Querier, movieRatingUpdate db.UpdateMoveRatingParams) (db.RatedMovie, error) {
	movieRating, err := querier.UpdateMoveRating(ctx, movieRatingUpdate)
	if err != nil {
		return db.RatedMovie{}, err
	}
	return movieRating, nil
}
