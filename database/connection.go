package database

import (
	"github.com/matheusferreira165/jwt-authentication/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	conn, err := gorm.Open(postgres.Open("postgres://postgres:develope123@localhost:5432/develope_db"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}
	DB = conn

	conn.AutoMigrate(&models.User{})
}
