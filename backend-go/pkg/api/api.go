package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/KavyaGopal/Go-organic/backend-go/pkg/db"
	"github.com/KavyaGopal/Go-organic/backend-go/pkg/model"
	"github.com/KavyaGopal/Go-organic/backend-go/pkg/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/price"
	"github.com/stripe/stripe-go/v72/webhook"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

//init product variable for mock
var fruits []model.FruitMock
var snacks []model.SnacksMock
var vegetables []model.VegetablesMock
var groceries []model.GroceriesMock
var cosmetics []model.CosmeticsMock

type App struct {
	Router *mux.Router
}

func (a *App) Run(addr string) {}

func handleCors(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

//get all fruits
func GetFruits(w http.ResponseWriter, r *http.Request) {
	handleCors(&w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fruits)
	return
}

//get all snacks
func GetSnacks(w http.ResponseWriter, r *http.Request) {
	handleCors(&w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snacks)

}

//get all vegetables
func GetVegetables(w http.ResponseWriter, r *http.Request) {
	handleCors(&w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vegetables)

}

//get all cosmetics
func GetCosmetics(w http.ResponseWriter, r *http.Request) {
	handleCors(&w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cosmetics)
	return

}

//get all groceries
func GetGroceries(w http.ResponseWriter, r *http.Request) {
	handleCors(&w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groceries)

}

//adding health check
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, "API is up and running")
}

//get all the products from the database
func GetAllProductsFromDB(w http.ResponseWriter, r *http.Request) {

	log.Println("GetAllProductsFromDB function called")

	w.Header().Set("Content-Type", "application/json")
	handleCors(&w, r)
	var productMaster []model.ProdMaster
	//db query from sqlite
	setupDB.DB.Find(&productMaster)

	err := json.NewEncoder(w).Encode(productMaster)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//get filtered query as item categories from db
func GetFilteredCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	handleCors(&w, r)
	params := mux.Vars(r) // get the params
	log.Println("GetFilteredCategory function called for ", params)
	// var productJsonArray []model.ProdMasterUpdate
	var productMaster []model.ProdMaster
	//get the json array from function : getAllProductsFromDBUpdate
	setupDB.DB.Where("item_category=?", params["itemCategory"]).Find(&productMaster)

	json.NewEncoder(w).Encode(productMaster)
}

