package models

import (
	"time"
	"gorm.io/gorm"
)

type Order struct {
	ID           string      `gorm:"primaryKey;type:varchar(36)"`
	TableID      string      `gorm:"not null;index"`
	Table        SeatTable   `gorm:"foreignKey:TableID"`
	RestaurantID string      `gorm:"not null;index"`
	Restaurant   Restaurant  `gorm:"foreignKey:RestaurantID"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderID"`
	Status       OrderStatus `gorm:"type:varchar(20);default:'PENDING'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
