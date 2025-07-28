package config

import (
	"fmt"
	"log"
	"os"
	"github.com/Ka10ken1/giftcard_service/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB


func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_NAME", "giftcard"),
		getEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&model.User{} , &model.Role{}, &model.Order{}, &model.GiftCardType{}, &model.GiftCard{})
	if err != nil {
		log.Fatalf("Failed to auto migrate models: %v", err)
	}

	DB = db
	log.Println("Database connection established")
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
