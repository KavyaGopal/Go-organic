package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("ProductMock.db"))

	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
}

// Product Struct (Model)
type ProdMaster struct {
	ID          int64   `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not_null" json:"id"`
	Name        string  `gorm:"type:varchar(255);NOT NULL" json:"name"`
	Category    string  `"gorm:varchar(255)" json="category"`
	Price       float32 `gorm:"type:decimal(10,2)" json:"price"`
	Description string  `gorm:"type:text" json:"description"`
}

type FruitMock struct {
	ID            int64   `json:"id"`
	ImageSource   string   `json:"imgSrc"`
	ItemName      string `json:"itemName"`
	ItemDesc      string   `json:"itemDesc"`
	ItemWeight 	  int64   `json:"itemWt"`
	ItemQuantity  int64 `json:"itemQuantity"`
	ItemCost      int64 `json:"itemCost"`
}

type VegetablesMock struct {
	ID            int64   `json:"id"`
	ImageSource   string   `json:"imgSrc"`
	ItemName      string `json:"itemName"`
	ItemDesc      string   `json:"itemDesc"`
	ItemWeight 	  int64   `json:"itemWt"`
	ItemQuantity  int64 `json:"itemQuantity"`
	ItemCost      int64 `json:"itemCost"`
}

type GroceriesMock struct {
	ID            int64   `json:"id"`
	ImageSource   string   `json:"imgSrc"`
	ItemName      string `json:"itemName"`
	ItemDesc      string   `json:"itemDesc"`
	ItemWeight 	  int64   `json:"itemWt"`
	ItemQuantity  int64 `json:"itemQuantity"`
	ItemCost      int64 `json:"itemCost"`
}

type SnacksMock struct {
	ID            int64   `json:"id"`
	ImageSource   string   `json:"imgSrc"`
	ItemName      string `json:"itemName"`
	ItemDesc      string   `json:"itemDesc"`
	ItemWeight 	  int64   `json:"itemWt"`
	ItemQuantity  int64 `json:"itemQuantity"`
	ItemCost      int64 `json:"itemCost"`
}

type CosmeticsMock struct {
	ID            int64   `json:"id"`
	ImageSource   string   `json:"imgSrc"`
	ItemName      string `json:"itemName"`
	ItemDesc      string   `json:"itemDesc"`
	ItemWeight 	  int64   `json:"itemWt"`
	ItemQuantity  int64 `json:"itemQuantity"`
	ItemCost      int64 `json:"itemCost"`
}