package models

import (
	"bootcamp_demo/config"
	"gorm.io/gorm"
	"time"
)

type SqlModel struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *SqlModel) Db() *gorm.DB {
	return config.MySqlDb()
}
