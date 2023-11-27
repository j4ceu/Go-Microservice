package user_controller

import "github.com/gin-gonic/gin"

type UserController interface {
	CreateUser(ctx *gin.Context)
	CreateAdmin(ctx *gin.Context)
	Login(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	FindUserByEmail(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
	FindAllUsers(ctx *gin.Context)
}
