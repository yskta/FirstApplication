package database

import (
	"log"
	"os"
	"fmt"
	"time"
	"github.com/yskta/taskmaster/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	fmt.Println(dsn)
	if dsn == "" {
		dsn = "host=db user=postgres password=password dbname=taskmaster port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	}

	var err error
	var db *gorm.DB

	// リトライの設定
	maxRetries := 5
	retryInterval := time.Second * 5

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			log.Printf("Retrying in %v...", retryInterval)
			time.Sleep(retryInterval)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %v", maxRetries, err)
	}

	DB = db

	// マイグレーションの実行
	err = DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Task{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully")
    return DB, nil
}