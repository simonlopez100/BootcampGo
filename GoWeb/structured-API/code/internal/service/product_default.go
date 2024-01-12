package service

import (
	"app/structured-api/internal"
	"fmt"
)

// is a struct that represents the default implementation of a product service
type ProductDefault struct {
	rp internal.ProductRepository
}

func NewProductDefault(rp internal.ProductRepository) *ProductDefault {
	return &ProductDefault{
		rp: rp,
	}
}

// Save saves a movie
func (pd *ProductDefault) Save(product *internal.Product) (err error) {

	// bussines logic

	// validation for name != empty
	if (*product).Name == "" {
		return fmt.Errorf("%w: name", internal.ErrFieldRequired)
	}
	// validation for name != empty
	if (*product).Quantity == 0 {
		return fmt.Errorf("%w: quantity", internal.ErrFieldRequired)
	}
	// validation for name != empty
	if (*product).CodeValue == "" {
		return fmt.Errorf("%w: code value", internal.ErrFieldRequired)
	}

	// save product in the repository
	err = pd.rp.Save(product)
	if err != nil {
		switch err {
		case internal.ErrProductNameAlreadyExists:
			return fmt.Errorf("%w: name", internal.ErrProductAlreadyExists)
		}
		return
	}

	return

}
