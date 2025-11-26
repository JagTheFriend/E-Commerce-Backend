package products

import (
	"net/http"

	"github.com/JagTheFriend/ecommerce/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products := []string{"Product 1", "Product 2", "Product 3"}
	json.Write(w, 200, products)
}
