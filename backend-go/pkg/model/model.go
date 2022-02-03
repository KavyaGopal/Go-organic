package model

// Product Struct (Model)
type Product struct{
	ID   string `json:"id"`
	Name string `json:"name"`
}


// func main() {
// 	//init router
// 	r := mux.NewRouter()

// 	//Mock data TODO: get the data from the sqlite
// 	products = append(products, Product{ID: "1", Name: "Groceries"})
// 	products = append(products, Product{ID: "2", Name: "Fruits"})
// 	products = append(products, Product{ID: "3", Name: "Snacks"})
// 	products = append(products, Product{ID: "4", Name: "Vegetables"})
// 	products = append(products, Product{ID: "5", Name: "Dairy"})

// 	r.HandleFunc("/api/products", getProducts).Methods("GET")
// 	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")

// 	log.Fatal(http.ListenAndServe(":8000",r))

// }