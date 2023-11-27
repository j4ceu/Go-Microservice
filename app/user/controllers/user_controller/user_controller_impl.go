package user_controller

import (
	"Go-Microservice/app/user/dto/payload"
	"Go-Microservice/app/user/services/user_service"
	"Go-Microservice/pkg/response"
	"Go-Microservice/pkg/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service user_service.UserService
}

func NewUserController(service user_service.UserService) UserController {
	return &userController{service: service}
}

func (c *userController) CreateUser(ctx *gin.Context) {
	var payload payload.UserPayload

	if err := ctx.Bind(&payload); err != nil {
		response.Error(ctx, "failed", http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	user, err := c.service.CreateUser(payload)
	if err != nil {
		response.Error(ctx, "failed", http.StatusNoContent, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, user)

}

func (c *userController) CreateAdmin(ctx *gin.Context) {
	var payload payload.UserPayload

	if err := ctx.Bind(&payload); err != nil {
		response.Error(ctx, "failed", http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	user, err := c.service.CreateAdmin(payload)
	if err != nil {
		response.Error(ctx, "failed", http.StatusNoContent, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, user)

}

func (c *userController) Login(ctx *gin.Context) {
	var payload payload.UserPayloadLogin

	if err := ctx.Bind(&payload); err != nil {
		response.Error(ctx, "failed", http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	token, err := c.service.Login(payload)
	if err != nil {
		if err.Error() == "401" {
			response.Error(ctx, "failed", http.StatusUnauthorized, errors.New("wrong password"))
		} else if err.Error() == "404" {
			response.Error(ctx, "failed", http.StatusNotFound, errors.New("user not found"))
		} else {
			response.Error(ctx, "failed", http.StatusInternalServerError, err)
		}
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, token)

}

func (c *userController) UpdateUser(ctx *gin.Context) {
	var payload payload.UpdateUserPayload

	if err := ctx.Bind(&payload); err != nil {
		response.Error(ctx, "failed", http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	claims := utils.GetTokenClaims(ctx)

	user, err := c.service.UpdateUser(payload, claims.UserID)
	if err != nil {
		response.Error(ctx, "failed", http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, user)

}

func (c *userController) FindUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	user, err := c.service.FindUserByEmail(email)
	if err != nil {
		if err.Error() == "204" {
			response.Error(ctx, "failed", http.StatusNoContent, errors.New("user not found"))
		} else {
			response.Error(ctx, "failed", http.StatusInternalServerError, err)
		}
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, user)
}

func (c *userController) ChangePassword(ctx *gin.Context) {
	var payload payload.ChangePasswordPayload

	if err := ctx.Bind(&payload); err != nil {
		response.Error(ctx, "failed", http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	claims := utils.GetTokenClaims(ctx)

	user, err := c.service.ChangePassword(payload, claims.UserID)
	if err != nil {
		if err.Error() == "403" {
			response.Error(ctx, "failed", http.StatusForbidden, errors.New("wrong old password"))
		} else if err.Error() == "500" {
			response.Error(ctx, "failed", http.StatusBadRequest, errors.New("confirm password not match or old password similar to new password"))
		} else {
			response.Error(ctx, "failed", http.StatusInternalServerError, err)
		}
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, user)

}

func (c *userController) FindAllUsers(ctx *gin.Context) {
	users, err := c.service.FindAllUsers()
	if err != nil {
		if err.Error() == "204" {
			response.Error(ctx, "failed", http.StatusNoContent, errors.New("user not found"))
		} else {
			response.Error(ctx, "failed", http.StatusInternalServerError, err)
		}
		ctx.Abort()
		return

	}

	response.Success(ctx, "success", http.StatusOK, users)
}
