package handler

import (
	"encoding/json"
	"fmt"
	"github.com/LCaparelli/banking-system/internal/web/request"
	"github.com/LCaparelli/banking-system/internal/web/response"
	"log"
	"net/http"
)

const (
	notEnoughBalance = "Not enough balance to withdraw %.2f"
	success          = "Successfully withdrew %.2f"
)

func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	body, err := reqBody(r)
	if err != nil {
		log.Printf("reqBody: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpMethod := r.Method
	switch httpMethod {
	case "POST":
		withdrawPOST(w, body)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		// TODO: decide whether should this should be logged and how.
	}
}

func withdrawPOST(w http.ResponseWriter, body []byte) {
	var req request.WithdrawPOST
	err, status, respBody := initReq(&req, body)
	if err != nil {
		w.WriteHeader(status)
		if _, err = w.Write(respBody); err != nil {
			log.Printf("withdrawPOST: Write: %v", err)
		}
		return
	}

	if _, err = accountService.GetAccount(req.Id); err != nil {
		log.Printf("withdrawPOST: GetAccount: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ok, msg := true, fmt.Sprintf(success, req.Amount)
	if err = accountService.Withdraw(req.Id, req.Amount); err != nil {
		ok = false
		msg = fmt.Sprintf(notEnoughBalance, req.Amount)
	}

	respBody, err = json.Marshal(response.WithdrawPOST{Ok: ok, Msg: msg})
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
