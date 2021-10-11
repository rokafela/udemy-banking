package dto

import (
	"strings"

	"github.com/rokafela/udemy-banking/errs"
)

type NewTransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

type NewTransactionResponse struct {
	TransactionId  string  `json:"transaction_id"`
	UpdatedBalance float64 `json:"updated_balance"`
}

func (req NewTransactionRequest) Validate() *errs.AppError {
	if req.AccountId == "" {
		return errs.NewValidationError("Required parameter: account_id")
	}

	if req.Amount < 1 {
		return errs.NewValidationError("Required parameter: amount. Minimum amount: 1")
	}

	if req.TransactionType == "" {
		return errs.NewValidationError("Required parameter: transaction_type")
	}
	lowered_trx_type := strings.ToLower(req.TransactionType)
	if lowered_trx_type != "withdrawal" && lowered_trx_type != "deposit" {
		return errs.NewValidationError("Allowed transaction_type: deposit, withdrawal")
	}

	return nil
}
