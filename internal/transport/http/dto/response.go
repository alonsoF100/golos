package dto

import (
	"time"

	"github.com/alonsoF100/golos/internal/models"
)

type ErrorResponse struct {
	Error     string    `json:"error"`
	TimeStamp time.Time `json:"timestamp"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error:     err.Error(),
		TimeStamp: time.Now(),
	}
}

type UserResponse struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Nickname:  user.Nickname,
		Password:  user.Password,
		CreatedAt: user.UpdatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type UsersResponse struct {
	users []*UserResponse
}

func NewUsersResponse(users []*models.User) UsersResponse {
	responseUsers := UsersResponse{
		users: make([]*UserResponse, 0, len(users)),
	}

	for _, user := range users {
		temp := &UserResponse{
			ID:        user.ID,
			Nickname:  user.Nickname,
			Password:  user.Password,
			CreatedAt: user.UpdatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		responseUsers.users = append(responseUsers.users, temp)
	}

	return responseUsers
}
