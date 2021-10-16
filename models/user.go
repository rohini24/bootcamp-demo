package models

import (
	"bootcamp_demo/core/models"
	"gorm.io/gorm"
)

type User struct {
	models.SqlModel
	MobileNumber string `gorm:"unique" json:"mobile_number"`
	Email        string `json:"email"`
}

func (u *User) GetById(id int) (*gorm.DB, User) {
	var model User
	result := u.Db().First(&model, id)
	return result, model
}
