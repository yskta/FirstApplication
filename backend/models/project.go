package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"userId"`
	Tasks       []Task `gorm:"foreignKey:ProjectID" json:"tasks"`
}