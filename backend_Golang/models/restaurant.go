package models

import (
	"time"
)

type Restaurant struct {
	ID        string           `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Name      string           `gorm:"not null" json:"name"`
	Address   *string          `json:"address,omitempty"`
	Image     *string          `json:"image,omitempty"`
	Users     []RestaurantUser `gorm:"foreignKey:RestaurantID" json:"-"`
	Tables    []SeatTable      `gorm:"foreignKey:RestaurantID" json:"-"`
	MenuItems []MenuItem       `gorm:"foreignKey:RestaurantID" json:"-"`
	Orders    []Order          `gorm:"foreignKey:RestaurantID" json:"-"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}
