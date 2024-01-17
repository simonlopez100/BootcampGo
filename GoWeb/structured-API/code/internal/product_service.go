package internal

import "errors"

var (
	ErrFieldRequired        = errors.New("field is required")
	ErrProductAlreadyExists = errors.New("product already exists")
)

// ProductService is an interface that represents the product service
// - business logic
// - validation
// - external services (apis, dbs...)
type ProductService interface {
	// Save saves a product
	Save(product *Product) (err error)
	// GetByID gets a product by its ID
	GetByID(id int) (product Product, err error)
	// Update updates a product
	Update(product *Product) (err error)
	// Delete deletes a product
	Delete(id int) (err error)
}
