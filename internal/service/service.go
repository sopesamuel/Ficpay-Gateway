package service

import (
	"fmt"
	"mart-gateway/internal/models"
	"slices"

	"github.com/google/uuid"
)

type MainService struct {
	bank BankInterface
	repository RepositoryInterface
}

func ActivateService(bankService BankInterface, repo RepositoryInterface) *MainService {
	return &MainService{bank: bankService, repository: repo}
}

//state machine for communication with the db, checks if 

func StateMachine(from, to string) bool {
	state := map[string][]string{
		"PENDING": {"AUTHORIZED"},
		"AUTHORIZED": {"CAPTURED", "VOIDED"},
		"CAPTURED": {"REFUNDED"},
		
	}

	if slices.Contains(state[from], to){
		return true
	}
	return false
}

//Authorize to bank from ficmart

func (state *MainService) AuthorizePayment(req models.Martrequest, bankreq models.Bankauthrequest, key string) (models.Payment,error){
	//send the request to the bank TO INITIATE PENDING AND AUTHORIZATION states
	
	auth_payment := models.Payment{
		PaymentID: uuid.New().String(), 
		OrderID: req.Order_id, 
		CustomerID: req.Customer_id,
		Amount: req.Amount, 
		Currency: "USD", 
		Status: "PENDING", 
		}

	err := state.repository.CreatePayment(&auth_payment)
	if err != nil{
		return models.Payment{},  fmt.Errorf("starting payment failed (pending state): %w", err)
	}

	//THEN RECORD THE PENDING STATE
	err = state.repository.CreateStateHistory(&models.StateHistory{ PaymentID: auth_payment.PaymentID, ToStatus: "PENDING",})
	if err != nil{
		return models.Payment{},  fmt.Errorf("failed to store pending state history to stae history db: %w", err)
	}

	//Then we initiate authorization request to bank
	res, err := state.bank.Authorize(bankreq, key)
	if err != nil{
		return models.Payment{},  fmt.Errorf("bank authorization failed: %w", err)
	}

	auth_id_rep := res.Authorization_id
	pending_state := "PENDING"
	auth_payment.AuthID = &auth_id_rep
	auth_payment.Status = "AUTHORIZED"

	err = state.repository.UpdatePaymentState(&auth_payment)
	if err != nil{
		return models.Payment{},  fmt.Errorf("failed to store authorization record to payment db: %w", err)
	}
	
	err = state.repository.CreateStateHistory(&models.StateHistory{PaymentID: auth_payment.PaymentID, FromStatus: &pending_state, ToStatus: "AUTHORIZED"})
	if err != nil{
		return models.Payment{},  fmt.Errorf("failed to store history OF authorization record to state history db: %w", err)
	}

	return  auth_payment, nil
}

func (state *MainService) CapturePayment(payment_id string, bankreq models.Bankcapturerequest, key string) (models.Payment, error){
	res_payment ,err := state.repository.GetPaymentByID(payment_id)
	if err != nil{
		return models.Payment{},  fmt.Errorf("failed to get payment reference from db: %w", err)
	}

	if !StateMachine(res_payment.Status, "CAPTURED"){
		return models.Payment{},  fmt.Errorf("Invalid payment transition from: %s to %s",res_payment.Status, "CAPTURED")
	}

	cap_res, err := state.bank.Capture(bankreq, key)
	if err != nil {
		return models.Payment{},  fmt.Errorf("failed to capture payment from bank: %w", err)
	}

	prev_status := res_payment.Status
	res_payment.CaptureID = &cap_res.Capture_id
	res_payment.Status = "CAPTURED"

	err = state.repository.UpdatePaymentState(&res_payment)
	if err != nil{
		return models.Payment{},  fmt.Errorf("failed to store authorization record to payment db: %w", err)
	}

	//THEN RECORD THE CAPTURED STATE
	err = state.repository.CreateStateHistory(&models.StateHistory{ PaymentID: res_payment.PaymentID, FromStatus: &prev_status, ToStatus: "CAPTURED",})
	if err != nil{
		return models.Payment{},  fmt.Errorf("failed to store captured state history to state history db: %w", err)
	}

	return res_payment, nil
}

func (state *MainService) VoidPayment(payment_id string, bankreq models.Bankvoidrequest, key string) (models.Payment, error){
	res_payment ,err := state.repository.GetPaymentByID(payment_id)
	if err != nil{
		return models.Payment{},  fmt.Errorf("failed to get payment reference from db for void transaction: %w", err)
	}

	if !StateMachine(res_payment.Status, "VOIDED"){
		return models.Payment{},  fmt.Errorf("Invalid payment transition from: %s to %s",res_payment.Status, "VOIDED")
	}

	void_res, err := state.bank.Void(bankreq, key)
	if err != nil{
		return models.Payment{},  fmt.Errorf("Bank unable to void transaction: %w", err)
	}

	prev_status := res_payment.Status
	res_payment.Status = "VOIDED"
	res_payment.VoidID = &void_res.Void_id

	err = state.repository.UpdatePaymentState(&res_payment)
	if err != nil{
		return models.Payment{},  fmt.Errorf("Unable to update voided transaction to payment db: %w", err)
	}

	err = state.repository.CreateStateHistory(&models.StateHistory{PaymentID: res_payment.PaymentID, FromStatus: &prev_status, ToStatus: "VOIDED",})
	if err != nil{
		return models.Payment{},  fmt.Errorf("Unable to store voided transaction to state history db: %w", err)
	}

	return res_payment, nil
}

func (state *MainService) RefundPayment(payment_id string, refundreq models.Bankrefundrequest, key string,) (models.Payment, error){

	res_payment ,err := state.repository.GetPaymentByID(payment_id)
	if err != nil{
		return models.Payment{},  fmt.Errorf("failed to get payment reference from db for refund transaction: %w", err)
	}

	if !StateMachine(res_payment.Status, "REFUNDED"){
		return models.Payment{},  fmt.Errorf("Invalid payment transition from: %s to %s",res_payment.Status, "REFUNDED")
	}

	void_res, err := state.bank.Refund(refundreq, key)
	if err != nil{
		return models.Payment{},  fmt.Errorf("Bank unable to refund transaction: %w", err)
	}

	prev_status := res_payment.Status
	res_payment.Status = "REFUNDED"
	res_payment.RefundID = &void_res.Refund_id

	err = state.repository.UpdatePaymentState(&res_payment)
	if err != nil{
		return models.Payment{},  fmt.Errorf("Unable to update refund transaction to payment db: %w", err)
	}

	err = state.repository.CreateStateHistory(&models.StateHistory{PaymentID: res_payment.PaymentID, FromStatus: &prev_status, ToStatus: "REFUNDED",})
	if err != nil{
		return models.Payment{},  fmt.Errorf("Unable to store refund transaction to state history db: %w", err)
	}

	return res_payment, nil
}

func (state *MainService) GetPayment(payment_id string) (models.Payment, error){
	res_payment, err := state.repository.GetPaymentByID(payment_id)
	if err != nil{
		return models.Payment{},  fmt.Errorf("failed to get payment reference from db for id: %w", err)
	}

	return res_payment, nil
}