package api
import (
	"net/http"
)

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /capture", captureRequestFromFicmart)
	mux.HandleFunc("POST /authorize", authorizationRequestFromFicmart)
	mux.HandleFunc("POST /void", voidRequestFromFicmart)
	mux.HandleFunc("POST /refund", refundRequestFromFicmart)
	mux.HandleFunc("GET /payments/{order_id}", getPaymentStatus)
	mux.HandleFunc("GET /customers/{customer_id}/payments", getPaymentHistory)

	return mux
}