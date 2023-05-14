package restaurant

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jesslyn-ctrl/go-restaurant-app/internal/model"
	"github.com/jesslyn-ctrl/go-restaurant-app/internal/model/constant"
	"github.com/jesslyn-ctrl/go-restaurant-app/internal/repository/menu"
	"github.com/jesslyn-ctrl/go-restaurant-app/internal/repository/order"
	"github.com/jesslyn-ctrl/go-restaurant-app/internal/repository/user"
)

type restaurantUsecase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
	userRepo  user.Repository
}

func GetUsecase(menuRepo menu.Repository, orderRepo order.Repository, userRepo user.Repository) Usecase {
	return &restaurantUsecase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

func (r *restaurantUsecase) GetMenuList(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenuList(menuType)
}

func (r *restaurantUsecase) Order(request model.OrderMenuRequest) (model.Order, error) {
	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))

	var totalPrice int64 = 0
	for i, orderProduct := range request.OrderProducts {
		menuData, err := r.menuRepo.GetMenu(orderProduct.OrderCode)
		if err != nil {
			return model.Order{}, err
		}

		productOrderData[i] = model.ProductOrder{
			Id:        uuid.New().String(),
			OrderCode: menuData.OrderCode,
			Quantity:  orderProduct.Quantity,
			Price:     menuData.Price * int64(orderProduct.Quantity),
			Status:    constant.ProductOrderStatusPreparing,
		}

		totalPrice += productOrderData[i].Price
	}

	orderData := model.Order{
		Id:            uuid.New().String(),
		ReferenceId:   request.ReferenceId,
		Status:        constant.OrderStatusProcessed,
		ProductOrders: productOrderData,
		TotalPrice:    totalPrice,
	}

	createOrderData, err := r.orderRepo.CreateOrder(orderData)
	if err != nil {
		return model.Order{}, err
	}

	return createOrderData, nil
}

func (r *restaurantUsecase) GetOrderInfo(request model.GetOrderInfoRequest) (model.Order, error) {
	orderData, err := r.orderRepo.GetOrderInfo(request.OrderId)
	if err != nil {
		return orderData, err
	}

	return orderData, nil
}

func (r *restaurantUsecase) RegisterUser(request model.RegisterRequest) (model.User, error) {
	userRegistered, err := r.userRepo.CheckRegistered(request.Username)
	if err != nil {
		return model.User{}, err
	}
	if userRegistered {
		return model.User{}, errors.New("User already registered")
	}

	userHash, err := r.userRepo.GenerateUserHash(request.Password)
	if err != nil {
		return model.User{}, err
	}

	userData, err := r.userRepo.RegisterUser(model.User{
		Id:       uuid.New().String(),
		Username: request.Username,
		Hash:     userHash,
	})
	if err != nil {
		return model.User{}, err
	}

	return userData, nil
}
