package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"movie_backend_go/crudl"
	"movie_backend_go/models"

	"net/http"
)

// @Summary      Show user
// @Description  Get user by id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.UserData
// @Failure      404  {object}	map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /user/{id} [get]
func GetUserHandlerMake(db *sql.DB) http.HandlerFunc {
	GetUserHandler := func(rw http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		user, err := crudl.GetUserDB(db, id)
		if err != nil {
			log.Println(err)
			http.Error(rw, fmt.Sprintf("Can't get user id: %s\n", id), 404)
			return
		}
		writeResponseBody(rw, user, "user")

	}
	return GetUserHandler
}

// @Summary      Show user list
// @Description  Get user list
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.UserDataList
// @Failure      404  {object}	map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /user [get]
func GetUserListHandlerMake(db *sql.DB) http.HandlerFunc {
	GetUserListHandler := func(rw http.ResponseWriter, r *http.Request) {
		userList, err := crudl.GetUserListDB(db)
		if err != nil {
			log.Println(err)
			http.Error(rw, "Can't get user list", 500)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		writeResponseBody(rw, userList, "user list")
	}
	return GetUserListHandler
}

// @Summary      Update user
// @Description  Update user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Param        request 		body	models.UpdateUserDataRequest  true  "User creation data"
// @Success      200  {object}  models.UserData
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /user/{id} [PATCH]
func UpdateUserHandlerMake(db *sql.DB) http.HandlerFunc {
	UpdateUserHandler := func(rw http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields() // Strict parsing

		var updateUserdata models.UpdateUserDataRequest
		err := decoder.Decode(&updateUserdata)
		if err != nil && err != io.EOF {
			log.Println(err)
			http.Error(rw, "Can't proceed body request", 400)
			return
		}
		user_id := r.PathValue("id")

		fmt.Println(updateUserdata)
		user, err := crudl.UpdateUserDataDB(db, updateUserdata, user_id)
		if err != nil {
			log.Println(err)
			http.Error(rw, "Can't update user", 404)
			return
		}

		writeResponseBody(rw, user, "user")
	}
	return UpdateUserHandler
}

// @Summary      Create user
// @Description  Create user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request 		body	models.CreateUserDataRequest  true  "User creation data"
// @Success      201  {object}  models.UserData
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /user [post]
func CreateUserHandlerMake(db *sql.DB) http.HandlerFunc {
	CreateUserHandler := func(rw http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields() // Strict parsing

		var createUserData models.CreateUserDataRequest
		err := decoder.Decode(&createUserData)
		if err != nil && err != io.EOF {
			log.Println(err)
			http.Error(rw, "Can't proceed body request", 400)
			return
		}

		user, err := crudl.CreateUserDataDB(db, createUserData)
		if err != nil {
			log.Println(err)
			http.Error(rw, "Can't create user", 404)
			return
		}

		rw.WriteHeader(201) // 204 - Created
		writeResponseBody(rw, user, "user")
	}
	return CreateUserHandler
}

// @Description  Delete user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      204  {object}  models.UserData
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /user/{id} [delete]
func DeleteUserHandler(db *sql.DB) http.HandlerFunc {
	DeleteUserHandler := func(rw http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		err := crudl.DeleteUserDB(db, id)
		if err != nil {
			log.Println(err)
			http.Error(rw, fmt.Sprintf("Can't delete user id: %s", id), 404)
			return
		}
		rw.WriteHeader(204) // 204 - Success without returning body
	}
	return DeleteUserHandler
}
