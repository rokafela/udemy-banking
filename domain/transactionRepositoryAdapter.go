package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/rokafela/udemy-banking/errs"
	"github.com/rokafela/udemy-banking/logger"
)

type TransactionRepositoryDb struct {
	dbPool *sqlx.DB
}

func SetTransactionRepositoryDb(activeDbPool *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{
		dbPool: activeDbPool,
	}
}

func (db TransactionRepositoryDb) Echo(str string) string {
	return str
}

func (db TransactionRepositoryDb) SelectAccountBalance(account_id string) (float64, *errs.AppError) {
	var account_balance float64

	selectSql := "SELECT amount FROM accounts WHERE status = 1 AND account_id = ?"

	err := db.dbPool.Get(&account_balance, selectSql, account_id)
	if err != nil {
		return 0, errs.NewNotFoundError("account not found")
	}

	return account_balance, nil
}

func (db TransactionRepositoryDb) InsertTransactionAndUpdateBalance(trx_data *Transaction, account_data *Account) (string, *errs.AppError) {
	tx := db.dbPool.MustBegin()

	sqlUpdate := "UPDATE accounts SET amount = ? WHERE account_id = ?;"
	tx.MustExec(sqlUpdate, account_data.Amount, account_data.AccountId)

	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?);"
	result := tx.MustExec(sqlInsert, trx_data.AccountId, trx_data.Amount, trx_data.TransactionType, trx_data.TransactionDate)

	err := tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return "0", errs.NewUnexpectedError("Unexpected error from database")
	}

	trx_id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id: " + err.Error())
		return "0", errs.NewUnexpectedError("Unexpected error from database")
	}

	return strconv.FormatInt(trx_id, 10), nil
}
