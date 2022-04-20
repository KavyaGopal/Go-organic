package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/KavyaGopal/Go-organic/backend-go/pkg/db"
	"github.com/KavyaGopal/Go-organic/backend-go/pkg/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	router.HandleFunc("/api/fetchItemQuantity", FetchItemQuantity).Methods("POST")
	router.HandleFunc("/api/fetchUserTestimonials", GetUserTestimonials).Methods("GET")

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

		expectedFruitsJson := `[{"id":1,"imgSrc":"../../../assets/items/apple.png","itemName":"Apple","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12,"itemInventory":200},{"id":2,"imgSrc":"../../../assets/items/cherry.png","itemName":"Cherry","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":3,"imgSrc":"../../../assets/items/orange.png","itemName":"Orange","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":4,"imgSrc":"../../../assets/items/pineapple.jpeg","itemName":"Pineapple","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":5,"imgSrc":"../../../assets/items/jackfruit.jpeg","itemName":"Jackfruit","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":6,"imgSrc":"../../../assets/items/watermelon.jpeg","itemName":"Watermelon","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200}]`
		expectedVegetablesJson := `[{"id":7,"imgSrc":"../../../assets/items/potato.jpeg","itemName":"Potato","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":8,"imgSrc":"../../../assets/items/tomato.jpeg","itemName":"Tomato","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":9,"imgSrc":"../../../assets/items/cauliflower.jpeg","itemName":"Cauliflower","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":10,"imgSrc":"../../../assets/items/brinjal.jpeg","itemName":"Brinjal","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":11,"imgSrc":"../../../assets/items/onion.jpeg","itemName":"Onion","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200},{"id":12,"imgSrc":"../../../assets/items/carrot.jpeg","itemName":"Carrot","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200}]`
		expectedCosmeticsJson := `[{"id":19,"imgSrc":"../../../assets/items/cleanser.jpeg","itemName":"Cleanser","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12,"itemInventory":200},{"id":20,"imgSrc":"../../../assets/items/bodyBalm.jpeg","itemName":"Body Balm","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12,"itemInventory":200},{"id":21,"imgSrc":"../../../assets/items/oil.jpeg","itemName":"Hair Oil","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":13,"itemInventory":200},{"id":22,"imgSrc":"../../../assets/items/faceMask.jpeg","itemName":"Face Mask","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":23,"imgSrc":"../../../assets/items/kajal.jpeg","itemName":"Kajal","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200},{"id":24,"imgSrc":"../../../assets/items/soap.jpeg","itemName":"Soap","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200}]`
		expectedSnacksJson := `[{"id":13,"imgSrc":"../../../assets/items/cashew.jpeg","itemName":"Roasted Cashew","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":14,"imgSrc":"../../../assets/items/peanuts.jpeg","itemName":"Roasted Peanut","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":11,"itemInventory":200},{"id":15,"imgSrc":"../../../assets/items/riceCake.jpeg","itemName":"Carrot Sticks","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":13,"itemInventory":200},{"id":16,"imgSrc":"../../../assets/items/mango.jpeg","itemName":"Dried Mango","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":17,"imgSrc":"../../../assets/items/mix.jpeg","itemName":"Trial Mix","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200},{"id":18,"imgSrc":"../../../assets/items/riceCake.jpeg","itemName":"Rice cake","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200}]`
		expectedGroceriesJson := `[{"id":25,"imgSrc":"../../../assets/items/rice.jpeg","itemName":"Rice","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":26,"imgSrc":"../../../assets/items/wheat.jpeg","itemName":"Wheat","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":27,"imgSrc":"../../../assets/items/sugar.jpeg","itemName":"Sugar","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":28,"imgSrc":"../../../assets/items/sugar.jpeg","itemName":"Salt","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":29,"imgSrc":"../../../assets/items/chilli.png","itemName":"Chilli Powder","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200},{"id":30,"imgSrc":"../../../assets/items/chilli.png","itemName":"Garam Masala","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200}]`

		var jsonMap map[string]string

		jsonMap = make(map[string]string)

		jsonMap["Fruits"] = expectedFruitsJson
		jsonMap["Vegetables"] = expectedVegetablesJson
		jsonMap["Cosmetics"] = expectedCosmeticsJson
		jsonMap["Snacks"] = expectedSnacksJson
		jsonMap["Groceries"] = expectedGroceriesJson

		//check equality of json for each product
		require.JSONEq(t, jsonMap[categoriesArray[i]], rr.Body.String())

		log.Println("######## Test Case Passed For Filtered Product API ###########")

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
	expected := `[{"id":1,"imgSrc":"../../../assets/items/apple.png","itemName":"Apple","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12,"itemInventory":200},{"id":2,"imgSrc":"../../../assets/items/cherry.png","itemName":"Cherry","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":3,"imgSrc":"../../../assets/items/orange.png","itemName":"Orange","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":4,"imgSrc":"../../../assets/items/pineapple.jpeg","itemName":"Pineapple","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":5,"imgSrc":"../../../assets/items/jackfruit.jpeg","itemName":"Jackfruit","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":6,"imgSrc":"../../../assets/items/watermelon.jpeg","itemName":"Watermelon","itemCategory":"Fruits","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":7,"imgSrc":"../../../assets/items/potato.jpeg","itemName":"Potato","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":8,"imgSrc":"../../../assets/items/tomato.jpeg","itemName":"Tomato","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":9,"imgSrc":"../../../assets/items/cauliflower.jpeg","itemName":"Cauliflower","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":10,"imgSrc":"../../../assets/items/brinjal.jpeg","itemName":"Brinjal","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":11,"imgSrc":"../../../assets/items/onion.jpeg","itemName":"Onion","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200},{"id":12,"imgSrc":"../../../assets/items/carrot.jpeg","itemName":"Carrot","itemCategory":"Vegetables","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200},{"id":13,"imgSrc":"../../../assets/items/cashew.jpeg","itemName":"Roasted Cashew","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":14,"imgSrc":"../../../assets/items/peanuts.jpeg","itemName":"Roasted Peanut","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":11,"itemInventory":200},{"id":15,"imgSrc":"../../../assets/items/riceCake.jpeg","itemName":"Carrot Sticks","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":13,"itemInventory":200},{"id":16,"imgSrc":"../../../assets/items/mango.jpeg","itemName":"Dried Mango","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":17,"imgSrc":"../../../assets/items/mix.jpeg","itemName":"Trial Mix","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200},{"id":18,"imgSrc":"../../../assets/items/riceCake.jpeg","itemName":"Rice cake","itemCategory":"Snacks","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200},{"id":19,"imgSrc":"../../../assets/items/cleanser.jpeg","itemName":"Cleanser","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12,"itemInventory":200},{"id":20,"imgSrc":"../../../assets/items/bodyBalm.jpeg","itemName":"Body Balm","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":12,"itemInventory":200},{"id":21,"imgSrc":"../../../assets/items/oil.jpeg","itemName":"Hair Oil","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":13,"itemInventory":200},{"id":22,"imgSrc":"../../../assets/items/faceMask.jpeg","itemName":"Face Mask","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":23,"imgSrc":"../../../assets/items/kajal.jpeg","itemName":"Kajal","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200},{"id":24,"imgSrc":"../../../assets/items/soap.jpeg","itemName":"Soap","itemCategory":"Cosmetics","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200},{"id":25,"imgSrc":"../../../assets/items/rice.jpeg","itemName":"Rice","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":26,"imgSrc":"../../../assets/items/wheat.jpeg","itemName":"Wheat","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":27,"imgSrc":"../../../assets/items/sugar.jpeg","itemName":"Sugar","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":10,"itemInventory":200},{"id":28,"imgSrc":"../../../assets/items/sugar.jpeg","itemName":"Salt","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":9,"itemInventory":200},{"id":29,"imgSrc":"../../../assets/items/chilli.png","itemName":"Chilli Powder","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200},{"id":30,"imgSrc":"../../../assets/items/chilli.png","itemName":"Garam Masala","itemCategory":"Groceries","itemDesc":"This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.","itemWt":500,"itemQuantity":1,"itemCost":8,"itemInventory":200}]`

	//check equality of json for all products
	require.JSONEq(t, expected, rr.Body.String())

	log.Println("########## Test Case Passed For Getting All Products from DB ############")

}

