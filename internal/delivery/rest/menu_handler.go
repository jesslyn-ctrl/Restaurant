package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetMenuList(c echo.Context) error {
	menuType := c.FormValue("type")

	menuData, err := h.restaurantUsecase.GetMenuList(menuType)
	if err != nil {
		fmt.Printf("Got error %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": menuData,
	})
}
