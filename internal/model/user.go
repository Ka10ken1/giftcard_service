package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	FirstName string         `gorm:"size:100;not null"`
	LastName  string         `gorm:"size:100;not null"`
	Email     string         `gorm:"size:100;uniqueIndex;not null"`
	Password  string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"` // soft delete

	Orders []Order `gorm:"foreignKey:UserID"` // One user to many orders
}

