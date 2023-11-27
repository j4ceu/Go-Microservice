package employee_repository

import (
	"Go-Microservice/app/employee/models"

	"gorm.io/gorm"
)

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) CreateEmployee(employee models.Employee) (models.Employee, error) {
	err := r.db.Create(&employee).Error

	return employee, err
}

func (r *employeeRepository) FindEmployeeByID(id string) (models.Employee, error) {
	var employee models.Employee
	err := r.db.Where("id = ?", id).First(&employee).Error

	return employee, err

}

func (r *employeeRepository) UpdateEmployee(employee models.Employee, id string) (models.Employee, error) {
	var employeeUpdate models.Employee

	err := r.db.Model(&employeeUpdate).Where("id = ?", id).Updates(employee).Error
	if err != nil {
		return models.Employee{}, err
	}

	return r.FindEmployeeByID(id)
}

func (r *employeeRepository) FindAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee

	err := r.db.Find(&employees).Error

	return employees, err
}
