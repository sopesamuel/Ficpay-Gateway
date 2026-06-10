package models
import ("time")

//payment life cycle
type Payment struct {
	System_id string
	Order_id string
	Customer_id string
	Amount int
	Authorization_id string
	Authorized_creation time.Time
	Authorization_expiry time.Time
	Currency string
	Status string
	Capture_id string
	Capture_at time.Time
	Void_id string
	Refund_id string
}

type Card struct {
  	Card_number string 
  	Cvv string
  	Expiry_month int
  	Expiry_year int
}

// type Card struct {
//   	Card_number string `json:"card_number"`
//   	Cvv string `json:"cvv"`
//   	Expiry_month int `json:"expiry_month"`
//   	Expiry_year int `json:"expiry_year"`
// }
//request from the ficmart
type Martrequest struct {
	Order_id string
	Customer_id string
	Amount int
	System_id string
	Carddetails Card
}


// request to the bank
type Bankauthrequest struct {
	Amount int
	Carddetails Card
}

type Bankcapturerequest struct {
	Amount int
	Authorization_id string
}

type Bankvoidrequest struct {
	Authorization_id string
}

type Bankrefundrequest struct {
	Amount int
	Capture_id string
}

//response from the bank

type Bankauthresponse struct {
	Amount int `json:"amount"`
	Authorization_id string `json:"authorization_id"`
	Created_at time.Time `json:"created_at"`
	Currency string `json:"currency"`
	Expires_at time.Time `json:"expires_at"`
	Status string `json:"status"`
}

type Captureresponse struct {
	Amount int `json:"amount"`
	Authorization_id string `json:"authorization_id"`
	Capture_id string `json:"capture_id"`
	Captured_at time.Time `json:"captured_at"`
	Currency string `json:"currency"`
	Status string `json:"status"`
}

type Voidresponse struct {
	Authorization_id string `json:"authorization_id"`
	Status string `json:"status"`
	Void_id string `json:"void_id"`
	Voided_At time.Time `json:"voided_at"`
}

type Refundresponse struct {
	Amount int `json:"amount"`
	Capture_id string `json:"capture_id"`
	Currency string `json:"currency"`
	Refund_id string `json:"refund_id"`
	Refunded_at time.Time `json:"refunded_at"`
	Status string `json:"status"`
}

type Errorresponse struct {
	Error string `json:"error"`
	Message string `json:"message"`
}