package database

import (
	"log"
	"os"
	"fmt"
	"github.com/yskta/taskmaster/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=your_password dbname=taskmaster port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// マイグレーションの実行
	err = DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Task{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully")
    return DB, nil
}