func TestRegisterAPI(t *testing.T) {

	setupDB.ConnectDatabase()
	var userDel model.User
	setupDB.DB.Where("email = ?", "test1").Delete(&userDel)
	user := &model.User{
		Name:     "test",
		Email:    "test1",
		Address:  "test",
		Password: "test",
		Phone:    "3528881112",
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

	//success case
	assert.Equal(t, 200, rr.Code, "OK response is expected")
	assert.Equal(t, "{\"status\":200,\"message\":\"New User Successfully Added: test\"}\n", rr.Body.String(), "Register Success")
	log.Println("###### Test Case Passed For Register Success #########")

	//fail case
	request, _ = http.NewRequest("POST", "/registerUser", bytes.NewBuffer(jsonPayload))
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(RegisterUser)
	handler.ServeHTTP(rr, request)
	assert.Equal(t, "{\"status\":500,\"message\":\"User already exists\"}\n", rr.Body.String(), "Register Failure")
	log.Println("###### Test Case Passed For Register Failure #########")

}

func TestLoginAPI(t *testing.T) {

	setupDB.ConnectDatabase()
	loginUser := &model.Login{

		Email: "test1",

		Password: "test",
	}
	jsonPayload, _ := json.Marshal(loginUser)
	request, _ := http.NewRequest("POST", "/loginUser", bytes.NewBuffer(jsonPayload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginUser)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//success case
	assert.Equal(t, 200, rr.Code, "OK response is expected")
	assert.Equal(t, "{\"name\":\"test\",\"email\":\"test1\",\"status\":200,\"message\":\"User successfully logged in\"}\n", rr.Body.String(), "Login Success")
	log.Println("###### Test Case Passed For Login Success #########")
	loginUser = &model.Login{

		Email: "test1",

		Password: "test2",
	}
	jsonPayload, _ = json.Marshal(loginUser)
	request, _ = http.NewRequest("POST", "/loginUser", bytes.NewBuffer(jsonPayload))
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(LoginUser)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//fail case
	request, _ = http.NewRequest("POST", "/loginUser", bytes.NewBuffer(jsonPayload))
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(LoginUser)
	handler.ServeHTTP(rr, request)
	assert.Equal(t, "{\"status\":500,\"message\":\"Unauthorized Access\"}\n", rr.Body.String(), "Login Failure")
	log.Println("###### Test Case Passed For Login Failure #########")

}

func TestResetPassword(t *testing.T) {

	setupDB.ConnectDatabase()
	loginUser := &model.Login{

		Email: "test1",

		Password: "test1",
	}
	jsonPayload, _ := json.Marshal(loginUser)
	request, _ := http.NewRequest("POST", "/resetPassword", bytes.NewBuffer(jsonPayload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ResetPassword)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//success case
	assert.Equal(t, 200, rr.Code, "OK response is expected")
	assert.Equal(t, "{\"status\":200,\"message\":\"Password successfully updated.\"}\n", rr.Body.String(), "Password Reset Success")
	log.Println("###### Test Case Passed For Password Reset Success #########")

	//case for re-enter login with old password

	loginUser = &model.Login{

		Email: "test1",

		Password: "test",
	}
	jsonPayload, _ = json.Marshal(loginUser)
	request, _ = http.NewRequest("POST", "/loginUser", bytes.NewBuffer(jsonPayload))
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(LoginUser)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	assert.Equal(t, "{\"status\":500,\"message\":\"Unauthorized Access\"}\n", rr.Body.String(), "Login Attempt failed with old password")
	log.Println("###### Test Case Passed For Re-enter Login with Old Password #########")

}

func TestCartCheckout(t *testing.T) {

	setupDB.ConnectDatabase()
	ItemList := `[{
		"itemID": 1,
		"itemQuantity": 2
	}, {
		"itemID": 2,
		"itemQuantity": 5
	}]`

	jsonPayload, _ := json.Marshal(ItemList)
	request, _ := http.NewRequest("POST", "/api/fetchItemQuantity", bytes.NewBuffer(jsonPayload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FetchItemQuantity)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	assert.Equal(t, 200, rr.Code, "OK response is expected")
	assert.Equal(t, "{\"status\":200,\"message\":\"Checkout of All Items Done Successfully\"}\n", rr.Body.String(), "Cart Checkout Success")
	log.Println("###### Test Case Passed For Cart Checkout #########")

}

func TestUserTestimonials(t *testing.T){

	request, err := http.NewRequest("GET", "/api/fetchUserTestimonials", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserTestimonials)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//expected json for all user testimonials fetched from the database
	expected := `[{"id":1,"userName":"Zack Wright","imgSrc":"../../../assets/testimonial-images/img-m1.png","userDesc":"I am in my mid thirties and my weight gain hit me hard. I have tried every diet there is but instead of dieting now...I ordered go-organic products and who knew this was too good for you! Now I eat and truly enjoy my food."},{"id":2,"userName":"Alexa Brighton","imgSrc":"../../../assets/testimonial-images/img-fm1.jpeg","userDesc":"These products are easy and super healthy. The choices are great. I would highly recommend to use go-organic products to my friends and first timers."},{"id":3,"userName":"Joe Tribbiani","imgSrc":"../../../assets/testimonial-images/img-m2.jpeg","userDesc":"I've been dabbling with the Vegan diet on and off since the beginning of the year. These products have really given me inspiration in terms of meal plans and different recipes to try."},{"id":4,"userName":"Myra Philips","imgSrc":"../../../assets/testimonial-images/img-fm2.jpeg","userDesc":"I bought this to refresh my diet and nutrition and wanted some new healthy diets. Every single thing I have tried has been a winner and products are relatively cheaper."},{"id":5,"userName":"Jay Shah","imgSrc":"../../../assets/testimonial-images/img-m3.jpeg","userDesc":"Mostly tasty products. I am not a vegan and enjoyed most of the go-organic products. I highly recommended these products if you want to become plant eater and not feel hungry."},{"id":6,"userName":"Phoebe Buffet","imgSrc":"../../../assets/testimonial-images/img-fm3.jpeg","userDesc":"Being healthy is in my genes and I highly appreciate that go-organic is promoting healty diets in the form of organic based products."}]`

	//check equality of json for all products
	require.JSONEq(t, expected, rr.Body.String())

	log.Println("########## Test Case Passed For Getting User Testimonials from DB ############")
	
}
