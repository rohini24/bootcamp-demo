package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func MySqlDb() *gorm.DB {
	username := os.Getenv("DB_MYSQL_USER")
	password := os.Getenv("DB_MYSQL_PASSWORD")
	database := os.Getenv("DB_MYSQL_DATABASE")
	host := os.Getenv("DB_MYSQL_HOST")
	port := os.Getenv("DB_MYSQL_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect MYSQL: " + err.Error())
	}

	return conn
}
