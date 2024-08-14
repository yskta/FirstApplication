package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `json:"-"` // パスワードはJSONレスポンスには含めない
	Projects []Project `gorm:"foreignKey:UserID" json:"projects"`
}