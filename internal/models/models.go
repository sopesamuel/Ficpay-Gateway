package models
import ("time")

//payment life cycle
type payment struct {
	system_id string
	order_id string
	customer_id string
	amount int

	authorization_id string
	authorized_creation time.Time
	authorization_expiry time.Time
	currency string
	status string
	
	capture_id string
	capture_at time.Time

	void_id string
	voided_At time.Time

	refund_id string
	refunded_at time.Time

}

type card struct {
  	card_number string
  	cvv string
  	expiry_month int
  	expiry_year int
}
//request from the ficmart
type martrequest struct {
	order_id string
	customer_id string
	amount int
	system_id string
	carddetails card
}


// request to the bank
type bankauthrequest struct {
	amount int
	carddetails card
}

type bankcapturerequest struct {
	amount int
	authorization_id string
}

type bankvoidrequest struct {
	authorization_id string
}

type bankrefundrequest struct {
	amount int
	capturen_id string
}
