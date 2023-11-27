package init

import (
	"Go-Microservice/app/employee/controllers/employee_controller"
	"Go-Microservice/app/employee/repositories/employee_repository"
	"Go-Microservice/app/employee/services/employee_service"
	"log"
	"time"

	"github.com/joho/godotenv"
)

var employeeRepository employee_repository.EmployeeRepository
var employeeService employee_service.EmployeeService
var EmployeeController employee_controller.EmployeeController

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Println(err.Error())
	}
	time.Local = location
	initRepositories()
	initServices()
	initController()

}

func initRepositories() {
	db := connectDatabase()
	DB = db

	employeeRepository = employee_repository.NewEmployeeRepository(db)

}

func initServices() {
	employeeService = employee_service.NewEmployeeService(employeeRepository)

}

func initController() {
	EmployeeController = employee_controller.NewEmployeeController(employeeService)

}
