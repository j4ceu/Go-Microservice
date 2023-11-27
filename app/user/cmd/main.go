package main

import (
	initializers "Go-Microservice/app/user/init"
	"Go-Microservice/app/user/routers"
	"log"
	"os"
)

func main() {
	initializers.Setup()
	router := routers.RouterSetup()

	log.Println(":" + os.Getenv("USER_SERVICE_PORT"))
	router.Run(":" + os.Getenv("USER_SERVICE_PORT"))
}
