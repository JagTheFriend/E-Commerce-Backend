package orders

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

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var tempOrder createOrderParams
	if err := json.Read(r, &tempOrder); err != nil {
		slog.Error("Error while reading request body", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdOrder, err := h.service.PlaceOrder(r.Context(), tempOrder)
	if err != nil {
		slog.Error("Error while placing order", "error", err)

		if err == ErrProductNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, createdOrder)
}
