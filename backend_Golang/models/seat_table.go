package models

import (
	"time"

	"gorm.io/gorm"
)

type SeatTable struct {
	ID               string         `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Name             string         `gorm:"not null" json:"name"`
	QRCode           *string        `json:"qrCode,omitempty"`
	RestaurantID     string         `gorm:"not null;index" json:"restaurantId"`
	Restaurant       Restaurant     `gorm:"foreignKey:RestaurantID" json:"-"`
	Orders           []Order        `gorm:"foreignKey:TableID" json:"-"`
	SessionID        *string        `json:"sessionId,omitempty"`
	IsActive         bool           `gorm:"default:false" json:"isActive"`
	CurrentOccupancy *int           `json:"currentOccupancy,omitempty"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
