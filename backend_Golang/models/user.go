package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              string           `gorm:"primaryKey;type:varchar(36)"`
	Username        string           `gorm:"unique;not null"`
	Email           *string          `gorm:"unique"`
	Password        string           `gorm:"not null"`
	Role            GlobalRole       `gorm:"type:varchar(20);default:'CUSTOMER'"`
	RestaurantLinks []RestaurantUser `gorm:"foreignKey:UserID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt 	`gorm:"index"`
}

func (User) TableName() string {
	return "user"
}
