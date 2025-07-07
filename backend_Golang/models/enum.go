package models

type GlobalRole string

const (
	GlobalRoleDev        GlobalRole = "DEV"
	GlobalRoleSuperAdmin GlobalRole = "SUPERADMIN"
	GlobalRoleCustomer   GlobalRole = "CUSTOMER"
)

type RestaurantRole string

const (
	RestaurantRoleOwner    RestaurantRole = "OWNER"
	RestaurantRoleManager  RestaurantRole = "MANAGER"
	RestaurantRoleEmployee RestaurantRole = "EMPLOYEE"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusConfirmed OrderStatus = "CONFIRMED"
	OrderStatusCompleted OrderStatus = "COMPLETED"
	OrderStatusPaid      OrderStatus = "PAID"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)
