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
}
