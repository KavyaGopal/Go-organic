package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/KavyaGopal/Go-organic/backend-go/pkg/model"
	"encoding/json"
)
	
//init product variable for mock
var products []model.ProductMockUpdate

//get all products
func getProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)

}

//get single product
func getProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get the params
	// loop through products and find the id
	for _, item := range products {
		if item .ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
	json.NewEncoder(w).Encode(&model.ProductMock{})
}


func main() {
	//init router
	r := mux.NewRouter()

	//Mock data TODO: get the data from the sqlite
	// products = append(products, model.ProductMock{ID: "1", Name: "Groceries"})
	// products = append(products, model.ProductMock{ID: "2", Name: "Fruits"})
	// products = append(products, model.ProductMock{ID: "3", Name: "Snacks"})
	// products = append(products, model.ProductMock{ID: "4", Name: "Vegetables"})
	// products = append(products, model.ProductMock{ID: "5", Name: "Dairy"})

	products = append(products, model.ProductMockUpdate{ID: "1", Name: "Groceries", Category: []string{"VeggieSalad", "VeggieBurrito","KetoNuts"}})
	products = append(products, model.ProductMockUpdate{ID: "2", Name: "Fruits", Category: []string{"Mango", "Orange","Banana"}})
	products = append(products, model.ProductMockUpdate{ID: "3", Name: "Vegetables", Category: []string{"Spinach", "Tomatoes","Cabbage"}})
	products = append(products, model.ProductMockUpdate{ID: "4", Name: "Snacks", Category: []string{"Peanut Butter", "Waffles","ChocoWalnut"}})
	products = append(products, model.ProductMockUpdate{ID: "5", Name: "Dairy", Category: []string{"Milk", "Curd","Smoothie"}})
	
	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000",r))

}