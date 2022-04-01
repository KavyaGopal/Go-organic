package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/KavyaGopal/Go-organic/backend-go/pkg/model"
	"github.com/KavyaGopal/Go-organic/backend-go/pkg/db"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"bytes"
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

	//register api endpoint
	router.HandleFunc("/registerUser", RegisterUser).Methods("POST")
	//login api endpoint
	router.HandleFunc("/loginUser", LoginUser).Methods("POST")

	return router
}

//test case for fetching one product->fruits/groceries/snacks/cosmetics/vegetables
//one at a time
func TestGetFilteredCategory(t *testing.T) {
	// 	//write test case for fetch single product from db
	setupDB.ConnectDatabase()
	categoriesArray := [5]string{"Fruits", "Snacks", "Vegetables", "Cosmetics", "Groceries"}
	for i := 0; i < len(categoriesArray); i++ {
		request, _ := http.NewRequest("GET", "/api/fetchProduct/"+categoriesArray[i], nil)
		rr := httptest.NewRecorder()
		Router().ServeHTTP(rr, request)
		// fmt.Println("response is", rr.Body.String())

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expectedFruitsJson := `[{"id":1,"imgSrc":"../../../assets/items/apple.png","itemName":"Apple","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12},{"id":2,"imgSrc":"../../../assets/items/cherry.png","itemName":"Cherry","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":3,"imgSrc":"../../../assets/items/orange.png","itemName":"Orange","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":4,"imgSrc":"../../../assets/items/pineapple.jpeg","itemName":"Pineapple","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":5,"imgSrc":"../../../assets/items/jackfruit.jpeg","itemName":"Jackfruit","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":6,"imgSrc":"../../../assets/items/watermelon.jpeg","itemName":"Watermelon","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9}]`
		expectedVegetablesJson := `[{"id":7,"imgSrc":"../../../assets/items/potato.jpeg","itemName":"Potato","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":8,"imgSrc":"../../../assets/items/tomato.jpeg","itemName":"Tomato","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":9,"imgSrc":"../../../assets/items/cauliflower.jpeg","itemName":"Cauliflower","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":10,"imgSrc":"../../../assets/items/brinjal.jpeg","itemName":"Brinjal","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":11,"imgSrc":"../../../assets/items/onion.jpeg","itemName":"Onion","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8},{"id":12,"imgSrc":"../../../assets/items/carrot.jpeg","itemName":"Carrot","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8}]`
		expectedCosmeticsJson := `[{"id":19,"imgSrc":"../../../assets/items/cleanser.jpeg","itemName":"Cleanser","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12},{"id":20,"imgSrc":"../../../assets/items/bodyBalm.jpeg","itemName":"Body Balm","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12},{"id":21,"imgSrc":"../../../assets/items/oil.jpeg","itemName":"Hair Oil","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":13},{"id":22,"imgSrc":"../../../assets/items/faceMask.jpeg","itemName":"Face Mask","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":23,"imgSrc":"../../../assets/items/kajal.jpeg","itemName":"Kajal","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8},{"id":24,"imgSrc":"../../../assets/items/soap.jpeg","itemName":"Soap","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8}]`
		expectedSnacksJson := `[{"id":13,"imgSrc":"../../../assets/items/cashew.jpeg","itemName":"Roasted Cashew","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":14,"imgSrc":"../../../assets/items/peanuts.jpeg","itemName":"Roasted Peanut","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":11},{"id":15,"imgSrc":"../../../assets/items/riceCake.jpeg","itemName":"Carrot Sticks","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":13},{"id":16,"imgSrc":"../../../assets/items/mango.jpeg","itemName":"Dried Mango","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":17,"imgSrc":"../../../assets/items/mix.jpeg","itemName":"Trial Mix","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8},{"id":18,"imgSrc":"../../../assets/items/riceCake.jpeg","itemName":"Rice cake","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8}]`
		expectedGroceriesJson := `[{"id":25,"imgSrc":"../../../assets/items/rice.jpeg","itemName":"Rice","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":26,"imgSrc":"../../../assets/items/wheat.jpeg","itemName":"Wheat","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":27,"imgSrc":"../../../assets/items/sugar.jpeg","itemName":"Sugar","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":28,"imgSrc":"../../../assets/items/sugar.jpeg","itemName":"Salt","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":29,"imgSrc":"../../../assets/items/chilli.png","itemName":"Chilli Powder","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8},{"id":30,"imgSrc":"../../../assets/items/chilli.png","itemName":"Garam Masala","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8}]`
		
		var jsonMap map[string]string

    	jsonMap = make(map[string]string)

		jsonMap["Fruits"] = expectedFruitsJson
    	jsonMap["Vegetables"] = expectedVegetablesJson
    	jsonMap["Cosmetics"] = expectedCosmeticsJson
		jsonMap["Snacks"] = expectedSnacksJson
		jsonMap["Groceries"] = expectedGroceriesJson

		//check equality of json for each product
		require.JSONEq(t, jsonMap[categoriesArray[i]], rr.Body.String())

	}
	

}