//register api
func RegisterUser(w http.ResponseWriter, r *http.Request) {

	log.Println("Register API Invoked ")

	w.Header().Set("Content-Type", "application/json")
	handleCors(&w, r)
	db, err := gorm.Open("sqlite3", "ProductData.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var user model.User

	err2 := decoder.Decode(&user)
	if err2 != nil {
		panic(err2)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	result := db.Create(&model.User{Name: user.Name, Email: user.Email, Address: user.Address, Password: string(hashedPassword), Phone: user.Phone})

	var jsonMessage model.JsonMessage

	if result.Error != nil {
		fmt.Println(result.Error)
		log.Println("Unauthorized Access ")
		jsonMessage = model.JsonMessage{Status: 500, Message: "User already exists"}
		json.NewEncoder(w).Encode(jsonMessage)
	} else if user.Email == "" {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Email is Empty")
		jsonMessage = model.JsonMessage{Status: http.StatusUnauthorized, Message: "User Email is empty"}
		json.NewEncoder(w).Encode(jsonMessage)
		return
	} else if user.Phone == "" {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Phone is Empty")
		jsonMessage = model.JsonMessage{Status: http.StatusUnauthorized, Message: "User Phone is empty"}
		json.NewEncoder(w).Encode(jsonMessage)
		return
	} else {
		jsonMessage = model.JsonMessage{Status: 200, Message: "New User Successfully Added: " + user.Name}
		json.NewEncoder(w).Encode(jsonMessage)
	}
}

//log in user
func LoginUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	handleCors(&w, r)
	db, err := gorm.Open("sqlite3", "ProductData.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var login model.Login

	err2 := decoder.Decode(&login)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if login.Email == "" || login.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user_ model.User
	// Get the existing entry present in the database for the given username
	db.Table("users").Where("Email = ?", login.Email).Find(&user_)

	var jsonMessage model.JsonMessage
	var jsonLoginResponse model.JsonLoginResponse

	if err != nil {
		// If there is an issue with the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Internal Server Error")
		jsonMessage = model.JsonMessage{Status: 500, Message: "Internal Server Error"}
		json.NewEncoder(w).Encode(jsonMessage)
		return
	}

	if user_.Email == "" {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Unauthorized Access ")
		jsonMessage = model.JsonMessage{Status: http.StatusUnauthorized, Message: "Unauthorized Access"}
		json.NewEncoder(w).Encode(jsonMessage)
		return
	}
	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(user_.Password), []byte(login.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Unauthorized Access ")
		jsonMessage = model.JsonMessage{Status: 500, Message: "Unauthorized Access"}
		json.NewEncoder(w).Encode(jsonMessage)

		return
	}

	log.Println("User successfully logged in ", user_.Name)
	// var jsonMessage = model.JsonMessage{200, "User Name "+ user_.Name+ "successfully logged in"}
	jsonLoginResponse = model.JsonLoginResponse{Name: user_.Name, Email: user_.Email, Status: 200, Message: "User successfully logged in"}
	json.NewEncoder(w).Encode(jsonLoginResponse)

}

//reset password api
func ResetPassword(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	handleCors(&w, r)
	db, err := gorm.Open("sqlite3", "ProductData.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	var login model.Login
	var jsonResetResponse model.JsonMessage

	err2 := decoder.Decode(&login)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if login.Email == "" || login.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user_ model.User

	db.Table("users").Where("Email = ?", login.Email).Find(&user_)
	if user_.Email == "" {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Unauthorized Access ")
		jsonResetResponse = model.JsonMessage{Status: http.StatusUnauthorized, Message: "Unauthorized Access"}
		json.NewEncoder(w).Encode(jsonResetResponse)
		return
	}
	if login.Password == user_.Password {
		jsonResetResponse = model.JsonMessage{Status: 500, Message: "Password is already matched with the database"}
		json.NewEncoder(w).Encode(jsonResetResponse)
		return
	}

	// Get the existing entry present in the database for the given username

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(login.Password), 8)
	db.Model(&user_).Update("password", hashedPassword)
	jsonResetResponse = model.JsonMessage{Status: 200, Message: "Password successfully updated."}
	json.NewEncoder(w).Encode(jsonResetResponse)

}

func FetchItemQuantity(w http.ResponseWriter, r *http.Request) {
	var items []model.Item
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Internal Server Error")
		jsonMessage := model.JsonMessage{Status: 500, Message: "Internal Server Error"}
		json.NewEncoder(w).Encode(jsonMessage)
		return
	}
	json.Unmarshal(body, &items)

	itemsProd := make([]model.ProdMaster, len(items))

	for i, x := range items {
		id := x.ItemID
		quantity := x.ItemQuantity

		setupDB.DB.First(&itemsProd[i], id)
		if itemsProd[i].ItemInventory < quantity {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Internal Server Error")
			jsonMessage := model.JsonMessage{Status: 500, Message: "Item " + itemsProd[i].ItemName + " is not available for selected quantity"}
			json.NewEncoder(w).Encode(jsonMessage)
			return

		}

	}

	for i, x := range items {
		id := x.ItemID

		quantity := x.ItemQuantity
		var product model.ProdMaster
		setupDB.DB.First(&product, id)
		setupDB.DB.Model(&product).Update("item_inventory", itemsProd[i].ItemInventory-quantity)

	}

	jsonMessage := model.JsonMessage{Status: 200, Message: "Checkout of All Items Done Successfully"}
	json.NewEncoder(w).Encode(jsonMessage)

}

//get all user testimonials from the database
func GetUserTestimonials(w http.ResponseWriter, r *http.Request) {

	log.Println("GetUserTestimonials is called")

	w.Header().Set("Content-Type", "application/json")
	handleCors(&w, r)
	var userTestimonialMaster []model.UserTestimonial
	//db query from sqlite
	setupDB.DB.Find(&userTestimonialMaster)

	err := json.NewEncoder(w).Encode(userTestimonialMaster)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//payment integration 
func handleConfig(w http.ResponseWriter, r *http.Request) {

	log.Printf("handle config called")

	w.Header().Set("Content-Type", "application/json")
	handleCors(&w, r)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	// Fetch a price, use it's unit amount and currency
	p, _ := price.Get(
		os.Getenv("PRICE"),
		nil,
	)
	utils.WriteJSON(w, struct {
		PublicKey  string `json:"publicKey"`
		UnitAmount int64  `json:"unitAmount"`
		Currency   string `json:"currency"`
	}{
		PublicKey:  os.Getenv("STRIPE_PUBLISHABLE_KEY"),
		UnitAmount: p.UnitAmount,
		Currency:   string(p.Currency),
	})
}

//this is a server which can check if payment transaction is initiated
// and return secret key
func handleWebhook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	handleCors(&w, r)
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("ioutil.ReadAll: %v", err)
		return
	}

	event, err := webhook.ConstructEvent(b, r.Header.Get("Stripe-Signature"), os.Getenv("STRIPE_WEBHOOK_SECRET"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("webhook.ConstructEvent: %v", err)
		return
	}

	if event.Type == "checkout.session.completed" {
		fmt.Println("Checkout Session completed!")
	}

	utils.WriteJSON(w, nil)
}

//check the environment of the .env file
func checkEnv() {
	price := os.Getenv("PRICE")
	if price == "price_12345" || price == "" {
		log.Fatal("You must set a Price ID from your Stripe account. See the README for instructions.")
	}
}

//create the session for the user on the server
func handleCheckoutSession(w http.ResponseWriter, r *http.Request) {
	log.Printf("handle checkout session called")
	w.Header().Set("Content-Type", "application/json")
	handleCors(&w, r)
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	sessionID := r.URL.Query().Get("sessionId")
	s, _ := session.Get(sessionID, nil)
	utils.WriteJSON(w, s)
}

//this is first invoked when stripe checkout is called
func handleCreateCheckoutSession(w http.ResponseWriter, r *http.Request) {

	log.Printf("handle create checkout session called")
	w.Header().Set("Content-Type", "application/json")
	handleCors(&w,r)

	domainURL := os.Getenv("DOMAIN")

	params := &stripe.CheckoutSessionParams{
		SuccessURL:         stripe.String(domainURL + "/checkout-success"),
		CancelURL:          stripe.String(domainURL + "/checkout-cancel"),
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: 			GetLineItems(w,r),
		
	}
	s, err := session.New(params)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while creating session %v", err.Error()), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}

//fetch the product keys of stripe corresponding to the product id
func GetStripeProductKeys() (productKeyJson string){

	log.Println("GetProductIDMappingList function called")
	
	var productIDMaster []model.ProductIDMaster
	//db query from sqlite
	setupDB.DB.Find(&productIDMaster)

	u, err := json.Marshal(productIDMaster)
	if err != nil {
		panic(err)
	}
	
	return string(u)
}

//get the LineItems object for the stripe payment
func GetLineItems(w http.ResponseWriter, r *http.Request) ([]*stripe.CheckoutSessionLineItemParams){

    log.Printf("GetLineItems called")

	//process items first

	var items []model.Item

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Internal Server Error")
		jsonMessage := model.JsonMessage{Status: 500, Message: "Internal Server Error"}
		json.NewEncoder(w).Encode(jsonMessage)
	}

	json.Unmarshal(body, &items)

	// items processing ends here

	//process productkey list from getstripe

	var productKeyList []model.ProductIDMaster

    var productKeys string
    productKeys = GetStripeProductKeys()

    err1:= json.Unmarshal([]byte(productKeys), &productKeyList)

    if err1 != nil{
        panic(err1)
    }

	//object to return
	var lineItems []*stripe.CheckoutSessionLineItemParams

	//filter according to the item ID and item quantity
	//need to iterate through the items 

	for _, x := range items {

		for _, productKeyData := range productKeyList {

			if productKeyData.ID == x.ItemID {
	
				var lineItem stripe.CheckoutSessionLineItemParams

				lineItem.Price = stripe.String(productKeyData.Key)
				lineItem.Quantity = stripe.Int64(x.ItemQuantity)

				lineItems = append(lineItems,&lineItem)
			
			}
	
			
		}

	}

    return lineItems

}

func main() {
	//init router
	a := &App{}

	a.Router = mux.NewRouter()
	setupDB.ConnectDatabase()

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

	// payment api
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	checkEnv()

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	//need to handle the mux compatibility as well
	http.Handle("/", http.FileServer(http.Dir(os.Getenv("STATIC_DIR"))))
	http.HandleFunc("/config", handleConfig)
	http.HandleFunc("/webhook", handleWebhook)

	a.Router.HandleFunc("/getFruits", GetFruits).Methods("GET")
	a.Router.HandleFunc("/getSnacks", GetSnacks).Methods("GET")
	a.Router.HandleFunc("/getVegetables", GetVegetables).Methods("GET")
	a.Router.HandleFunc("/getCosmetics", GetCosmetics).Methods("GET")
	a.Router.HandleFunc("/getGroceries", GetGroceries).Methods("GET")

	//register api endpoint
	a.Router.HandleFunc("/registerUser", RegisterUser).Methods("POST")
	//login api endpoint
	a.Router.HandleFunc("/loginUser", LoginUser).Methods("POST")
	a.Router.HandleFunc("/resetPassword", ResetPassword).Methods("POST")
	//add apis to fetch data from db
	a.Router.HandleFunc("/api/fetchAllProductsFromDB", GetAllProductsFromDB).Methods("GET")
	a.Router.HandleFunc("/api/fetchProduct/{itemCategory}", GetFilteredCategory).Methods("GET")
	a.Router.HandleFunc("/api/fetchUserTestimonials", GetUserTestimonials).Methods("GET")

	a.Router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	a.Router.HandleFunc("/api/fetchItemQuantity", FetchItemQuantity).Methods("POST")
	// http.Handle("/", a.Router)

	// log.Fatal(http.ListenAndServe(":8000", a.Router))
	// a.Run(":8010")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(a.Router)))

}
