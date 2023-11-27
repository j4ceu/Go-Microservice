package user_service

import (
	"Go-Microservice/app/user/constant"
	"Go-Microservice/app/user/dto/payload"
	"Go-Microservice/app/user/dto/response"
	"Go-Microservice/app/user/models"
	"Go-Microservice/app/user/repositories/user_repository"
	"Go-Microservice/pkg/utils"
	"errors"
	"log"

	"gorm.io/gorm"
)

type userService struct {
	repository user_repository.UserRepository
}

func NewUserService(userRepository user_repository.UserRepository) UserService {
	return &userService{repository: userRepository}
}

func (s *userService) CreateUser(payload payload.UserPayload) (*response.UserResponse, error) {
	user := models.User{
		Email:    payload.Email,
		Password: payload.Password,
		Role:     constant.UserRoleUser,
	}

	if err := user.HashPassword(user.Password); err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	user, err := s.repository.CreateUser(user)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewUserResponse(user), nil

}

func (s *userService) CreateAdmin(payload payload.UserPayload) (*response.UserResponse, error) {
	user := models.User{
		Email:    payload.Email,
		Password: payload.Password,
		Role:     constant.UserRoleAdmin,
	}

	if err := user.HashPassword(user.Password); err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	user, err := s.repository.CreateUser(user)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewUserResponse(user), nil

}

func (s *userService) UpdateUser(payload payload.UpdateUserPayload, userId string) (*response.UserResponse, error) {
	user, err := s.repository.FindUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("204")
		}
		return nil, err
	}

	if payload.Email != "" {
		user.Email = payload.Email
	}

	// Update User
	user, err = s.repository.UpdateUser(user, userId)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewUserResponse(user), nil

}

func (s *userService) Login(payload payload.UserPayloadLogin) (*response.LoginResponse, error) {
	user, err := s.repository.FindUserByEmail(payload.Email)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, errors.New("404")
	}

	if err := user.CheckPassword(payload.Password); err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, errors.New("401")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewLoginResponse(token), nil
}

func (s *userService) ChangePassword(payload payload.ChangePasswordPayload, userId string) (*response.UserResponse, error) {
	user, err := s.repository.FindUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("204")
		}
		return nil, err
	}

	err = user.CheckPassword(payload.OldPassword)
	if err != nil {
		log.Println(string("\033[31m"), errors.New("403"))
		return nil, err
	}

	if payload.NewPassword != payload.ConfirmPassword || payload.OldPassword == payload.NewPassword {
		return nil, errors.New("500")
	}

	user.Password = payload.ConfirmPassword

	if err := user.HashPassword(user.Password); err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	user, err = s.repository.UpdateUser(user, userId)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewUserResponse(user), nil

}

func (s *userService) FindUserByEmail(email string) (*response.UserResponse, error) {
	user, err := s.repository.FindUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("204")
		}
		return nil, err
	}

	return response.NewUserResponse(user), nil
}

func (s *userService) FindAllUsers() (*[]response.UserResponse, error) {
	users, err := s.repository.FindAllUsers()
	if err != nil {
		return nil, err
	}

	return response.NewUserResponses(users), nil
}
