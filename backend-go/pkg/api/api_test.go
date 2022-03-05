package main

import(
	"testing"
	"github.com/gorilla/mux"
	// "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"fmt"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getFruits", GetFruits).Methods("GET")
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	return router
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
	r := mux.NewRouter()
    r.ServeHTTP(rr, req)

    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

//test case for fetching one product->fruits/groceries/snacks/cosmetics/vegetables
//one at a time
func TestGetFilteredCategory(t *testing.T){
// 	//write test case for fetch single product from db

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(getFruits)
// 	handler.ServeHTTP(rr, req)
// 	fmt.Println(rr.Code)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

}

//test case for fetching all products->fruits/groceries/snacks/cosmetics/vegetables
//all at once
func TestGetAllProductsFromDB(t *testing.T){
// 	//write test cases for fetching all products from db
	request, err := http.NewRequest("GET", "/getFruits", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetFruits)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	fmt.Println("response is ", rr.Body.String())
	// assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestHealthCheck(t *testing.T){

	//write test cases for api health
	request, _ := http.NewRequest("GET", "/health-check", nil)
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheck)
	handler.ServeHTTP(response,request)
		
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println("response got is "+ response.Body.String())

	expected := `API is up and running`
	if response.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
		response.Body.String(), expected)
	}
}

func TestGetFruits(t *testing.T){
	
	//write test cases for api health
	request, _ := http.NewRequest("GET", "/getFruits", nil)
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(GetFruits)
	handler.ServeHTTP(response,request)
		
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println("response got is "+ response.Body.String())

	// expected := `API is up and running`
	// if response.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 	response.Body.String(), expected)
	// }
}
