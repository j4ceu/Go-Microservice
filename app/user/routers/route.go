package routers

import (
	"Go-Microservice/app/user/constant"
	initializers "Go-Microservice/app/user/init"
	"Go-Microservice/app/user/middlewares"
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

	router.POST("/register", initializers.UserController.CreateUser)
	router.POST("/register/admin", initializers.UserController.CreateAdmin)
	router.POST("/login", initializers.UserController.Login)
	router.PUT("/update", middlewares.Auth(""), initializers.UserController.UpdateUser)
	router.GET("/:email", middlewares.Auth(""), initializers.UserController.FindUserByEmail)
	router.PUT("/change-password", middlewares.Auth(""), initializers.UserController.ChangePassword)
	router.GET("/all", middlewares.Auth(constant.UserRoleAdmin.String()), initializers.UserController.FindAllUsers)

	return router
}
