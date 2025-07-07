package models

import (
	"time"
	"gorm.io/gorm"
)

type SeatTable struct {
	ID               string       `gorm:"primaryKey;type:varchar(36)"`
	Name             string       `gorm:"not null"`
	QRCode           *string
	RestaurantID     string       `gorm:"not null;index"`
	Restaurant       Restaurant   `gorm:"foreignKey:RestaurantID"`
	Orders           []Order      `gorm:"foreignKey:TableID"`
	SessionID        *string
	IsActive         bool         `gorm:"default:false"`
	CurrentOccupancy *int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
