package domain

import "github.com/jmoiron/sqlx"

type TransactionRepositoryDb struct {
	dbPool *sqlx.DB
}

func SetTransactionRepositoryDb(activeDbPool *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{
		dbPool: activeDbPool,
	}
}

func (d TransactionRepositoryDb) Echo(str string) string {
	return str
}
