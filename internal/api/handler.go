package api

import (
	"encoding/json"
	"mart-gateway/internal/models"
	"net/http"
)

type Handler struct {
	service HandlerInterface
}

func NewHandler(service HandlerInterface) *Handler{
	return &Handler{service: service}
}

//handlers for capture, void, refund, authorization
func (h *Handler) authorizationRequestFromFicmart(w http.ResponseWriter, r *http.Request) {
	var req models.Martrequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return 
	}

	key := r.Header.Get("Idempotency-Key")
	if key == ""{
		http.Error(w, "missing idempotency key", http.StatusBadRequest)
		return
	}

	payment, err := h.service.AuthorizePayment(req, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payment)

}

func captureRequestFromFicmart(w http.ResponseWriter, r *http.Request) {

}
func voidRequestFromFicmart(w http.ResponseWriter, r *http.Request) {

}

func refundRequestFromFicmart(w http.ResponseWriter, r *http.Request) {

}

func getPaymentStatus(w http.ResponseWriter, r *http.Request) {

}

func getPaymentHistory(w http.ResponseWriter, r *http.Request) {

}