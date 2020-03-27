package handler

import (
	"banking-system/internal/web/request"
	"banking-system/internal/web/response"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	notEnoughBalance = "Not enough balance to withdraw %.2f"
	success          = "Successfully withdrew %.2f"
)

func WithdrawPOST(w http.ResponseWriter, body []byte) {
	var req request.WithdrawPOST
	err, status, respBody := initReq(&req, body)
	if err != nil {
		w.WriteHeader(status)
		if _, err = w.Write(respBody); err != nil {
			log.Printf("WithdrawPOST: Write: %v", err)
		}
		return
	}

	if _, err = accountService.GetAccount(req.Id); err != nil {
		log.Printf("WithdrawPOST: GetAccount: %v", err)
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
		log.Printf("WithdrawPOST: Marshal: %v", err)
		return
	}

	_, err = w.Write([]byte(respBody))
	if err != nil {
		log.Printf("WithdrawPOST: Write: %v", err)
	}
}
