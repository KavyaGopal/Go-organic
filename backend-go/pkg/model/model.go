package model

// Product Struct (Model)
type ProductMock struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductMockUpdate struct {
	ID       string   `json:"productId"`
	Name     string   `json:"productName"`
	Category []string `json:"productCategory"`
}

type CategoryMock struct {
	ID          string `json:"categoryId"`
	Name        string `json:"categoryName"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Rating      string `json:"rating"`
}
