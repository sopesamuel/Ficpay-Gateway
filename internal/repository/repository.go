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

//we basically inserting into the created table values that have filled our payment struct
// This triggers on authorize, first transaction.
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

func (r *Repository) CreateStateHistory(state *models.StateHistory) error {
	_, err := r.DB.Exec(
		`INSERT INTO state_history (payment_id, from_status, to_status)
		VALUES (?, ?, ?)`,
		state.PaymentID,
		state.FromStatus,
		state.ToStatus,
	)
	return err
}

func (r *Repository) UpdatePaymentState(payment *models.Payment) error {
	_, err := r.DB.Exec(
		`UPDATE payments SET status = ?, capture_id = ?, void_id = ?, refund_id = ? 
		WHERE payment_id = ?`,
		payment.Status,
		payment.CaptureID,
		payment.VoidID,
		payment.RefundID,
		payment.PaymentID,
	)
	return err
}



func (r *Repository) GetPaymentByID(paymentID string) (models.Payment, error){
	payment := models.Payment{}

	stmt := `SELECT payment_id, order_id, customer_id, amount, currency, status, auth_id, capture_id, void_id, refund_id, created_at, updated_at 
	FROM payments WHERE payment_id = ? `

	err := r.DB.QueryRow(stmt,paymentID).Scan(
		&payment.PaymentID,
 		&payment.OrderID,
 		&payment.CustomerID,
 		&payment.Amount,
 		&payment.Currency,
 		&payment.Status,
 		&payment.AuthID,
 		&payment.CaptureID,
 		&payment.VoidID,
 		&payment.RefundID,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)

	if err != nil {
		return models.Payment{}, err
	}

	return payment,nil
}

func (r *Repository) GetStatusByID(orderID string) (string, error) {
	var status string

	stmt := `SELECT status FROM payments WHERE order_id = ?`

	err := r.DB.QueryRow(stmt, orderID).Scan(&status)
	if err != nil {
		return "", err
	}

	return status, nil
}

	