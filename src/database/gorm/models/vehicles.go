package models

import (
	"time"
)

type Vehicle struct {
	Vehicle_ID  uint      `gorm:"primaryKey" json:"vehicle_id"`
	Name        string    `json:"name" valid:"type(string), required"`
	Location    string    `json:"location" valid:"type(string), required"`
	Description string    `json:"description" valid:"type(string), required"`
	Price       string    `json:"price" valid:"type(string), required"`
	Status      string    `json:"status" valid:"type(string), required"`
	Stock       int       `json:"stock" valid:"type(int), required"`
	Category    string    `json:"category" valid:"type(string), required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Image       string    `json:"image" valid:"type(string)"`
	Rating      int       `json:"rating" valid:"type(int), required"`
}

type Vehicles []Vehicle
