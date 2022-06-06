package histories

import (
	"time"
)

type History struct {
	History_Id uint      `gorm:"primaryKey" json:"history_id"`
	Id_User    uint      `gorm:"foreignKey" json:"id_user" valid:"type(uint), required"`
	Id_Vehicle uint      `gorm:"foreignKey" json:"id_vehicle" valid:"type(uint), required"`
	Start_Date string    `json:"start_date" valid:"type(string), required"`
	End_Date   string    `json:"end_date" valid:"type(string), required"`
	Prepayment string    `json:"prepayment" valid:"type(string), required"`
	Status     string    `json:"status" valid:"type(string), required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Quantity   string    `json:"quantity" valid:"type(string), required"`
}

type Result struct {
	Users      string `json:"users"`
	Vehicle    string `json:"vehicle"`
	Start_Date string `json:"start_date"`
	End_Date   string `json:"end_date"`
	Prepayment string `json:"prepayment"`
	Status     string `json:"status"`
	Quantity   string `json:"quantity"`
}

type Results []Result
type Histories []History
