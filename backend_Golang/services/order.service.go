package services

import (
	"errors"
	"strings"
	"time"

	"github.com/nineard99/restaurant-Golang/config"
	"github.com/nineard99/restaurant-Golang/models"
	"github.com/nineard99/restaurant-Golang/types"
	"github.com/nineard99/restaurant-Golang/utils"
	"gorm.io/gorm"
)

func CreateOrder(input types.CreateOrderInput) (*models.Order, error) {
	if strings.TrimSpace(input.RestaurantID) == "" {
		return nil, errors.New("restaurant ID is required")
	}
	if strings.TrimSpace(input.SessionID) == "" {
		return nil, errors.New("session ID is required")
	}
	if len(input.OrderItems) == 0 {
		return nil, errors.New("order items are required")
	}

	var restaurant models.Restaurant
	if err := config.DB.First(&restaurant, "id = ?", input.RestaurantID).Error; err != nil {
		return nil, errors.New("restaurant not found")
	}

	var seatTable models.SeatTable
	if err := config.DB.First(&seatTable, "session_id = ?", input.SessionID).Error; err != nil {
		return nil, errors.New("session not found")
	}

	order := models.Order{
		ID:           utils.GenerateUUID(),
		TableID:      seatTable.ID,
		RestaurantID: input.RestaurantID,
		Status:       "PENDING",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		for _, item := range input.OrderItems {
			orderItem := models.OrderItem{
				ID:         utils.GenerateUUID(),
				OrderID:    order.ID,
				MenuItemID: item.MenuItemID,
				Quantity:   item.Quantity,
			}
			if err := tx.Create(&orderItem).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Load order items
	if err := config.DB.Preload("OrderItems.MenuItem").First(&order, "id = ?", order.ID).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func GetOrdersBySession(restaurantID, sessionID string) ([]models.Order, error) {
	if strings.TrimSpace(restaurantID) == "" || strings.TrimSpace(sessionID) == "" {
		return nil, errors.New("restaurant ID and session ID are required")
	}

	var seatTable models.SeatTable
	if err := config.DB.First(&seatTable, "session_id = ?", sessionID).Error; err != nil {
		return nil, errors.New("session not found")
	}

	var orders []models.Order
	err := config.DB.Preload("OrderItems.MenuItem").
		Where("restaurant_id = ? AND table_id = ?", restaurantID, seatTable.ID).
		Order("created_at desc").
		Find(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func GetAllOrders(restaurantID string) ([]types.OrderResponse, error) {
	if strings.TrimSpace(restaurantID) == "" {
		return nil, errors.New("restaurant ID is required")
	}

	var orders []models.Order
	err := config.DB.Preload("OrderItems.MenuItem").Preload("Table").
		Where("restaurant_id = ?", restaurantID).
		Order("created_at asc").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	var responses []types.OrderResponse
	for _, o := range orders {
		items := []types.OrderItemResponse{}
		for _, item := range o.OrderItems {
			items = append(items, types.OrderItemResponse{
				ItemID:   item.MenuItem.ID,
				Name:     item.MenuItem.Name,
				Quantity: item.Quantity,
			})
		}
		responses = append(responses, types.OrderResponse{
			ID:        o.ID,
			TableName: o.Table.Name,
			CreatedAt: o.CreatedAt.Format(time.RFC3339),
			Status:    string(o.Status),
			Items:     items,
		})
	}
	return responses, nil
}

func UpdateOrderStatus(orderID, restaurantID, newStatus string) (*models.Order, error) {
	validStatuses := map[string]bool{
		"PENDING": true, "CONFIRMED": true, "COMPLETED": true, "CANCELLED": true,
	}

	if !validStatuses[newStatus] {
		return nil, errors.New("invalid order status")
	}

	var order models.Order
	if err := config.DB.First(&order, "id = ? AND restaurant_id = ?", restaurantID, orderID).Error; err != nil {
		return nil, errors.New("order not found")
	}

	order.Status = utils.ParseOrderStatus(newStatus)
	if err := config.DB.Save(&order).Error; err != nil {
		return nil, errors.New("failed to update order status")
	}

	return &order, nil
}
