package types

type OrderItemInput struct {
	MenuItemID string `json:"menuItemId"`
	Quantity   int    `json:"quantity"`
}

type CreateOrderInput struct {
	SessionID    string           `json:"sessionId"`
	RestaurantID string           `json:"restaurantId"`
	OrderItems   []OrderItemInput `json:"orderItems"`
}

type OrderItemResponse struct {
	ItemID   string `json:"itemId"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type OrderResponse struct {
	ID        string               `json:"id"`
	TableName string               `json:"tableName"`
	CreatedAt string               `json:"createdAt"`
	Status    string               `json:"status"`
	Items     []OrderItemResponse  `json:"items"`
}