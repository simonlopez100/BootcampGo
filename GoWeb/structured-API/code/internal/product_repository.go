package internal

import "errors"

var (
	ErrProductNameAlreadyExists = errors.New("product name already exists")
)

type ProductRepository interface {
	// Save saves a product in the repository
	Save(product *Product) (err error)
}
