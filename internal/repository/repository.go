package repository

import (
	"database/sql"
	"mart-gateway/internal/models"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreatePayment(payment *models.Payment) error {
	_, err := r.DB.Exec(
		`INSERT INTO payments (system_id, order_id, customer_id, amount, authorization_id, authorized_creation, authorization_expiry, currency, status, capture_id, capture_at, void_id, refund_id)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
		payment.PaymentID,
		payment.OrderID,
		payment.CustomerID,
		payment.Amount,
		payment.AuthID,
		payment.Currency,
		payment.Status,
		payment.CaptureID,
		payment.CreatedAt,
		payment.RefundID,
		payment.VoidID,
		payment.UpdatedAt,
	)
	return err
}
