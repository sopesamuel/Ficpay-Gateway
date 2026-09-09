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

func (state MainService) AuthorizePayment(req models.Martrequest, bankreq models.Bankauthrequest, key string) (models.Payment,error){
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
		return models.Payment{},  fmt.Errorf("failed to store history OF authorization record to stae history db: %w", err)
	}

	return  auth_payment, nil
}