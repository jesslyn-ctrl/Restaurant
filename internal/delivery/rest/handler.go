package rest

import "github.com/jesslyn-ctrl/go-restaurant-app/internal/usecase/restaurant"

type handler struct {
	restaurantUsecase restaurant.Usecase
}

func NewHandler(restaurantUsecase restaurant.Usecase) *handler {
	return &handler{
		restaurantUsecase: restaurantUsecase,
	}
}
