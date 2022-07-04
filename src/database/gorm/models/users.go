package models

import (
	"time"
)

type User struct {
	User_ID     uint      `gorm:"primaryKey" json:"user_id"`
	Name        string    `json:"name" valid:"type(string)"`
	Gender      string    `json:"gender" valid:"type(string)"`
	Email       string    `json:"email" valid:"email"`
	Phone       string    `json:"phone" valid:"type(string)"`
	Birth       string    `json:"birth" valid:"type(string)"`
	Address     string    `json:"address" valid:"type(string)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
}

type Users []User
