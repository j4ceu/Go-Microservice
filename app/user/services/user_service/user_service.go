package user_service

import (
	"Go-Microservice/app/user/dto/payload"
	"Go-Microservice/app/user/dto/response"
)

type UserService interface {
	// User
	CreateUser(payload payload.UserPayload) (*response.UserResponse, error)
	CreateAdmin(payload payload.UserPayload) (*response.UserResponse, error)
	ChangePassword(payload payload.ChangePasswordPayload, userId string) (*response.UserResponse, error)
	UpdateUser(payload payload.UpdateUserPayload, userId string) (*response.UserResponse, error)
	Login(payload payload.UserPayloadLogin) (*response.LoginResponse, error)
	FindUserByEmail(email string) (*response.UserResponse, error)

	//Admin
	FindAllUsers() (*[]response.UserResponse, error)
}
