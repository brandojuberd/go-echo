package models

type GetUserFilter struct {
	ID       uint64 `json:"ID"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserLogin struct {
	// Email    string `json:"email" validate:"required,email"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}
