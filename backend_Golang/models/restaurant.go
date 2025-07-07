package models

import (
	"time"
	"gorm.io/gorm"
)

type Restaurant struct {
	ID        string           `gorm:"primaryKey;type:varchar(36)"`
	Name      string           `gorm:"not null"`
	Address   *string
	Image     *string
	Users     []RestaurantUser `gorm:"foreignKey:RestaurantID"`
	Tables    []SeatTable      `gorm:"foreignKey:RestaurantID"`
	MenuItems []MenuItem       `gorm:"foreignKey:RestaurantID"`
	Orders    []Order          `gorm:"foreignKey:RestaurantID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
