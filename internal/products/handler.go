package products

import (
	"log/slog"
	"net/http"
	"strconv"

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

func (h *handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	id := vars.Get("id")
	if id == "" {
		slog.Error("Failed to get product id", "error", "id is empty")
		json.Write(w, http.StatusBadRequest, "Bad Request")
		return
	}

	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		slog.Error("Failed to parse product id", "error", err)
		json.Write(w, http.StatusBadRequest, "Bad Request")
		return
	}

	product, err := h.service.GetProductByID(r.Context(), parsedId)
	if err != nil {
		slog.Error("Failed to get product by id", "error", err)
		json.Write(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	json.Write(w, 200, product)
}
