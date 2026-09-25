package api
import (
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /capture", h.captureRequestFromFicmart)
	mux.HandleFunc("POST /authorize", h.authorizationRequestFromFicmart)
	mux.HandleFunc("POST /void", h.voidRequestFromFicmart)
	mux.HandleFunc("POST /refund", h.refundRequestFromFicmart)
	mux.HandleFunc("GET /payments/{order_id}", h.getPaymentStatus)
	mux.HandleFunc("GET /customers/{customer_id}/payments", h.getPaymentHistory)

	return mux
}