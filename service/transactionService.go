package service

import "github.com/rokafela/udemy-banking/domain"

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
}

func (trxSrv DefaultTransactionService) UseEcho(str string) string {
	return trxSrv.trxRepo.Echo(str)
}
