package application

import (
	"app/structured-api/internal"
	"app/structured-api/internal/handler"
	"app/structured-api/internal/repository"
	"app/structured-api/internal/service"
	"net/http"

	"github.com/go-chi/chi"
)

type ServerChi struct {
	// address is the address to listen to
	adress string
}

func NewServerChi(adress string) *ServerChi {
	defaultAdress := ":8080"
	if adress != "" {
		defaultAdress = adress
	}

	return &ServerChi{
		adress: defaultAdress,
	}
}

func (s *ServerChi) Run() (err error) {
	// Initialize Dependencies
	// db - repository
	products := make(map[int]internal.Product)
	lastID := 0
	rp := repository.NewProductMap(products, lastID)

	// service
	sv := service.NewProductDefault(rp)

	// handler
	hd := handler.NewDefaultProducts(sv)

	// router
	rt := chi.NewRouter()

	// endpoints
	rt.Get("/products/{id}", hd.GetByID())
	rt.Post("/products", hd.Create())
	rt.Put("/products/{id}", hd.Update())
	rt.Patch("/products/{id}", hd.UpdatePartial())

	// run server
	return http.ListenAndServe(s.adress, rt)

}
