package employee_controller

import "github.com/gin-gonic/gin"

type EmployeeController interface {
	CreateEmployee(ctx *gin.Context)
	UpdateEmployee(ctx *gin.Context)
	FindEmployeeByID(ctx *gin.Context)
	FindAllEmployees(ctx *gin.Context)
}
