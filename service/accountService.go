package service

import (
	"time"

	"github.com/rokafela/udemy-banking/domain"
	"github.com/rokafela/udemy-banking/dto"
	"github.com/rokafela/udemy-banking/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (das DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	transformedAccountRequest := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := das.repo.Save(transformedAccountRequest)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()

	return &response, nil
}

func NewAccountService(accountRepo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: accountRepo}
}
