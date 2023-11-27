package user_repository

import "Go-Microservice/app/user/models"

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	FindUserById(id string) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	UpdateUser(user models.User, id string) (models.User, error)
	DeleteUser(id string) (models.User, error)
	FindAllUsers() ([]models.User, error)
}
