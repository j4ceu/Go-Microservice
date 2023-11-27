package employee_repository

import "Go-Microservice/app/employee/models"

type EmployeeRepository interface {
	CreateEmployee(employee models.Employee) (models.Employee, error)
	UpdateEmployee(employee models.Employee, id string) (models.Employee, error)
	FindEmployeeByID(id string) (models.Employee, error)
	FindAllEmployees() ([]models.Employee, error)
}
