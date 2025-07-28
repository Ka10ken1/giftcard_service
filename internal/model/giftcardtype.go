package model

type GiftCardType struct {
	ID       uint    `gorm:"primaryKey;autoIncrement"`
	Name     string  `gorm:"size:100;not null;uniqueIndex"`
	Currency string  `gorm:"size:10;not null"`

	GiftCards []GiftCard `gorm:"foreignKey:GiftCardTypeID"` // One gift card type to many gift cards
}

