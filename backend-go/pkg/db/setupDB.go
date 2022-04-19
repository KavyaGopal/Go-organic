package setupDB

import (
	"github.com/KavyaGopal/Go-organic/backend-go/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("ProductData.db"))

	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
}

func ConnectEndPointDatabase() {
	database, err := gorm.Open(sqlite.Open("pkg/api/ProductData.db"))

	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
}

func CreateDatabase() {

	db, err := gorm.Open(sqlite.Open("pkg/api/ProductData.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.ProdMaster{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserTestimonial{})
	db.AutoMigrate(&model.ProductIDMaster{})

	var products = []model.ProdMaster{
		//fruits
		{ImageSource: "../../../assets/items/apple.png", ItemName: "Apple", ItemCategory: "Fruits", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 12},
		{ImageSource: "../../../assets/items/cherry.png", ItemName: "Cherry", ItemCategory: "Fruits", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 10},
		{ImageSource: "../../../assets/items/orange.png", ItemName: "Orange", ItemCategory: "Fruits", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 9},
		{ImageSource: "../../../assets/items/pineapple.jpeg", ItemName: "Pineapple", ItemCategory: "Fruits", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 9},
		{ImageSource: "../../../assets/items/jackfruit.jpeg", ItemName: "Jackfruit", ItemCategory: "Fruits", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 9},
		{ImageSource: "../../../assets/items/watermelon.jpeg", ItemName: "Watermelon", ItemCategory: "Fruits", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 9},
		//vegetables
		{ImageSource: "../../../assets/items/potato.jpeg", ItemName: "Potato", ItemCategory: "Vegetables", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 10},
		{ImageSource: "../../../assets/items/tomato.jpeg", ItemName: "Tomato", ItemCategory: "Vegetables", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 9},
		{ImageSource: "../../../assets/items/cauliflower.jpeg", ItemName: "Cauliflower", ItemCategory: "Vegetables", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 10},
		{ImageSource: "../../../assets/items/brinjal.jpeg", ItemName: "Brinjal", ItemCategory: "Vegetables", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 9},
		{ImageSource: "../../../assets/items/onion.jpeg", ItemName: "Onion", ItemCategory: "Vegetables", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 8},
		{ImageSource: "../../../assets/items/carrot.jpeg", ItemName: "Carrot", ItemCategory: "Vegetables", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 8},
		//snacks
		{ImageSource: "../../../assets/items/cashew.jpeg", ItemName: "Roasted Cashew", ItemCategory: "Snacks", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 10},
		{ImageSource: "../../../assets/items/peanuts.jpeg", ItemName: "Roasted Peanut", ItemCategory: "Snacks", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 11},
		{ImageSource: "../../../assets/items/riceCake.jpeg", ItemName: "Carrot Sticks", ItemCategory: "Snacks", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 13},
		{ImageSource: "../../../assets/items/mango.jpeg", ItemName: "Dried Mango", ItemCategory: "Snacks", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 9},
		{ImageSource: "../../../assets/items/mix.jpeg", ItemName: "Trial Mix", ItemCategory: "Snacks", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 8},
		{ImageSource: "../../../assets/items/riceCake.jpeg", ItemName: "Rice cake", ItemCategory: "Snacks", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 8},
		//cosmetics
		{ImageSource: "../../../assets/items/cleanser.jpeg", ItemName: "Cleanser", ItemCategory: "Cosmetics", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 12},
		{ImageSource: "../../../assets/items/bodyBalm.jpeg", ItemName: "Body Balm", ItemCategory: "Cosmetics", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 12},
		{ImageSource: "../../../assets/items/oil.jpeg", ItemName: "Hair Oil", ItemCategory: "Cosmetics", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 13},
		{ImageSource: "../../../assets/items/faceMask.jpeg", ItemName: "Face Mask", ItemCategory: "Cosmetics", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 10},
		{ImageSource: "../../../assets/items/kajal.jpeg", ItemName: "Kajal", ItemCategory: "Cosmetics", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 8},
		{ImageSource: "../../../assets/items/soap.jpeg", ItemName: "Soap", ItemCategory: "Cosmetics", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 8},
		//groceries
		{ImageSource: "../../../assets/items/rice.jpeg", ItemName: "Rice", ItemCategory: "Groceries", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 10},
		{ImageSource: "../../../assets/items/wheat.jpeg", ItemName: "Wheat", ItemCategory: "Groceries", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 9},
		{ImageSource: "../../../assets/items/sugar.jpeg", ItemName: "Sugar", ItemCategory: "Groceries", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 10},
		{ImageSource: "../../../assets/items/sugar.jpeg", ItemName: "Salt", ItemCategory: "Groceries", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 9},
		{ImageSource: "../../../assets/items/chilli.png", ItemName: "Chilli Powder", ItemCategory: "Groceries", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 8},
		{ImageSource: "../../../assets/items/chilli.png", ItemName: "Garam Masala", ItemCategory: "Groceries", ItemDesc: "This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.", ItemWeight: 500, ItemQuantity: 1, ItemInventory: 200, ItemCost: 8}}
	
		db.Create(&products)

		var testimonials = []model.UserTestimonial{
            
            {ImageSource: "../../../assets/testimonial-images/img-m1.png", UserDescription: "I am in my mid thirties and my weight gain hit me hard. I have tried every diet there is but instead of dieting now...I ordered go-organic products and who knew this was too good for you! Now I eat and truly enjoy my food."},
            {ImageSource: "../../../assets/testimonial-images/img-fm1.jpeg", UserDescription: "These products are easy and super healthy. The choices are great. I would highly recommend to use go-organic products to my friends and first timers."},
            {ImageSource: "../../../assets/testimonial-images/img-m2.jpeg", UserDescription: "I've been dabbling with the Vegan diet on and off since the beginning of the year. These products have really given me inspiration in terms of meal plans and different recipes to try."},
            {ImageSource: "../../../assets/testimonial-images/img-fm2.jpeg", UserDescription: "I bought this to refresh my diet and nutrition and wanted some new healthy diets. Every single thing I have tried has been a winner and products are relatively cheaper."},
            {ImageSource: "../../../assets/testimonial-images/img-m3.jpeg", UserDescription: "Mostly tasty products. I am not a vegan and enjoyed most of the go-organic products. I highly recommended these products if you want to become plant eater and not feel hungry."},
            {ImageSource: "../../../assets/testimonial-images/img-fm3.jpeg",UserDescription: "Being healthy is in my genes and I highly appreciate that go-organic is promoting healty diets in the form of organic based products."}}
        
        db.Create(&testimonials)
		
		var productIdMapping = []model.ProductIDMaster{
            //fruits
            {Key: "price_1Kpg3AEdowRWU1QhLGGU9OJ0"},
            {Key: "price_1KpgN3EdowRWU1QhOKpvaaVr"},
			{Key: "price_1KpgNREdowRWU1Qh2ukoR2yT"},
			{Key: "price_1KpgNtEdowRWU1QhMjbadG77"},
			{Key: "price_1KpgOHEdowRWU1Qhs4ajwrYY"},
			{Key: "price_1KpgOoEdowRWU1QhjNPZamXT"},
			//vegetables
			{Key: "price_1Kpg3zEdowRWU1QhM01asQ1Q"},
            {Key: "price_1KpgPJEdowRWU1Qh0U398QzW"},
			{Key: "price_1KpgQ3EdowRWU1QhtWaCdKrk"},
			{Key: "price_1KpgQOEdowRWU1QhSiZVE9XV"},
			{Key: "price_1KpgQlEdowRWU1Qh4MGl5f1r"},
			{Key: "price_1KpgRBEdowRWU1QhzItAHXkc"},
			//snacks
			{Key: "price_1KpgReEdowRWU1QhAMWJeNiT"},
            {Key: "price_1KpgS4EdowRWU1Qh3dL5nbp9"},
			{Key: "price_1KpgSSEdowRWU1QhuvjmlUNg"},
			{Key: "price_1KpgSrEdowRWU1QhRPAYhv1C"},
			{Key: "price_1KpgTNEdowRWU1QhZYwE6FJS"},
			{Key: "price_1KpgTrEdowRWU1Qh9ONXiqXw"},
			//cosmetics
			{Key: "price_1KpgUDEdowRWU1QhBJxNS1rV"},
            {Key: "price_1KpgUXEdowRWU1QhS3DzPKdQ"},
			{Key: "price_1KpgUuEdowRWU1QhEBNpmpPs"},
			{Key: "price_1KpgVEEdowRWU1QhUZmlToJs"},
			{Key: "price_1KpgVVEdowRWU1Qh9vah3xX3"},
			{Key: "price_1KpgVtEdowRWU1QhjKyOBX2j"},
			//groceries
			{Key: "price_1KpgWPEdowRWU1QhU5h4dvIL"},
            {Key: "price_1KpgWlEdowRWU1QhUqXdoLZL"},
			{Key: "price_1KpgX9EdowRWU1QhQvt4RPCc"},
			{Key: "price_1KpgXQEdowRWU1QhqAuGMF3y"},
			{Key: "price_1KpgXmEdowRWU1QhaAAJficP"},
			{Key: "price_1KpgY3EdowRWU1QhRblhZSQi"}}
        
        db.Create(&productIdMapping)
		
}
