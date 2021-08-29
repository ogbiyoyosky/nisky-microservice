package data

import "time"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func GetProduct() []*Product {
	return productList
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "AirPod Pro",
		Description: "Apple AirPod Pro",
		Price:       2.45,
		SKU:         "app-1",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Air Jordan 1",
		Description: "Air Jordan",
		Price:       2.95,
		SKU:         "aj-1",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
