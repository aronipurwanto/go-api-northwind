package models

type LoginRequest struct {
	Username string `json:"username" validate:"required"` // bisa username atau email
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username   string `json:"username" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6"`
	Role       string `json:"role"`
	EmployeeID *int   `json:"employee_id"` // opsional
}
