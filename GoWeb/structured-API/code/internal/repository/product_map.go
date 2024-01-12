package repository

import (
	"app/structured-api/internal"
)

type ProductMap struct {
	// db is a map of [product ID]Product that represent a database
	db map[int]internal.Product

	// Key: id of the product
	// Value: product
	lastId int
}

// NewProductMap creates a new ProductMap instance
func NewProductMap(db map[int]internal.Product, lastId int) *ProductMap {
	return &ProductMap{
		db:     db,
		lastId: lastId,
	}
}

func (pm *ProductMap) Save(product *internal.Product) (err error) {
	// Validate product (consistency)
	//  title unic
	for _, v := range (*pm).db {
		if v.Name == product.Name {
			return internal.ErrProductNameAlreadyExists
		}
	}

	// id autroincrement
	(*pm).lastId++
	// set id
	(*product).Id = (*pm).lastId
	// save product
	(*pm).db[(*product).Id] = *product

	return
}
