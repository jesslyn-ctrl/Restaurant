package order

import "github.com/jesslyn-ctrl/go-restaurant-app/internal/model"

type Repository interface {
	CreateOrder(order model.Order) (model.Order, error)
	GetOrderInfo(orderId string) (model.Order, error)
}
