package services

import (
	"fmt"
	"notification/entities"
	"notification/externalServices"
)

type OrderService struct {
	Notifier externalServices.Notifier
}

func (orderService *OrderService) CreateOrder(order *entities.Order) *entities.Order {
	fmt.Printf("Order created: %v\n", order)

	orderService.Notifier = externalServices.NewNotifier(order.NotificationType)

	orderService.Notifier.SendNotify(order.UserId, "Order created")

	return order
}

func NewOrderService() *OrderService {
	return &OrderService{}
}
