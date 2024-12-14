package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint64 `json:"ID"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"password"`
	// CreatedAt string `json:"password"`
}