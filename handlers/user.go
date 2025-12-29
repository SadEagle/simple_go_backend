package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"movie_backend_go/db/crudl"
	"movie_backend_go/models"

	"net/http"
)

// TODO: check is it correct signature `rw` interface and not pointer-like... It's weird
func writeResponseBody(rw http.ResponseWriter, user models.UserData) {
	response_user := user.ToResponse()
	response_user_byte, err := json.Marshal(response_user)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Can't convert user data to response json", 500)
	}

	rw.Header().Set("Content-Type", "application/json")
	_, err = rw.Write(response_user_byte)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Can't send user data", 500)
	}
}

// @Summary      Show user
// @Description  Get user by id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.UserDataResponse
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
		writeResponseBody(rw, user)

	}
	return GetUserHandler
}

// @Summary      Update user
// @Description  Update user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Param        request 		body	models.UpdateUserDataRequest  true  "User creation data"
// @Success      200  {object}  models.UserDataResponse
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /user/{id} [PATCH]
// @Router       /user/{id} [PUT]
func UpdateUserHandler(db *sql.DB) http.HandlerFunc {
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
		id := r.PathValue("id")

		user, err := crudl.UpdateUserDataDB(db, updateUserdata, id)
		if err != nil {
			log.Println(err)
			http.Error(rw, "Can't update user", 404)
			return
		}

		writeResponseBody(rw, user)
	}
	return UpdateUserHandler
}

// @Summary      Create user
// @Description  Create user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request 		body	models.CreateUserDataRequest  true  "User creation data"
// @Success      201  {object}  models.UserDataResponse
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /user [post]
func CreateUserHandler(db *sql.DB) http.HandlerFunc {
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
		writeResponseBody(rw, user)
	}
	return CreateUserHandler
}

// @Description  Delete user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      204  {object}  models.UserDataResponse
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
