package utils

import (
	"strings"

	"github.com/google/uuid"

	"github.com/nineard99/restaurant-Golang/models"
)

func NullableString(s string) *string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	return &s
}

func ParseGlobalRole(roleStr string) models.GlobalRole {
	switch strings.ToUpper(roleStr) {
	case "DEV":
		return models.GlobalRoleDev
	case "SUPERADMIN":
		return models.GlobalRoleSuperAdmin
	case "CUSTOMER":
		return models.GlobalRoleCustomer
	default:
		return models.GlobalRoleCustomer
	}
}
func ParseOrderStatus(status string) models.OrderStatus {
	switch strings.ToUpper(status) {
	case "PENDING":
		return models.OrderStatusPending
	case "CONFIRMED":
		return models.OrderStatusConfirmed
	case "COMPLETED":
		return models.OrderStatusCompleted
	case "PAID":
		return models.OrderStatusPaid
	case "CANCELLED":
		return models.OrderStatusCancelled
	default:
		return models.OrderStatusPending
	}
}

func GenerateUUID() string {
	return uuid.New().String()
}
