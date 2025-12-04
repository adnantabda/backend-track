package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status" gorm:"default:'pending'"`
	UserID      uint   `json:"user_id" gorm:"not null"`
	User        User   `json:"-" gorm:"foreignKey:UserID"`
}
