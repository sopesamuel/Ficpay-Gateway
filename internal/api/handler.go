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
func (h *Handler) authorizationRequestFromFicmart(w http.ResponseWriter, r *http.Request){
	var req models.Martrequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	key := r.Header.Get("Idempotency-key")
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



func (h *Handler) captureRequestFromFicmart(w http.ResponseWriter, r *http.Request) {
	var req models.PaymentRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	key := r.Header.Get("Idempotency-key")
	if key == ""{
		http.Error(w, "missing idempotency key", http.StatusBadRequest)
		return
	}



	capturedpayment, err := h.service.CapturePayment(req.PaymentID, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(capturedpayment)
}


func (h *Handler) voidRequestFromFicmart(w http.ResponseWriter, r *http.Request) {
	var req models.PaymentRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	key := r.Header.Get("Idempotency-key")
	if key == ""{
		http.Error(w, "missing idempotency key", http.StatusBadRequest)
		return
	}



	voidedpayment, err := h.service.VoidPayment(req.PaymentID, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(voidedpayment)
}

func (h *Handler) refundRequestFromFicmart(w http.ResponseWriter, r *http.Request) {
	var req models.PaymentRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	key := r.Header.Get("Idempotency-key")
	if key == ""{
		http.Error(w, "missing idempotency key", http.StatusBadRequest)
		return
	}



	refundedpayment, err := h.service.RefundPayment(req.PaymentID, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(refundedpayment)
}

func (h *Handler) getPaymentStatus(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getPaymentHistory(w http.ResponseWriter, r *http.Request) {

}