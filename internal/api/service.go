package api

import (
	"mart-gateway/internal/models"
)

//what can handler ask service to do

type HandlerInterface interface {
	AuthorizePayment(req models.Martrequest, key string) (models.Payment,error)
	CapturePayment(payment_id string, key string) (models.Payment, error)
	VoidPayment(payment_id string, key string) (models.Payment, error)
	RefundPayment(payment_id string, key string,) (models.Payment, error)
	GetPayment(payment_id string) (models.Payment, error)
	GetStatus(order_id string) (string, error)
	GetHistory(customerID string) ([]models.Payment, error)
}