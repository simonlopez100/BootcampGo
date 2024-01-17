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

// ValidateProduct validates a product (business logic)
func ValidateProduct(product *internal.Product) (err error) {
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
	return
}

// Save saves a movie
func (pd *ProductDefault) Save(product *internal.Product) (err error) {

	// bussines logic
	// validation
	if err = ValidateProduct(product); err != nil {
		return
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

// GetByID gets a product by its ID
func (pd *ProductDefault) GetByID(id int) (product internal.Product, err error) {
	product, err = pd.rp.GetById(id)
	if err != nil {
		switch err {
		case internal.ErrProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrProductNotFound)
		}
		return
	}
	return
}

// Update updates a product
func (pd *ProductDefault) Update(product *internal.Product) (err error) {
	if err = ValidateProduct(product); err != nil {
		return
	}

	// update product in the repository
	err = pd.rp.Update(product)
	if err != nil {
		switch err {
		case internal.ErrProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrProductNotFound)
		}
		return
	}
	return
}

// Delete deletes a product
func (pd *ProductDefault) Delete(id int) (err error) {
	err = pd.rp.Delete(id)
	if err != nil {
		switch err {
		case internal.ErrProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrProductNotFound)
		}
		return
	}
	return
}
