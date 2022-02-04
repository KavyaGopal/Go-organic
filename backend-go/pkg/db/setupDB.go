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
	db.AutoMigrate(&model.ProdMaster{}, &model.CategoryMaster{})

}
