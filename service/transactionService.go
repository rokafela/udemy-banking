package service

import (
	"time"

	"github.com/rokafela/udemy-banking/domain"
	"github.com/rokafela/udemy-banking/dto"
	"github.com/rokafela/udemy-banking/errs"
)

type DefaultTransactionService struct {
	trxRepo domain.TransactionRepository
}

func SetTransactionService(activeRepo domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{
		trxRepo: activeRepo,
	}
}

type TransactionService interface {
	UseEcho(string) string
	GetAccountBalance(string) (float64, *errs.AppError)
	SaveTransactionAndBalance(*dto.NewTransactionRequest, float64) (string, *errs.AppError)
}

func (trxSrv DefaultTransactionService) UseEcho(str string) string {
	return trxSrv.trxRepo.Echo(str)
}

func (trxSrv DefaultTransactionService) GetAccountBalance(account_id string) (float64, *errs.AppError) {
	account_balance, err := trxSrv.trxRepo.SelectAccountBalance(account_id)
	if err != nil {
		return 0, err
	}
	return account_balance, nil
}

func (trxSrv DefaultTransactionService) SaveTransactionAndBalance(received_body *dto.NewTransactionRequest, updated_balance float64) (string, *errs.AppError) {
	trx_data := domain.Transaction{
		AccountId:       received_body.AccountId,
		Amount:          received_body.Amount,
		TransactionType: received_body.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	account_data := domain.Account{
		AccountId: received_body.AccountId,
		Amount:    updated_balance,
	}

	trx_id, err := trxSrv.trxRepo.InsertTransactionAndUpdateBalance(&trx_data, &account_data)
	if err != nil {
		return "0", err
	}

	return trx_id, nil
}
