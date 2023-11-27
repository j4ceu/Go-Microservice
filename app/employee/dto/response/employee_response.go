package response

import "Go-Microservice/app/employee/models"

type EmployeeResponse struct {
	UserID     string  `json:"user_id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	JobTitle   string  `json:"job_title"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

func NewEmployeeResponse(employee models.Employee) *EmployeeResponse {
	response := &EmployeeResponse{
		UserID:     employee.ID.String(),
		FirstName:  employee.FirstName,
		LastName:   employee.LastName,
		JobTitle:   employee.JobTitle,
		Department: employee.Department,
		Salary:     employee.Salary,
	}

	return response
}

func NewEmployeeResponses(employees []models.Employee) *[]EmployeeResponse {
	var responses []EmployeeResponse

	for _, employee := range employees {
		responses = append(responses, *NewEmployeeResponse(employee))
	}

	return &responses
}
