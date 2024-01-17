package internal

import "errors"

var (
	ErrProductNameAlreadyExists = errors.New("product name already exist")
	ErrProductNotFound          = errors.New("product not found")
)

type ProductRepository interface {
	// Save saves a product in the repository
	Save(product *Product) (err error)
	// GetById gets a product by its ID
	GetById(id int) (product Product, err error)
	// Update updates a product in the repository
	Update(product *Product) (err error)
	// Delete deletes a product from the repository
	Delete(id int) (err error)
}
