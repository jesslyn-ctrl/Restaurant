package database

import (
	"github.com/jesslyn-ctrl/go-restaurant-app/internal/model"
	"github.com/jesslyn-ctrl/go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {

	// Migrate the schema
	db.AutoMigrate(&model.User{}, &model.MenuItem{}, &model.Order{}, &model.ProductOrder{})

	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     36000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Rica-rica",
			OrderCode: "ayam_rica_rica",
			Price:     32500,
			Type:      constant.MenuTypeFood,
		},
	}

	drinksMenu := []model.MenuItem{
		{
			Name:      "Lemon Tea",
			OrderCode: "lemon_tea",
			Price:     18000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Ocha",
			OrderCode: "ocha",
			Price:     10000,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinksMenu)
	}
}
