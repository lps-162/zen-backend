package handlers

import (
	"encoding/json"
	"net/http"
	"zentools/models"
)

var orders = []models.Order{
	models.Order{OrderNumber: 10100, OrderDate: "2003-05-20",
		ShippedDate: "2003-05-25", Status: "Shipped"},
	models.Order{OrderNumber: 10100, OrderDate: "2003-05-21",
		ShippedDate: "2003-05-22", Status: "Shipped"},
	models.Order{OrderNumber: 10100, OrderDate: "2003-05-24",
		ShippedDate: "2003-05-29", Status: "Shipped"},
}

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)

}
