package app

import (
	"net/http"

	"github.com/gorilla/mux"
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
