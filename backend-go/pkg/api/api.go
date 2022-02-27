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
var snacks []model.SnacksMock
var vegetables []model.VegetablesMock
var groceries []model.GroceriesMock
var cosmetics []model.CosmeticsMock

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

//get all snacks
func getSnacks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(snacks)

}

//get all vegetables
func getVegetables(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(vegetables)

}

//get all cosmetics
func getCosmetics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(cosmetics)

}

//get all groceries
func getGroceries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(groceries)

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

	snacks = append(snacks, model.SnacksMock{ID: 11, ImageSource: "../../../assets/items/cashew.jpeg", ItemName: "Apple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 12})
	snacks = append(snacks, model.SnacksMock{ID: 12, ImageSource: "../../../assets/items/cashew.jpeg", ItemName: "Cherry", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 15})
	snacks = append(snacks, model.SnacksMock{ID: 13, ImageSource: "../../../assets/items/cashew.jpeg", ItemName: "Orange", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 5})
	snacks = append(snacks, model.SnacksMock{ID: 14, ImageSource: "../../../assets/items/cashew.jpeg", ItemName: "Pineapple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 9})
	snacks = append(snacks, model.SnacksMock{ID: 15, ImageSource: "../../../assets/items/cashew.jpeg", ItemName: "Jackfruit", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 12})
	snacks = append(snacks, model.SnacksMock{ID: 16, ImageSource: "../../../assets/items/cashew.jpegn", ItemName: "Watermelon", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 20})

	vegetables = append(vegetables, model.VegetablesMock{ID: 21, ImageSource: "../../../assets/items/potato.jpeg", ItemName: "Apple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 12})
	vegetables = append(vegetables, model.VegetablesMock{ID: 22, ImageSource: "../../../assets/items/potato.jpeg", ItemName: "Cherry", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 15})
	vegetables = append(vegetables, model.VegetablesMock{ID: 23, ImageSource: "../../../assets/items/potato.jpeg", ItemName: "Orange", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 5})
	vegetables = append(vegetables, model.VegetablesMock{ID: 24, ImageSource: "../../../assets/items/potato.jpeg", ItemName: "Pineapple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 9})
	vegetables = append(vegetables, model.VegetablesMock{ID: 25, ImageSource: "../../../assets/items/potato.jpeg", ItemName: "Jackfruit", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 12})
	vegetables = append(vegetables, model.VegetablesMock{ID: 26, ImageSource: "../../../assets/items/potato.jpeg", ItemName: "Watermelon", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 20})

	cosmetics = append(cosmetics, model.CosmeticsMock{ID: 31, ImageSource: "../../../assets/items/apple.png", ItemName: "Apple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 12})
	cosmetics = append(cosmetics, model.CosmeticsMock{ID: 32, ImageSource: "../../../assets/items/cherry.png", ItemName: "Cherry", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 15})
	cosmetics = append(cosmetics, model.CosmeticsMock{ID: 33, ImageSource: "../../../assets/items/orange.png", ItemName: "Orange", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 5})
	cosmetics = append(cosmetics, model.CosmeticsMock{ID: 34, ImageSource: "../../../assets/items/pineapple.jpeg", ItemName: "Pineapple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 9})
	cosmetics = append(cosmetics, model.CosmeticsMock{ID: 35, ImageSource: "../../../assets/items/jackfruit.jpeg", ItemName: "Jackfruit", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 12})
	cosmetics = append(cosmetics, model.CosmeticsMock{ID: 36, ImageSource: "../../../assets/items/watermelon.jpeg", ItemName: "Watermelon", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 20})

	groceries = append(groceries, model.GroceriesMock{ID: 41, ImageSource: "../../../assets/items/rice.jpeg", ItemName: "Apple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 12})
	groceries = append(groceries, model.GroceriesMock{ID: 42, ImageSource: "../../../assets/items/rice.jpeg", ItemName: "Cherry", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 15})
	groceries = append(groceries, model.GroceriesMock{ID: 43, ImageSource: "../../../assets/items/rice.jpeg", ItemName: "Orange", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 5})
	groceries = append(groceries, model.GroceriesMock{ID: 44, ImageSource: "../../../assets/items/rice.jpeg", ItemName: "Pineapple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 9})
	groceries = append(groceries, model.GroceriesMock{ID: 45, ImageSource: "../../../assets/items/rice.jpeg", ItemName: "Jackfruit", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 12})
	groceries = append(groceries, model.GroceriesMock{ID: 46, ImageSource: "../../../assets/items/rice.jpeg", ItemName: "Watermelon", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 200, ItemCost: 20})

	

	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/getFruits", getFruits).Methods("GET")
	r.HandleFunc("/getSnacks", getSnacks).Methods("GET")
	r.HandleFunc("/getVegetables", getVegetables).Methods("GET")
	r.HandleFunc("/getCosmetics", getCosmetics).Methods("GET")
	r.HandleFunc("/getGroceries", getGroceries).Methods("GET")
	r.HandleFunc("/api/productsFromDB", getAllProductsFromDB).Methods("GET")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
