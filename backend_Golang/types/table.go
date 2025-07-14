package types

type SeatTableInput struct {
	Name         string `form:"name" json:"name" binding:"required"`
	RestaurantID string `json:"restaurantId"`
}

type SeatTableUpdateInput struct {
	IsActive         bool   `json:"isActive"`
	CurrentOccupancy *int   `json:"currentOccupancy,omitempty"`
	RestaurantID     string `json:"restaurantId"`
	TableID          string `json:"tableId"`
}
