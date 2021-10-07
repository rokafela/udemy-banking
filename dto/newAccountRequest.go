package dto

import (
	"strings"

	"github.com/rokafela/udemy-banking/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (req NewAccountRequest) Validate() *errs.AppError {
	if req.Amount < 5000 {
		return errs.NewValidationError("Minimum amount for new account: 5000")
	}
	loweredAccountType := strings.ToLower(req.AccountType)
	if loweredAccountType != "saving" && loweredAccountType != "checking" {
		return errs.NewValidationError("Allowed account type: checking, saving")
	}
	return nil
}
