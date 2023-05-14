package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jesslyn-ctrl/go-restaurant-app/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *handler) RegisterUser(c echo.Context) error {
	var request model.RegisterRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("Got error %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userData, err := h.restaurantUsecase.RegisterUser(request)
	if err != nil {
		fmt.Printf("Got error %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, userData)
}
