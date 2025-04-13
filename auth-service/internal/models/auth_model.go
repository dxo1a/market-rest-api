package models

type LoginInput struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"fullName" binding:"required"`
	Address  string `json:"address" binding:"required"`
}
