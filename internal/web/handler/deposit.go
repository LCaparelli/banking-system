package handler

import (
	"encoding/json"
	"github.com/LCaparelli/banking-system/internal/web/request"
	"github.com/LCaparelli/banking-system/internal/web/response"
	"log"
	"net/http"
)

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	body, err := reqBody(r)
	if err != nil {
		log.Printf("reqBody: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpMethod := r.Method
	switch httpMethod {
	case "POST":
		depositPOST(w, body)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		// TODO: decide whether should this should be logged and how.
	}
}

func depositPOST(w http.ResponseWriter, body []byte) {
	var req request.DepositPOST
	err, status, respBody := initReq(&req, body)
	if err != nil {
		w.WriteHeader(status)
		if _, err = w.Write(respBody); err != nil {
			log.Printf("depositPOST: Write: %v", err)
		}
		return
	}

	if _, err = accountService.GetAccount(req.Id); err != nil {
		log.Printf("depositPOST: GetAccount: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	newBalance, err := accountService.Deposit(req.Id, req.Amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("withdrawPOST: Marshal: %v", err)
		return
	}

	respBody, err = json.Marshal(response.TransactionPOST{Ok: true, Msg: successDeposit, NewBalance: newBalance})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("withdrawPOST: Marshal: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(respBody))
	if err != nil {
		log.Printf("withdrawPOST: Write: %v", err)
	}
}
