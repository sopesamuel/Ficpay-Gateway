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
		"PENDING": {"AUTHORIZED", "VOIDED"},
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
	res, err := state.bank.Authorize(bankreq, key)
	if err != nil{
		return models.Payment{},  fmt.Errorf("bank authorization failed: %w", err)
	}
	
	auth_id := res.Authorization_id
	return models.Payment{PaymentID: uuid.New().String(), OrderID: req.Order_id, CustomerID: req.Customer_id,Amount: req.Amount, Currency: "USD", Status: "AUTHORIZED", AuthID: &auth_id, CaptureID: nil, VoidID: nil, RefundID: nil,} , nil
}