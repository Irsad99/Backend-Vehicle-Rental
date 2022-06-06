package histories

import (
	"time"
)

type History struct {
	History_Id uint      `gorm:"primaryKey" json:"history_id"`
	Id_User    uint      `gorm:"foreignKey" json:"id_user"`
	Id_Vehicle uint      `gorm:"foreignKey" json:"id_vehicle"`
	Start_Date string    `json:"start_date"`
	End_Date   string    `json:"end_date"`
	Prepayment string    `json:"prepayment"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Result struct {
	Users      string `json:"users"`
	Vehicle    string `json:"vehicle"`
	Start_Date string `json:"start_date"`
	End_Date   string `json:"end_date"`
	Prepayment string `json:"prepayment"`
	Status     string `json:"status"`
}

type Results []Result
type Histories []History
