package products

import (
	"log/slog"
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
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		slog.Error("Failed to list products", "error", err)
		json.Write(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	json.Write(w, 200, products)
}
