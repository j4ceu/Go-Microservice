package models

import (
	"Go-Microservice/app/user/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID         uuid.UUID   `gorm:"primaryKey"`
	UserID     uuid.UUID   `gorm:"size:191;foreignKey:ID"`
	User       models.User `gorm:"foreignKey:UserID; references:ID"`
	FirstName  string
	LastName   string
	JobTitle   string
	Department string
	Salary     float64
}

func (employee *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	employee.ID = uuid.New()
	return
}
