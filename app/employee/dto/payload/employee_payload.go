package payload

type EmployeePayload struct {
	UserID     string  `json:"user_id,omitempty"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	JobTitle   string  `json:"job_title"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}
