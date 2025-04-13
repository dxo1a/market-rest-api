package repository

import (
	"auth/config"
	"auth/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(cfg config.PostgresConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=%s",
		cfg.Host, cfg.User, cfg.DBName, cfg.Password, cfg.Port, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к Postgres")
	} else {
		log.Print("Подключено к Postgres.")
	}

	db.AutoMigrate(&models.User{})
	return db
}
