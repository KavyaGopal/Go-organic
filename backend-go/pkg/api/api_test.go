package main

import (
	"testing"

	// "github.com/gorilla/mux"

	// "github.com/stretchr/testify/assert"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/KavyaGopal/Go-organic/backend-go/pkg/model"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/getFruits", GetFruits).Methods("GET")
	router.HandleFunc("/getSnacks", GetSnacks).Methods("GET")
	router.HandleFunc("/getVegetables", GetVegetables).Methods("GET")
	router.HandleFunc("/getCosmetics", GetCosmetics).Methods("GET")
	router.HandleFunc("/getGroceries", GetGroceries).Methods("GET")

	//add apis to fetch data from db
	router.HandleFunc("/api/fetchAllProductsFromDB", GetAllProductsFromDB).Methods("GET")
	router.HandleFunc("/api/fetchProduct/{itemCategory}", GetFilteredCategory).Methods("GET")

	return router
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

//test case for fetching one product->fruits/groceries/snacks/cosmetics/vegetables
//one at a time
func TestGetFilteredCategory(t *testing.T) {
	// 	//write test case for fetch single product from db
	model.ConnectDatabase()
	categoriesArray := [5]string{"Fruits", "Snacks", "Vegetables", "Cosmetics", "Groceries"}
	for i := 0; i < len(categoriesArray); i++ {
		request, _ := http.NewRequest("GET", "/api/fetchProduct/"+categoriesArray[i], nil)
		rr := httptest.NewRecorder()
		// handler := http.HandlerFunc(getFruits)
		Router().ServeHTTP(rr, request)
		fmt.Println("response is", rr.Body.String())

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}

}

//test case for fetching all products->fruits/groceries/snacks/cosmetics/vegetables
//all at once
func TestGetAllProductsFromDB(t *testing.T) {
	// 	//write test cases for fetching all products from db
	model.ConnectDatabase()
	request, err := http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllProductsFromDB)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	fmt.Println("response is ", rr.Body.String())
	// assert.Equal(t, 200, response.Code, "OK response is expected")
}

// func TestHealthCheck(t *testing.T) {

// 	//write test cases for api health
// 	request, _ := http.NewRequest("GET", "/health-check", nil)
// 	response := httptest.NewRecorder()
// 	handler := http.HandlerFunc(HealthCheck)
// 	handler.ServeHTTP(response, request)

// 	if status := response.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	fmt.Println("response got is " + response.Body.String())

// 	expected := `API is up and running`
// 	if response.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			response.Body.String(), expected)
// 	}
// }

