package model

import "gorm.io/gorm"

// Product Struct (Model)
type ProdMaster struct {
	gorm.Model
	ID          int    `gorm:"primary_key" json:"id"`
	Name        string `gorm:"type:varchar(255);NOT NULL" json:"name"`
	Category    CategoryMaster
	Price       float32 `gorm:"type:decimal(10,2)" json:"price"`
	Description string  `gorm:"type:text" json:"description"`
}
type CategoryMaster struct {
	gorm.Model
	CatID        int    `gorm:"primary_key" json:"catid"`
	CategoryName string `gorm:"type:varchar(255)" json:"categoryname"`
}
