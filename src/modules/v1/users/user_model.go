package users

import (
	"time"
)

type User struct {
	User_ID   uint      `gorm:"primaryKey" json:"user_id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Email     *string    `json:"email"`
	Phone     string    `json:"phone"`
	Birth     string    `json:"birth"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User
