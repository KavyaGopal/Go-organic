package main

import (
	"github.com/KavyaGopal/Go-organic/backend-go/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("ProductMock.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.ProdMaster{})
	// &model.CategoryMaster{})
	// var categories = []model.CategoryMaster{{CategoryName: "Fruits"}, {CategoryName: "Vegetables"}, {CategoryName: "Snacks"}, {CategoryName: "Dairy"}, {CategoryName: "Groceries"}}

	var products = []model.ProdMaster{{Name: "Mango", Category: "Fruits", Price: 20.32, Description: "Mango"},
		{Name: "Apple", Category: "Fruits", Price: 25.10, Description: "Apple"},
		{Name: "Spinach", Category: "Vegetables", Price: 10.32, Description: "Spinach"},
		{Name: "Cabbage", Category: "Vegetables", Price: 15.32, Description: "Cabbage"},
		{Name: "Waffles", Category: "Snacks", Price: 30.32, Description: "Waffles"},
		{Name: "Muffins", Category: "Snacks", Price: 11.32, Description: "Muffins"},
		{Name: "Milk", Category: "Dairy", Price: 14.32, Description: "Milk"},
		{Name: "Curd", Category: "Dairy", Price: 16.00, Description: "Curd"},
		{Name: "Nuts", Category: "Groceries", Price: 18.00, Description: "Nuts"},
		{Name: "Beans", Category: "Groceries", Price: 26.00, Description: "Beans"}}

	db.Create(&products)
}
