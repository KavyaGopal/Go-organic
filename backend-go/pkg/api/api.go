package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KavyaGopal/Go-organic/backend-go/pkg/model"
	"github.com/gorilla/mux"
)

//init product variable for mock
var products []model.ProductMock

//get all products
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)

}

//get single product
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get the params
	// loop through products and find the id
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.ProductMock{})
}

//get the products from the database
func getAllProductsFromDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prodMockArray []model.ProdMaster
	model.DB.Find(&prodMockArray)
	json.NewEncoder(w).Encode(prodMockArray)

}

func main() {
	//init router
	r := mux.NewRouter()
	model.ConnectDatabase()
	//Mock data
	products = append(products, model.ProductMockUpdate{ID: "1", Name: "Groceries", Category: []string{"VeggieSalad", "VeggieBurrito", "KetoNuts"}})
	products = append(products, model.ProductMockUpdate{ID: "2", Name: "Fruits", Category: []string{"Mango", "Orange", "Banana"}})
	products = append(products, model.ProductMockUpdate{ID: "3", Name: "Vegetables", Category: []string{"Spinach", "Tomatoes", "Cabbage"}})
	products = append(products, model.ProductMockUpdate{ID: "4", Name: "Snacks", Category: []string{"Peanut Butter", "Waffles", "ChocoWalnut"}})
	products = append(products, model.ProductMockUpdate{ID: "5", Name: "Dairy", Category: []string{"Milk", "Curd", "Smoothie"}})

	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/productsFromDB", getAllProductsFromDB).Methods("GET")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
