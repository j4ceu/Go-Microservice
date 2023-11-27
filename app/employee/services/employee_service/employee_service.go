package employee_service

import (
	"Go-Microservice/app/employee/dto/payload"
	"Go-Microservice/app/employee/dto/response"
)

type EmployeeService interface {
	CreateEmployee(payload payload.EmployeePayload) (*response.EmployeeResponse, error)
	UpdateEmployee(payload payload.EmployeePayload, id string) (*response.EmployeeResponse, error)
	FindEmployeeByID(id string) (*response.EmployeeResponse, error)
	FindAllEmployees() (*[]response.EmployeeResponse, error)
}
