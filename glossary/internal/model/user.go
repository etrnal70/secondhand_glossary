package model

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserRole  string    `json:"user_role" gorm:"default:user"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique; not null"`
  Password  string    `json:"password" gorm:"not null; size:64"`
	Reviews   []*Review `json:"reviews"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
