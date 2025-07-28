package model


type Role struct {
    ID   uint   `gorm:"primaryKey;autoIncrement"`
    Name string `gorm:"unique;not null"`
    
    Users []User `gorm:"many2many:user_roles;"`
}

