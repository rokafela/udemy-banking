package domain

type TransactionRepository interface {
	Echo(string) string
}
