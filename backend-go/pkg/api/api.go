package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KavyaGopal/Go-organic/backend-go/pkg/model"
	"github.com/gorilla/mux"
)

//init product variable for mock
var products []model.ProductMockUpdate
var fruits []model.FruitMock

//get all products
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(products)

}

//get single product
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var prodMockArray []model.ProdMaster
	model.DB.Find(&prodMockArray)
	json.NewEncoder(w).Encode(prodMockArray)

}

//get all fruits
func getFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(fruits)

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

	fruits = append(fruits, model.FruitMock{ID: 1, ImageSource: "../../../assets/items/apple.png", ItemName: "Apple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 12})
	fruits = append(fruits, model.FruitMock{ID: 2, ImageSource: "../../../assets/items/cherry.png", ItemName: "Cherry", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 15})
	fruits = append(fruits, model.FruitMock{ID: 3, ImageSource: "../../../assets/items/orange.png", ItemName: "Orange", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 5})
	fruits = append(fruits, model.FruitMock{ID: 4, ImageSource: "../../../assets/items/pineapple.jpeg", ItemName: "Pineapple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 9})
	fruits = append(fruits, model.FruitMock{ID: 5, ImageSource: "../../../assets/items/jackfruit.jpeg", ItemName: "Jackfruit", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 12})
	fruits = append(fruits, model.FruitMock{ID: 6, ImageSource: "../../../assets/items/watermelon.jpeg", ItemName: "Watermelon", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 20})


	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/getFruits", getFruits).Methods("GET")
	r.HandleFunc("/api/productsFromDB", getAllProductsFromDB).Methods("GET")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
