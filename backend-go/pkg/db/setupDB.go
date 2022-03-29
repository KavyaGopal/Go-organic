package setupDB

import (
	"github.com/KavyaGopal/Go-organic/backend-go/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("ProductData.db"))

	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
}

func ConnectEndPointDatabase() {
	database, err := gorm.Open(sqlite.Open("pkg/api/ProductData.db"))

	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
}

func CreateDatabase(){

	db, err := gorm.Open(sqlite.Open("pkg/api/ProductData.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.ProdMaster{})
	db.AutoMigrate(&model.User{})

	var products = []model.ProdMaster{
		//fruits
		{ImageSource: "../../../assets/items/apple.png", ItemName: "Apple", ItemCategory: "Fruits", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 12},
		{ImageSource: "../../../assets/items/cherry.png", ItemName: "Cherry", ItemCategory: "Fruits", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 10},
		{ImageSource: "../../../assets/items/orange.png", ItemName: "Orange", ItemCategory: "Fruits", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9},
		{ImageSource: "../../../assets/items/pineapple.jpeg", ItemName: "Pineapple", ItemCategory: "Fruits", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9},
		{ImageSource: "../../../assets/items/jackfruit.jpeg", ItemName: "Jackfruit", ItemCategory: "Fruits", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9},
		{ImageSource: "../../../assets/items/watermelon.jpeg", ItemName: "Watermelon", ItemCategory: "Fruits", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9},
		//vegetables
		{ImageSource: "../../../assets/items/potato.jpeg", ItemName: "Potato", ItemCategory: "Vegetables", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 10},
		{ImageSource: "../../../assets/items/tomato.jpeg", ItemName: "Tomato", ItemCategory: "Vegetables", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9},
		{ImageSource: "../../../assets/items/cauliflower.jpeg", ItemName: "Cauliflower", ItemCategory: "Vegetables", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 10},
		{ImageSource: "../../../assets/items/brinjal.jpeg", ItemName: "Brinjal", ItemCategory: "Vegetables", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9},
		{ImageSource: "../../../assets/items/onion.jpeg", ItemName: "Onion", ItemCategory: "Vegetables", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 8},
		{ImageSource: "../../../assets/items/carrot.jpeg", ItemName: "Carrot", ItemCategory: "Vegetables", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 8},
		//snacks
		{ImageSource: "../../../assets/items/cashew.jpeg", ItemName: "Roasted Cashew", ItemCategory: "Snacks", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 10},
		{ImageSource: "../../../assets/items/peanuts.jpeg", ItemName: "Roasted Peanut", ItemCategory: "Snacks", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 11},
		{ImageSource: "../../../assets/items/riceCake.jpeg", ItemName: "Carrot Sticks", ItemCategory: "Snacks", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 13},
		{ImageSource: "../../../assets/items/mango.jpeg", ItemName: "Dried Mango", ItemCategory: "Snacks", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9},
		{ImageSource: "../../../assets/items/mix.jpeg", ItemName: "Trial Mix", ItemCategory: "Snacks", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 8},
		{ImageSource: "../../../assets/items/riceCake.jpeg", ItemName: "Rice cake", ItemCategory: "Snacks", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 8},
		//cosmetics
		{ImageSource: "../../../assets/items/cleanser.jpeg", ItemName: "Cleanser", ItemCategory: "Cosmetics", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 12},
		{ImageSource: "../../../assets/items/bodyBalm.jpeg", ItemName: "Body Balm", ItemCategory: "Cosmetics", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 12},
		{ImageSource: "../../../assets/items/oil.jpeg", ItemName: "Hair Oil", ItemCategory: "Cosmetics", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 13},
		{ImageSource: "../../../assets/items/faceMask.jpeg", ItemName: "Face Mask", ItemCategory: "Cosmetics", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 10},
		{ImageSource: "../../../assets/items/kajal.jpeg", ItemName: "Kajal", ItemCategory: "Cosmetics", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 8},
		{ImageSource: "../../../assets/items/soap.jpeg", ItemName: "Soap", ItemCategory: "Cosmetics", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 8},
		//groceries
		{ImageSource: "../../../assets/items/rice.jpeg", ItemName: "Rice", ItemCategory: "Groceries", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 10},
		{ImageSource: "../../../assets/items/wheat.jpeg", ItemName: "Wheat", ItemCategory: "Groceries", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9},
		{ImageSource: "../../../assets/items/sugar.jpeg", ItemName: "Sugar", ItemCategory: "Groceries", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 10},
		{ImageSource: "../../../assets/items/sugar.jpeg", ItemName: "Salt", ItemCategory: "Groceries", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9},
		{ImageSource: "../../../assets/items/chilli.png", ItemName: "Chilli Powder", ItemCategory: "Groceries", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 8},
		{ImageSource: "../../../assets/items/chilli.png", ItemName: "Garam Masala", ItemCategory: "Groceries", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 8}}

	db.Create(&products)

}
