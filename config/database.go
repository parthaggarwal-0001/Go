package config

import (
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

func InitDB(cfg Config) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    DB = db
    log.Println("Database connected successfully")
}

func AutoMigrate() {
    // Models will be migrated here
}