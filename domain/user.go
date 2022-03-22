package domain

import (
	"time"
)

type User struct {
	ID        int        `json:"id,omitempty"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type SignupRequest struct {
	FirstName string `json:"first_name" binding:"required,max=64"`
	LastName  string `json:"last_name" binding:"required,max=64"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required,password"`
}

type SigninRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,password"`
}

type UpdateProfileRequest struct {
	FirstName string `json:"first_name" binding:"required,max=64"`
	LastName  string `json:"last_name" binding:"required,max=64"`
}
