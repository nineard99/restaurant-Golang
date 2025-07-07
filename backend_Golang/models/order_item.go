package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderItem struct {
	ID         string   `gorm:"primaryKey;type:varchar(36)"`
	OrderID    string   `gorm:"not null;index"`
	Order      Order    `gorm:"foreignKey:OrderID"`
	MenuItemID string   `gorm:"not null;index"`
	MenuItem   MenuItem `gorm:"foreignKey:MenuItemID"`
	Quantity   int      `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
