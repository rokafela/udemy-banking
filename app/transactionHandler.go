package app

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rokafela/udemy-banking/dto"
	"github.com/rokafela/udemy-banking/errs"
	"github.com/rokafela/udemy-banking/service"
)

type TransactionHandler struct {
	trxSrv service.TransactionService
}

func (trxHdl TransactionHandler) CallUseEcho(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str := vars["tmp"]
	response := trxHdl.trxSrv.UseEcho(str)
	writeResponse(w, http.StatusOK, response)
}

func (trxHdl TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var received_body dto.NewTransactionRequest
	var response_body dto.NewTransactionResponse

	json.NewDecoder(r.Body).Decode(&received_body)
	err := received_body.Validate()
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.AsMessage())
		return
	}

	// get account balance
	account_balance, err := trxHdl.trxSrv.GetAccountBalance(received_body.AccountId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	var updated_balance float64

	// check transaction type
	lowered_trx_type := strings.ToLower(received_body.TransactionType)
	if lowered_trx_type == "withdrawal" {
		// check balance
		if received_body.Amount > account_balance {
			err := errs.NewValidationError("Insufficient balance")
			writeResponse(w, err.Code, err.AsMessage())
			return
		}
		// process withdrawal
		updated_balance = account_balance - received_body.Amount
	} else {
		// process deposit
		updated_balance = account_balance + received_body.Amount
	}

	// save transaction
	trx_id, err := trxHdl.trxSrv.SaveTransactionAndBalance(&received_body, updated_balance)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	response_body.TransactionId = trx_id
	response_body.UpdatedBalance = updated_balance
	writeResponse(w, http.StatusOK, response_body)

}
