package models

import (
	"time"
)

type MenuItem struct {
	ID           string      `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Name         string      `gorm:"not null" json:"name"`
	Description  *string     `json:"description"`
	Price        float64     `gorm:"not null" json:"price"`
	Image        *string     `json:"image"`
	RestaurantID string      `gorm:"not null;index" json:"restaurantId"`
	Restaurant   Restaurant  `gorm:"foreignKey:RestaurantID" json:"-"`
	OrderItems   []OrderItem `gorm:"foreignKey:MenuItemID" json:"-"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}
