package main

import (
	"github.com/jesslyn-ctrl/go-restaurant-app/internal/database"
	"github.com/jesslyn-ctrl/go-restaurant-app/internal/delivery/rest"
	mRepo "github.com/jesslyn-ctrl/go-restaurant-app/internal/repository/menu"
	oRepo "github.com/jesslyn-ctrl/go-restaurant-app/internal/repository/order"
	uRepo "github.com/jesslyn-ctrl/go-restaurant-app/internal/repository/user"
	rUsecase "github.com/jesslyn-ctrl/go-restaurant-app/internal/usecase/restaurant"
	"github.com/labstack/echo/v4"
)

const dbAddress = "host=localhost port=5432 user=postgres password=postgres dbname=restaurant sslmode=disable"

func main() {
	e := echo.New()

	db := database.GetDB(dbAddress)
	secret := "AES256Key-32Characters1234567890"

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32)
	if err != nil {
		panic(err)
	}

	restaurantUsecase := rUsecase.GetUsecase(menuRepo, orderRepo, userRepo)

	handler := rest.NewHandler(restaurantUsecase)

	rest.LoadMiddlewares(e)
	rest.LoadRoutes(e, handler)

	e.Logger.Fatal(e.Start(":14045"))
}
