package rest

import "github.com/labstack/echo/v4"

func LoadRoutes(e *echo.Echo, handler *handler) {
	// User
	e.POST("/user/register", handler.RegisterUser)

	// Menu
	e.GET("/menu", handler.GetMenuList)

	// Order
	e.POST("/order", handler.Order)
	e.GET("/order/:id", handler.GetOrderInfo)
}
