package repository

import (
	"database/sql"
	"mart-gateway/internal/models"
)

type Repository struct {
	DB *sql.DB
}

func NewPaymentRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreatePayment(payment *models.Payment) error {
	_, err := r.DB.Exec(
		`INSERT INTO payments (payment_id, order_id, customer_id, amount, currency, status, auth_id, capture_id, void_id, refund_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		payment.PaymentID,
		payment.OrderID,
		payment.CustomerID,
		payment.Amount,
		payment.Currency,
		payment.Status,
		payment.AuthID,
		payment.CaptureID,
		payment.VoidID,
		payment.RefundID,
	)
	return err
}
