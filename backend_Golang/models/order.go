package models

import (
	"time"
)

type Order struct {
	ID           string      `gorm:"primaryKey;type:varchar(36)" json:"id"`
	TableID      string      `gorm:"not null;index" json:"tableId"`
	Table        SeatTable   `gorm:"foreignKey:TableID" json:"table"`
	RestaurantID string      `gorm:"not null;index" json:"restaurantId"`
	Restaurant   Restaurant  `gorm:"foreignKey:RestaurantID" json:"restaurant"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderID" json:"orderItems"`
	Status       OrderStatus `gorm:"type:varchar(20); default:'PENDING'" json:"status"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}
