package crudl

import (
	"context"
	db "movie_backend_go/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func CreateUser(ctx context.Context, querier db.Querier, userCreate db.CreateUserParams) (db.UserDatum, error) {
	user, err := querier.CreateUser(ctx, userCreate)
	if err != nil {
		return db.UserDatum{}, err
	}
	return user, nil
}

func DeleteUser(ctx context.Context, querier db.Querier, userID pgtype.UUID) error {
	numDel, err := querier.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}
	if numDel == 0 {
		return EmptyDeletionError
	}
	return nil
}

func GetUserByID(ctx context.Context, querier db.Querier, userID pgtype.UUID) (db.UserDatum, error) {
	user, err := querier.GetUserByID(ctx, userID)
	if err != nil {
		return db.UserDatum{}, err
	}
	return user, nil
}

func GetUserByLogin(ctx context.Context, querier db.Querier, login string) (db.UserDatum, error) {
	user, err := querier.GetUserByLogin(ctx, login)
	if err != nil {
		return db.UserDatum{}, err
	}
	return user, nil
}

func GetUserList(ctx context.Context, querier db.Querier) ([]db.UserDatum, error) {
	userList, err := querier.GetUserList(ctx)
	if err != nil {
		return []db.UserDatum{}, err
	}
	return userList, nil
}

func UpdateUser(ctx context.Context, querier db.Querier, userUpdate db.UpdateUserParams) (db.UserDatum, error) {
	user, err := querier.UpdateUser(ctx, userUpdate)
	if err != nil {
		return db.UserDatum{}, err
	}
	return user, nil
}
