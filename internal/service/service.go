package service

import "slices"

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