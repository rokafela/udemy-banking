package dto

type NewTransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

type NewTransactionResponse struct {
	TransactionId  string  `json:"transaction_id"`
	UpdatedBalance float64 `json:"updated_balance"`
}
