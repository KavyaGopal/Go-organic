package model

import(
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

// Product Struct (Model)
type ProdMaster struct {
	ID          int64   `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not_null" json:"id"`
	ImageSource string  `gorm:"type:varchar(255);NOT NULL" json:"imgSrc"`
	ItemName    string  `gorm:"type:varchar(255);NOT NULL" json:"itemName"`
	ItemCategory    string  `gorm:"type:varchar(255);NOT NULL" json:"itemCategory"`
	ItemDesc    string  `gorm:"type:text" json:"itemDesc"`
	ItemWeight  float32 `gorm:"type:decimal(10,2)" json:"itemWt"`
	ItemQuantity int64 	`gorm:"<-" json:"itemQuantity"`
	ItemCost     float32 `gorm:"type:decimal(10,2)" json:"itemCost"`
	
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