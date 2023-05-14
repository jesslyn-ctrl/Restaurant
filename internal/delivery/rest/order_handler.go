package rest

import (
	"encoding/json"
	"net/http"

	"github.com/jesslyn-ctrl/go-restaurant-app/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *handler) Order(c echo.Context) error {
	var request model.OrderMenuRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	orderData, err := h.restaurantUsecase.Order(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, orderData)
}

func (h *handler) GetOrderInfo(c echo.Context) error {
	orderId := c.Param("id")

	orderData, err := h.restaurantUsecase.GetOrderInfo(model.GetOrderInfoRequest{
		OrderId: orderId,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, orderData)
}
