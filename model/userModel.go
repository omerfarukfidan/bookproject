package model

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	PhoneNumber string         `json:"phone_number"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
