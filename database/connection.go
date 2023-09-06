package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/matheusferreira165/jwt-authentication/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Recupera as variáveis de ambiente
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Constrói a string de conexão com o banco de dados
	dns := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPassword)

	conn, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}
	DB = conn

	conn.AutoMigrate(&models.User{})
}
