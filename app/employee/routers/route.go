package routers

import (
	initialize "Go-Microservice/app/employee/init"
	"Go-Microservice/app/employee/middlewares"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions, http.MethodPut},
		AllowHeaders:     []string{"Content-Type", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/create", middlewares.Auth("Admin"), initialize.EmployeeController.CreateEmployee)
	router.PUT("/update/:id", middlewares.Auth("Admin"), initialize.EmployeeController.UpdateEmployee)
	router.GET("/:id", middlewares.Auth(""), initialize.EmployeeController.FindEmployeeByID)
	router.GET("/", middlewares.Auth("Admin"), initialize.EmployeeController.FindAllEmployees)

	return router
}