//test case for fetching all products->fruits/groceries/snacks/cosmetics/vegetables
//all at once
func TestGetAllProductsFromDB(t *testing.T) {
	// 	//write test cases for fetching all products from db
	setupDB.ConnectDatabase()
	request, err := http.NewRequest("GET", "/api/fetchAllProductsFromDB", nil)
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

	//expected json for all products fetched from the database
	expected := `[{"id":1,"imgSrc":"../../../assets/items/apple.png","itemName":"Apple","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12},{"id":2,"imgSrc":"../../../assets/items/cherry.png","itemName":"Cherry","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":3,"imgSrc":"../../../assets/items/orange.png","itemName":"Orange","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":4,"imgSrc":"../../../assets/items/pineapple.jpeg","itemName":"Pineapple","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":5,"imgSrc":"../../../assets/items/jackfruit.jpeg","itemName":"Jackfruit","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":6,"imgSrc":"../../../assets/items/watermelon.jpeg","itemName":"Watermelon","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":7,"imgSrc":"../../../assets/items/potato.jpeg","itemName":"Potato","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":8,"imgSrc":"../../../assets/items/tomato.jpeg","itemName":"Tomato","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":9,"imgSrc":"../../../assets/items/cauliflower.jpeg","itemName":"Cauliflower","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":10,"imgSrc":"../../../assets/items/brinjal.jpeg","itemName":"Brinjal","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":11,"imgSrc":"../../../assets/items/onion.jpeg","itemName":"Onion","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8},{"id":12,"imgSrc":"../../../assets/items/carrot.jpeg","itemName":"Carrot","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8},{"id":13,"imgSrc":"../../../assets/items/cashew.jpeg","itemName":"Roasted Cashew","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":14,"imgSrc":"../../../assets/items/peanuts.jpeg","itemName":"Roasted Peanut","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":11},{"id":15,"imgSrc":"../../../assets/items/riceCake.jpeg","itemName":"Carrot Sticks","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":13},{"id":16,"imgSrc":"../../../assets/items/mango.jpeg","itemName":"Dried Mango","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":17,"imgSrc":"../../../assets/items/mix.jpeg","itemName":"Trial Mix","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8},{"id":18,"imgSrc":"../../../assets/items/riceCake.jpeg","itemName":"Rice cake","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8},{"id":19,"imgSrc":"../../../assets/items/cleanser.jpeg","itemName":"Cleanser","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12},{"id":20,"imgSrc":"../../../assets/items/bodyBalm.jpeg","itemName":"Body Balm","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12},{"id":21,"imgSrc":"../../../assets/items/oil.jpeg","itemName":"Hair Oil","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":13},{"id":22,"imgSrc":"../../../assets/items/faceMask.jpeg","itemName":"Face Mask","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":23,"imgSrc":"../../../assets/items/kajal.jpeg","itemName":"Kajal","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8},{"id":24,"imgSrc":"../../../assets/items/soap.jpeg","itemName":"Soap","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8},{"id":25,"imgSrc":"../../../assets/items/rice.jpeg","itemName":"Rice","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":26,"imgSrc":"../../../assets/items/wheat.jpeg","itemName":"Wheat","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":27,"imgSrc":"../../../assets/items/sugar.jpeg","itemName":"Sugar","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10},{"id":28,"imgSrc":"../../../assets/items/sugar.jpeg","itemName":"Salt","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":29,"imgSrc":"../../../assets/items/chilli.png","itemName":"Chilli Powder","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8},{"id":30,"imgSrc":"../../../assets/items/chilli.png","itemName":"Garam Masala","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8}]`
	
	//check equality of json for all products
	require.JSONEq(t, expected, rr.Body.String())

	

}

//test register api
func TestRegisterAPI(t *testing.T) {
	//write test cases for the register api
	setupDB.ConnectDatabase()
	user := &model.User{
		Name:         "test",
		Email:        "test",
		Address:      "test",
		Password: 	  "test",
		Age:           10,
		Phone:        "3528889222",
	}
	jsonPayload, _ := json.Marshal(user)
	request, _ := http.NewRequest("POST", "/registerUser", bytes.NewBuffer(jsonPayload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RegisterUser)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	assert.Equal(t, 200, rr.Code, "OK response is expected")
	assert.Equal(t, "{\"status\":200,\"message\":\"New User Successfully Added: test\"}\n", rr.Body.String(), "Register Success")

}

//test for sample mock json->Fruit
func TestGetFruits(t *testing.T) {

	fruits = append(fruits, model.FruitMock{ID: 1, ImageSource: "../../../assets/items/apple.png", ItemName: "Apple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 12})
	fruits = append(fruits, model.FruitMock{ID: 2, ImageSource: "../../../assets/items/cherry.png", ItemName: "Cherry", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 15})
	fruits = append(fruits, model.FruitMock{ID: 3, ImageSource: "../../../assets/items/orange.png", ItemName: "Orange", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 5})
	fruits = append(fruits, model.FruitMock{ID: 4, ImageSource: "../../../assets/items/pineapple.jpeg", ItemName: "Pineapple", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 9})
	fruits = append(fruits, model.FruitMock{ID: 5, ImageSource: "../../../assets/items/jackfruit.jpeg", ItemName: "Jackfruit", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 12})
	fruits = append(fruits, model.FruitMock{ID: 6, ImageSource: "../../../assets/items/watermelon.jpeg", ItemName: "Watermelon", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemCost: 20})
	
	//write test cases for api health
	request, _ := http.NewRequest("GET", "/getFruits", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//expected json for the fruit
	expected := `[{"id":1,"imgSrc":"../../../assets/items/apple.png","itemName":"Apple","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12},{"id":2,"imgSrc":"../../../assets/items/cherry.png","itemName":"Cherry","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":15},{"id":3,"imgSrc":"../../../assets/items/orange.png","itemName":"Orange","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":5},{"id":4,"imgSrc":"../../../assets/items/pineapple.jpeg","itemName":"Pineapple","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9},{"id":5,"imgSrc":"../../../assets/items/jackfruit.jpeg","itemName":"Jackfruit","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12},{"id":6,"imgSrc":"../../../assets/items/watermelon.jpeg","itemName":"Watermelon","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":20}]`
	
	//check equality of json for the fruit product
	require.JSONEq(t, expected, response.Body.String())
}




