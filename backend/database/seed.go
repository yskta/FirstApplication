package database

import (
	"log"

	"github.com/yskta/taskmaster/backend/models"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) error {
	// ユーザーの作成
	user := models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123", // 実際のアプリケーションではハッシュ化すべき
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	// プロジェクトの作成
	project := models.Project{
		Name:        "Test Project",
		Description: "This is a test project",
		UserID:      user.ID,
	}
	if err := db.Create(&project).Error; err != nil {
		return err
	}

	// タスクの作成
	tasks := []models.Task{
		{Title: "Task 1", Description: "Description for task 1", Status: "TODO", ProjectID: project.ID},
		{Title: "Task 2", Description: "Description for task 2", Status: "IN_PROGRESS", ProjectID: project.ID},
		{Title: "Task 3", Description: "Description for task 3", Status: "DONE", ProjectID: project.ID},
	}
	for _, task := range tasks {
		if err := db.Create(&task).Error; err != nil {
			return err
		}
	}

	log.Println("Seed data inserted successfully")
	return nil
}