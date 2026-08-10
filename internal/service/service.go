package service

type MainService struct {
	bank BankInterface
	repository RepositoryInterface
}

func ActivateService(bankService BankInterface, repo RepositoryInterface) *MainService {
	return &MainService{bank: bankService, repository: repo}
}

//state machine