package main

import (
	initializers "Go-Microservice/app/employee/init"
	"Go-Microservice/app/employee/routers"
	"log"
	"os"
)

func main() {
	initializers.Setup()
	router := routers.RouterSetup()

	log.Println(":" + os.Getenv("EMPLOYEE_SERVICE_PORT"))
	router.Run(":" + os.Getenv("EMPLOYEE_SERVICE_PORT"))
}
