package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rokafela/udemy-banking/dto"
	"github.com/rokafela/udemy-banking/service"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ah AccountHandlers) CreateNewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		vars := mux.Vars(r)
		customer_id := vars["customer_id"]
		request.CustomerId = customer_id
		account, appError := ah.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
