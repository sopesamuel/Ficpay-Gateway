package service 

import (
	"mart-gateway/internal/models"
)

type RepositoryInterface interface {
	CreatePayment(payment *models.Payment) error
	CreateStateHistory(state *models.StateHistory) error
	UpdatePaymentState(payment *models.Payment) error
	GetPaymentByID(paymentID string) (models.Payment, error)
	GetStatusByID(orderID string) (string, error)
	GetIdempotencyKey(key string) (string, error)
	CreateIdempotencyKey(key, paymentID string) error
	GetHistoryByCustomerID(customerID string) ([]models.Payment, error)
}