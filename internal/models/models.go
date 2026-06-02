package models


//payment life cycle
type payment struct {
	system_id string
	order_id string
	customer_id string
	amount int

	authorization_id string
	authorized_creation string
	authorization_expiry string
	currency string
	status string
	
	capture_id string
	capture_at string

	void_id string
	voided_At string

	refund_id string
	refunded_at string

	idempotency_key string
}


//request from the ficmart
type request struct {

	order_id string
	customer_id string
	amount int
	system_id string

}


// request to the bank
type bankrequest struct {
	
}
