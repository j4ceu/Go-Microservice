package init

import (
	"Go-Microservice/app/user/controllers/user_controller"
	"Go-Microservice/app/user/repositories/user_repository"
	"Go-Microservice/app/user/services/user_service"
	"log"
	"time"

	"github.com/joho/godotenv"
)

//User
var UserController user_controller.UserController
var userService user_service.UserService
var userRepository user_repository.UserRepository

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

	userRepository = user_repository.NewUserRepository(db)

}

func initServices() {
	userService = user_service.NewUserService(userRepository)

}

func initController() {
	UserController = user_controller.NewUserController(userService)

}
