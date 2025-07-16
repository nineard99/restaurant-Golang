package models

type RestaurantUser struct {
	ID           string         `gorm:"primaryKey;type:varchar(36)"`
	UserID       string         `gorm:"not null;index"`
	User         User           `gorm:"foreignKey:UserID"`
	RestaurantID string         `gorm:"not null;index"`
	Restaurant   Restaurant     `gorm:"foreignKey:RestaurantID"`
	Role         RestaurantRole `gorm:"type:varchar(20);not null"`
}
