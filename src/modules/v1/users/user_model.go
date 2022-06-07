package users

import (
	"time"
)

type User struct {
	User_ID   uint      `gorm:"primaryKey" json:"user_id"`
	Name      string    `json:"name" valid:"type(string), required"`
	Gender    string    `json:"gender" valid:"type(string), required"`
	Email     *string   `json:"email" valid:"email, required"`
	Phone     string    `json:"phone" valid:"type(string), required"`
	Birth     string    `json:"birth" valid:"type(string), required"`
	Address   string    `json:"address" valid:"type(string), required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User
