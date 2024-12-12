package dto

type GetUserFilter struct {
	ID       uint64 `json:"ID"`
	Email    string `json:"email"`
	Username string `json:"username"`
}