func TestGetFruits(t *testing.T) {
	// model.ConnectDatabase()
	fruits = append(fruits, model.FruitMock{ID: 1, ImageSource: "../../../assets/items/apple.png", ItemName: "Apple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 12})
	fruits = append(fruits, model.FruitMock{ID: 2, ImageSource: "../../../assets/items/cherry.png", ItemName: "Cherry", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 15})
	fruits = append(fruits, model.FruitMock{ID: 3, ImageSource: "../../../assets/items/orange.png", ItemName: "Orange", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 5})
	fruits = append(fruits, model.FruitMock{ID: 4, ImageSource: "../../../assets/items/pineapple.jpeg", ItemName: "Pineapple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9})
	fruits = append(fruits, model.FruitMock{ID: 5, ImageSource: "../../../assets/items/jackfruit.jpeg", ItemName: "Jackfruit", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 12})
	fruits = append(fruits, model.FruitMock{ID: 6, ImageSource: "../../../assets/items/watermelon.jpeg", ItemName: "Watermelon", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 20})

	snacks = append(snacks, model.SnacksMock{ID: 11, ImageSource: "../../../assets/items/cashew.jpeg", ItemName: "Roasted Cashew", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 100, ItemQuantity: 1, ItemCost: 15})
	snacks = append(snacks, model.SnacksMock{ID: 12, ImageSource: "../../../assets/items/peanuts.jpeg", ItemName: "Roasted Peanut", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 100, ItemQuantity: 1, ItemCost: 12})
	snacks = append(snacks, model.SnacksMock{ID: 13, ImageSource: "../../../assets/items/riceCake.jpeg", ItemName: "Carrot Sticks", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 50, ItemQuantity: 1, ItemCost: 9})
	snacks = append(snacks, model.SnacksMock{ID: 14, ImageSource: "../../../assets/items/mango.jpeg", ItemName: "Dried Mango", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 100, ItemQuantity: 1, ItemCost: 8})
	snacks = append(snacks, model.SnacksMock{ID: 15, ImageSource: "../../../assets/items/mix.jpeg", ItemName: "Trial Mix", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 100, ItemQuantity: 1, ItemCost: 5})
	snacks = append(snacks, model.SnacksMock{ID: 16, ImageSource: "../../../assets/items/riceCake.jpeg", ItemName: "Rice cake", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 100, ItemQuantity: 1, ItemCost: 12})

	vegetables = append(vegetables, model.VegetablesMock{ID: 21, ImageSource: "../../../assets/items/potato.jpeg", ItemName: "Potato", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 6})
	vegetables = append(vegetables, model.VegetablesMock{ID: 22, ImageSource: "../../../assets/items/tomato.jpeg", ItemName: "Tomato", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9})
	vegetables = append(vegetables, model.VegetablesMock{ID: 23, ImageSource: "../../../assets/items/cauliflower.jpeg", ItemName: "Cauliflower", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 5})
	vegetables = append(vegetables, model.VegetablesMock{ID: 24, ImageSource: "../../../assets/items/brinjal.jpeg", ItemName: "Brinjal", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 13})
	vegetables = append(vegetables, model.VegetablesMock{ID: 25, ImageSource: "../../../assets/items/onion.jpeg", ItemName: "Onion", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 15})
	vegetables = append(vegetables, model.VegetablesMock{ID: 26, ImageSource: "../../../assets/items/carrot.jpeg", ItemName: "Carrot", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 12})

	cosmetics = append(cosmetics, model.CosmeticsMock{ID: 31, ImageSource: "../../../assets/items/cleanser.jpeg", ItemName: "Cleanser", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 12})
	cosmetics = append(cosmetics, model.CosmeticsMock{ID: 32, ImageSource: "../../../assets/items/bodyBalm.jpeg", ItemName: "Body Balm", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 15})
	cosmetics = append(cosmetics, model.CosmeticsMock{ID: 33, ImageSource: "../../../assets/items/oil.jpeg", ItemName: "Hair Oil", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 5})
	cosmetics = append(cosmetics, model.CosmeticsMock{ID: 34, ImageSource: "../../../assets/items/faceMask.jpeg", ItemName: "Face Mask", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9})
	cosmetics = append(cosmetics, model.CosmeticsMock{ID: 35, ImageSource: "../../../assets/items/kajal.jpeg", ItemName: "Kajal", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 12})
	cosmetics = append(cosmetics, model.CosmeticsMock{ID: 36, ImageSource: "../../../assets/items/soap.jpeg", ItemName: "Soap", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 20})

	groceries = append(groceries, model.GroceriesMock{ID: 41, ImageSource: "../../../assets/items/rice.jpeg", ItemName: "Rice", ItemDesc: "Rice is the seed of the grass species Oryza sativa or less commonly Oryza glaberrima.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 12})
	groceries = append(groceries, model.GroceriesMock{ID: 42, ImageSource: "../../../assets/items/wheat.jpeg", ItemName: "Wheat", ItemDesc: "Wheat is a member of the grass family that produces a dry, one-seeded fruit commonly called a kernel.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 15})
	groceries = append(groceries, model.GroceriesMock{ID: 43, ImageSource: "../../../assets/items/sugar.jpeg", ItemName: "Sugar", ItemDesc: "Sugar is the generic name for sweet-tasting, soluble carbohydrates, many of which are used in food.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 5})
	groceries = append(groceries, model.GroceriesMock{ID: 44, ImageSource: "../../../assets/items/sugar.jpeg", ItemName: "Salt", ItemDesc: "SAlt is the generic name for salt-tasting, soluble carbohydrates, many of which are used in food. It is salty in nature.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9})
	groceries = append(groceries, model.GroceriesMock{ID: 45, ImageSource: "../../../assets/items/chilli.png", ItemName: "Chilli Powder", ItemDesc: "Chili powder is the dried, pulverized fruit of one or more varieties of chili pepper, sometimes with the addition of other spices.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 12})
	groceries = append(groceries, model.GroceriesMock{ID: 46, ImageSource: "../../../assets/items/chilli.png", ItemName: "Garam Masala", ItemDesc: "Garam masala is a blend of ground spices originating from South Asia.It is common in Indian, Pakistani, Nepalese and Bangladeshi.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 20})

	//write test cases for api health
	request, _ := http.NewRequest("GET", "/getFruits", nil)
	response := httptest.NewRecorder()
	// handler := http.HandlerFunc(GetFruits)
	Router().ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println("response got is ", response.Body)

	// expected := `API is up and running`
	// if response.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 	response.Body.String(), expected)
	// }
}
