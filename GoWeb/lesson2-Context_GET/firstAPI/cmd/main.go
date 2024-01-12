package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// Global products map to storage serialize and save products
var globalProducts = make(map[int]Product)

// load products from file to slice
func LoadProducts() ([]Product, error) {
	// Using os library to read file
	content, err := os.ReadFile("products.json")
	//Check for errors in function read file
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}
	// Create a slice of type Product
	var products []Product
	// convert []byte content to json and save it in products []Product
	if err := json.Unmarshal(content, &products); err != nil {
		return nil, fmt.Errorf("Error conversion json: %v", err)
	}
	// return products in slice
	return products, nil
}

// load products to map [id, Product]
func SaveProductsToMemory(productsFromFile []Product) {
	for _, product := range productsFromFile {
		globalProducts[product.Id] = product
	}
}

// Endpoint to get pong
func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong!"))
}

// Endpoint to get all the products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Create a slice to store global products
	productList := make([]Product, 0, len(globalProducts))
	// Append products to slice
	for _, product := range globalProducts {
		productList = append(productList, product)
	}
	//set headers
	w.Header().Set("Content-Type", "application/json")
	//set status
	w.WriteHeader(http.StatusOK)
	// return json of product list
	json.NewEncoder(w).Encode(productList)
}

// Enpoint to get product by id
func GetProductById(w http.ResponseWriter, r *http.Request) {

	//  get id param from path
	productId := chi.URLParam(r, "productId")
	// convert id from patch param (string to int)
	productID, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Error convertion string to float", http.StatusBadRequest)
	}
	// If id does not exist!
	if productID > len(globalProducts) || productID < 0 {
		w.Write([]byte("User not found"))
		return
	}
	// Return product by id un global map
	productById := globalProducts[productID]

	//set headers
	w.Header().Set("Content-Type", "application/json")
	//set status
	w.WriteHeader(http.StatusOK)
	// return json of product list
	json.NewEncoder(w).Encode(productById)
}

// Endpoint to search by query params
func GetProductsByQueryParams(w http.ResponseWriter, r *http.Request) {
	// Create a slice to store global products
	productList := make([]Product, 0, len(globalProducts))
	// Append products to slice
	for _, product := range globalProducts {
		productList = append(productList, product)
	}
	// create a slice to list filtered products
	var filteredProducts []Product
	// get query param price from path
	price := r.URL.Query().Get("price")
	// conver string to float
	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		http.Error(w, "Error convertion string to float", http.StatusBadRequest)
	}
	// Filter product with price greater than query param price
	for _, product := range productList {
		if product.Price > priceFloat {
			filteredProducts = append(filteredProducts, product)
		}
	}
	if len(filteredProducts) == 0 {
		http.Error(w, "No products available", http.StatusBadRequest)
	}

	// See convertion in console
	fmt.Println(priceFloat)
	//set headers
	w.Header().Set("Content-Type", "application/json")
	//set status
	w.WriteHeader(http.StatusOK)
	// return json of product list
	json.NewEncoder(w).Encode(filteredProducts)
}

func main() {

	productsList, err := LoadProducts()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	SaveProductsToMemory(productsList)

	// Create a new router
	rt := chi.NewRouter()

	// Endpoint to get pong
	rt.Get("/ping", ping)

	rt.Route("/products", func(rt chi.Router) {
		// Endpoint to get all products
		rt.Get("/", GetProducts)
		// Endpoint to get a product by id
		rt.Get("/{productId}", GetProductById)
		// Endpoint to search products by query params
		rt.Get("/search", GetProductsByQueryParams)
	})

	// Start server
	if err := http.ListenAndServe(":8000", rt); err != nil {
		panic(err)
	}

}
