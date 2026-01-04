package models

import (
	"github.com/google/uuid"
	"time"
)

type UserData struct {
	ID        uuid.UUID `json:"user_data"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type UserDataList struct {
	UserList []UserData `json:"user_list"`
}

// Request structs
type CreateUserDataRequest struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UpdateUserDataRequest struct {
	Name     *string `json:"name"`
	Login    *string `json:"login"`
	Password *string `json:"password"`
}
