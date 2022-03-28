package model

// Product Struct (Model)
type ProdMaster struct {
	ID          int64   `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not_null" json:"id"`
	ImageSource string  `gorm:"type:varchar(255);NOT NULL" json:"imgSrc"`
	ItemName    string  `gorm:"type:varchar(255);NOT NULL" json:"itemName"`
	ItemCategory string  `gorm:"type:varchar(255);NOT NULL" json:"itemCategory"`
	ItemDesc    string  `gorm:"type:text" json:"itemDesc"`
	ItemWeight  float32 `gorm:"type:decimal(10,2)" json:"itemWt"`
	ItemQuantity int64 	`gorm:"<-" json:"itemQuantity"`
	ItemCost     float32 `gorm:"type:decimal(10,2)" json:"itemCost"`
	
}

//user profile
type User struct {
	ID           int64  `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not_null" json:"id"`
	Name         string `gorm:"type:varchar(255);NOT NULL" json:"name"`
	Email        string `gorm:"unique; Not null " json:"email"`
	Address      string `gorm:"type:varchar(255);NOT NULL" json:"address"`
	Password     string `gorm:"Not null " json:"password"`
	Age          int64 	`gorm:"<-" json:"age"`
	Phone      	 int64 `gorm:"unique; Not null " json:"phone"`
}

type Login struct {
	Email    string
	Password string
}


type Address struct {
	Street  string `gorm:"type:varchar(255);NOT NULL" json:"street"`
	Apartment  string `gorm:"type:varchar(255);NOT NULL" json:"apartment"`
	City     string `gorm:"type:varchar(255);NOT NULL" json:"city"`
	State    string `gorm:"type:varchar(255);NOT NULL" json:"state"`
	Zip      int64 	`gorm:"<-" json:"zipcode"`
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