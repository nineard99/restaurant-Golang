package models

import (
	"time"
)

type OrderItem struct {
	ID         string         `gorm:"primaryKey;type:varchar(36)" json:"id"`
	OrderID    string         `gorm:"not null;index" json:"orderId"`
	Order      Order          `gorm:"foreignKey:OrderID" json:"order"`
	MenuItemID string         `gorm:"not null;index" json:"menuItemId"`
	MenuItem   MenuItem       `gorm:"foreignKey:MenuItemID" json:"menuItem"`
	Quantity   int            `gorm:"not null" json:"quantity"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
}

