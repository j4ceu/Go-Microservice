package employee_service

import (
	"Go-Microservice/app/employee/dto/payload"
	"Go-Microservice/app/employee/dto/response"
	"Go-Microservice/app/employee/models"
	"Go-Microservice/app/employee/repositories/employee_repository"
	"errors"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type employeeService struct {
	repository employee_repository.EmployeeRepository
}

func NewEmployeeService(repository employee_repository.EmployeeRepository) EmployeeService {
	return &employeeService{repository: repository}
}

func (s *employeeService) CreateEmployee(payload payload.EmployeePayload) (*response.EmployeeResponse, error) {
	employee := models.Employee{
		UserID:     uuid.MustParse(payload.UserID),
		FirstName:  payload.FirstName,
		LastName:   payload.LastName,
		JobTitle:   payload.JobTitle,
		Department: payload.Department,
		Salary:     payload.Salary,
	}

	employee, err := s.repository.CreateEmployee(employee)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewEmployeeResponse(employee), nil

}

func (s *employeeService) UpdateEmployee(payload payload.EmployeePayload, id string) (*response.EmployeeResponse, error) {
	employee, err := s.repository.FindEmployeeByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("204")
		}
		return nil, err
	}

	if payload.FirstName != "" {
		employee.FirstName = payload.FirstName
	}
	if payload.LastName != "" {
		employee.LastName = payload.LastName
	}
	if payload.JobTitle != "" {
		employee.JobTitle = payload.JobTitle
	}
	if payload.Department != "" {
		employee.Department = payload.Department
	}
	if payload.Salary != 0 {
		employee.Salary = payload.Salary
	}

	employee, err = s.repository.UpdateEmployee(employee, id)
	if err != nil {
		log.Println(string("\033[31m"), err.Error())
		return nil, err
	}

	return response.NewEmployeeResponse(employee), nil

}

func (s *employeeService) FindEmployeeByID(id string) (*response.EmployeeResponse, error) {
	employee, err := s.repository.FindEmployeeByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("204")
		}
		return nil, err
	}

	return response.NewEmployeeResponse(employee), nil

}

func (s *employeeService) FindAllEmployees() (*[]response.EmployeeResponse, error) {
	employee, err := s.repository.FindAllEmployees()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("204")
		}
		return nil, err
	}

	return response.NewEmployeeResponses(employee), nil

}
