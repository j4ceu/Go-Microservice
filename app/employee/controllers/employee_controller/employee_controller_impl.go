package employee_controller

import (
	"Go-Microservice/app/employee/dto/payload"
	"Go-Microservice/app/employee/services/employee_service"
	"Go-Microservice/pkg/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type employeeController struct {
	service employee_service.EmployeeService
}

func NewEmployeeController(service employee_service.EmployeeService) EmployeeController {
	return &employeeController{service: service}
}

func (c *employeeController) CreateEmployee(ctx *gin.Context) {
	var payload payload.EmployeePayload

	if err := ctx.Bind(&payload); err != nil {
		response.Error(ctx, "failed", http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	employee, err := c.service.CreateEmployee(payload)
	if err != nil {
		response.Error(ctx, "failed", http.StatusNoContent, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, employee)

}

func (c *employeeController) UpdateEmployee(ctx *gin.Context) {
	var payload payload.EmployeePayload

	if err := ctx.Bind(&payload); err != nil {
		response.Error(ctx, "failed", http.StatusBadRequest, err)
		ctx.Abort()
		return
	}

	id := ctx.Param("id")

	employee, err := c.service.UpdateEmployee(payload, id)
	if err != nil {
		response.Error(ctx, "failed", http.StatusInternalServerError, err)
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, employee)

}
func (c *employeeController) FindEmployeeByID(ctx *gin.Context) {
	id := ctx.Param("id")

	employee, err := c.service.FindEmployeeByID(id)
	if err != nil {
		if err.Error() == "204" {
			response.Error(ctx, "failed", http.StatusNoContent, errors.New("employee not found"))
		} else {
			response.Error(ctx, "failed", http.StatusInternalServerError, err)
		}
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, employee)
}

func (c *employeeController) FindAllEmployees(ctx *gin.Context) {
	employees, err := c.service.FindAllEmployees()
	if err != nil {
		if err.Error() == "204" {
			response.Error(ctx, "failed", http.StatusNoContent, errors.New("employee not found"))
		} else {
			response.Error(ctx, "failed", http.StatusInternalServerError, err)
		}
		ctx.Abort()
		return
	}

	response.Success(ctx, "success", http.StatusOK, employees)

}
