package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Status      string `gorm:"not null" json:"status"`
	ProjectID   uint   `json:"projectId"`
}