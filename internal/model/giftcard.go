package model

import (
	"time"

	"github.com/google/uuid"
)


type GiftCard struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Code           string    `gorm:"uniqueIndex;not null"`
	Balance        float64   `gorm:"not null"`
	ExpirationDate time.Time `gorm:"not null"`
	Redeemed       bool      `gorm:"not null;default:false"`
	CreationDate   time.Time `gorm:"autoCreateTime"`
	GiftCardTypeID uint      `gorm:"not null"`
	OrderID        uint      `gorm:"not null"`

	GiftCardType GiftCardType `gorm:"foreignKey:GiftCardTypeID"`
	Order        Order        `gorm:"foreignKey:OrderID"`
}

