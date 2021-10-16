package models

import "bootcamp_demo/core/models"

type User struct {
	models.SqlModel
	MobileNumber string `gorm:"unique" json:"mobile_number"`
	Email        string `json:"email"`
}
