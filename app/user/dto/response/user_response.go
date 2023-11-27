package response

import (
	"Go-Microservice/app/user/models"
	"time"
)

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func NewUserResponse(user models.User) *UserResponse {
	response := &UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return response
}

func NewUserResponses(users []models.User) *[]UserResponse {
	var responses []UserResponse

	for _, response := range users {
		responses = append(responses, *NewUserResponse(response))
	}

	return &responses
}

func NewLoginResponse(token string) *LoginResponse {
	response := &LoginResponse{
		Token: token,
	}

	return response
}
