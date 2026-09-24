package api
import (
	"net/http"
)

type Handler struct {
	service HandlerInterface
}

func NewHandler(service HandlerInterface) *Handler{
	return &Handler{service: service}
}

//handlers for capture, void, refund, authorization
func authorizationRequestFromFicmart(w http.ResponseWriter, r *http.Request) {

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