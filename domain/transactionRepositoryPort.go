package domain

import "github.com/rokafela/udemy-banking/errs"

type TransactionRepository interface {
	Echo(string) string
	SelectAccountBalance(string) (float64, *errs.AppError)
	InsertTransactionAndUpdateBalance(*Transaction, *Account) (string, *errs.AppError)
}
