package vehicles

import (
	"time"
)

type Vehicle struct {
	Vehicle_ID  uint      `gorm:"primaryKey" json:"vehicle_id"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	Status      string    `json:"status"`
	Stock       int       `json:"stock"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Image       string    `json:"image"`
	Rating      int       `json:"rating"`
}

type Vehicles []Vehicle
