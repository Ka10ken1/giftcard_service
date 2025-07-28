package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint           `gorm:"primaryKey;autoIncrement"`
	UserID     uint           `gorm:"not null;index"`
	OrderDate  time.Time      `gorm:"not null"`
	Status     OrderStatus    `gorm:"size:50;not null"`
	TotalPrice float64        `gorm:"not null"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	GiftCards []GiftCard `gorm:"foreignKey:OrderID"` // One order to many gift cards
}


