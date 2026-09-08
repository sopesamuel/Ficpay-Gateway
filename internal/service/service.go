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
	//send the request to the bank and we get response 
	res, err := state.bank.Authorize(bankreq, key)
	if err != nil{
		return models.Payment{},  fmt.Errorf("bank authorization failed: %w", err)
	}
	Payment_ID := uuid.New().String()
	// we take details from the ficmart request and combine with the respomse from bank(auth id) and store in payments
	auth_id := res.Authorization_id
	// save to our payment db 
	auth_payment := models.Payment{PaymentID: Payment_ID, OrderID: req.Order_id, CustomerID: req.Customer_id,Amount: req.Amount, Currency: "USD", Status: "AUTHORIZED", AuthID: &auth_id, CaptureID: nil, VoidID: nil, RefundID: nil,}
	err = state.repository.CreatePayment(&auth_payment)
	if err != nil{
		return models.Payment{},  fmt.Errorf("failed to store authorization record to payment db: %w", err)
	}

	authorization_state := models.StateHistory{PaymentID: Payment_ID,FromStatus: nil, ToStatus: "AUTHORIZED" }
	err = state.repository.CreateStateHistory(&authorization_state)
	if err != nil{
		return models.Payment{},  fmt.Errorf("failed to store history authorization record to stae history db: %w", err)
	}

	return  auth_payment, nil
}