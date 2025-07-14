package types

type MenuInput struct {
	Name         string  `form:"name" json:"name" binding:"required"`
	Description  *string `form:"description" json:"description"`
	Price        float64 `form:"price" json:"price" binding:"required"`
	RestaurantID string  `json:"restaurantId"`
	Image        *string `json:"image"`
}
