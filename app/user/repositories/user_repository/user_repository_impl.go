package user_repository

import (
	"Go-Microservice/app/user/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *userRepository) FindUserById(id string) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error

	return user, err
}

func (r *userRepository) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error

	return user, err
}

func (r *userRepository) UpdateUser(user models.User, id string) (models.User, error) {
	var userUpdate models.User

	err := r.db.Model(&userUpdate).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return models.User{}, err
	}

	return r.FindUserById(id)
}

func (r *userRepository) DeleteUser(id string) (models.User, error) {
	err := r.db.Where("id = ?", id).Delete(&models.User{}).Error

	return models.User{}, err
}

func (r *userRepository) FindAllUsers() ([]models.User, error) {
	var users []models.User

	err := r.db.Find(&users).Error

	return users, err

}
