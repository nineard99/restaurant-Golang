package models

import (
	"time"
	"gorm.io/gorm"
)

type MenuItem struct {
	ID           string      `gorm:"primaryKey;type:varchar(36)"`
	Name         string      `gorm:"not null"`
	Description  *string
	Price        float64     `gorm:"not null"`
	Image        *string
	RestaurantID string      `gorm:"not null;index"`
	Restaurant   Restaurant  `gorm:"foreignKey:RestaurantID"`
	OrderItems   []OrderItem `gorm:"foreignKey:MenuItemID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
