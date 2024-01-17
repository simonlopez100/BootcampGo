package handler

import (
	"app/scaffolding/platform/web/request"
	"app/structured-api/internal"
	"app/structured-api/platform/web/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// 1- Se crea una estructura pensada para guardar los prodcutos

// DefaultProducts is an implementation with handlers for Products storage
type DefaultProducts struct {
	// sv is a product service
	sv internal.ProductService
}

// 2- Se crea una funcion para instanciar dicha estructura

// NewDefaultProducts creates a new DefaultProducts instance
func NewDefaultProducts(sv internal.ProductService) *DefaultProducts {
	return &DefaultProducts{
		sv: sv,
	}
}

// 3- Se crea una funcion para guardar una producto en formato JSON desde el request
type BodyRequestProductJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// 4- Se crea una estructura para guardar la respuesta del producto
type ProductJSONResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (d *DefaultProducts) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// A- REQUEST
		// validate if client request (unstructured) adapts to the structure data
		// This validation must be made at dynamic level

		// parse request body
		var body BodyRequestProductJSON
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// B- PROCESS
		// serialize internal product
		product := internal.Product{
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		if err := d.sv.Save(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrFieldRequired):
				// w.Header().Set("Content-Type", "text/plain")
				// w.WriteHeader(http.StatusBadRequest)
				// w.Write([]byte("invalid request body"))
				response.Text(w, http.StatusBadRequest, "invalid body")
			case errors.Is(err, internal.ErrProductAlreadyExists):
				// w.Header().Set("Content-Type", "text/plain")
				// w.WriteHeader(http.StatusBadRequest)
				// w.Write([]byte("product already exists"))
				response.Text(w, http.StatusBadRequest, "product already exists")
			default:
				// w.Header().Set("Content-Type", "text/plain")
				// w.WriteHeader(http.StatusInternalServerError)
				// w.Write([]byte("internal server error"))
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// C- RESPONSE
		// serialize movie response
		dataproductResponse := ProductJSONResponse{
			Id:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusCreated)
		// json.NewEncoder(w).Encode(map[string]any{
		// 	"message": "Product created successfully",
		// 	"data":    dataproductResponse,
		// })
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Product created",
			"data":    dataproductResponse,
		})
	}
}

// get product by id
func (d *DefaultProducts) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// get id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// get product by id
		product, err := d.sv.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		//response
		// serialize product response
		dataproductResponse := ProductJSONResponse{
			Id:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Product found",
			"data":    dataproductResponse,
		})
	}
}

func (d *DefaultProducts) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// validate key existance

		// to save the data flow in bytes so we can reuse it later
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		var bodyMap map[string]any
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		if err := ValidateKeyExistance(bodyMap, "name", "quantity", "code_value", "is_published", "expiration", "price"); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// get body from request
		var body BodyRequestProductJSON
		if err := json.Unmarshal(bytes, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}
		// process
		// serialize internal product
		product := internal.Product{
			Id:          id,
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}
		// update product
		if err := d.sv.Update(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrFieldRequired):
				response.Text(w, http.StatusBadRequest, "invalid body")
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}
		// response
		// serialize product response
		dataproductResponse := ProductJSONResponse{
			Id:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Product updated",
			"data":    dataproductResponse,
		})

	}
}

func (d *DefaultProducts) UpdatePartial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// get the movide from service
		product, err := d.sv.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}
		// serialize internal product
		reqBody := BodyRequestProductJSON{
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		// get body from request
		if err := request.JSON(r, &reqBody); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// response
		data := ProductJSONResponse{
			Id:          product.Id,
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  reqBody.Expiration,
			Price:       reqBody.Price,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Product updated",
			"data":    data,
		})

	}
}

// Delete product by id
func (d *DefaultProducts) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		if err := d.sv.Delete(id); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "Product deleted",
		})
	}
}

// validate key existance
func ValidateKeyExistance(mp map[string]any, keys ...string) (err error) {
	for _, key := range keys {
		if _, ok := mp[key]; !ok {
			return fmt.Errorf("key %s not found", key)
		}
	}
	return
}
