package handler

import (
	"github.com/LCaparelli/banking-system/internal/web/request"
	"log"
	"net/http"
)

func DepositPOST(w http.ResponseWriter, body []byte) {
	var req request.DepositPOST
	err, status, respBody := initReq(&req, body)
	if err != nil {
		w.WriteHeader(status)
		if _, err = w.Write(respBody); err != nil {
			log.Printf("DepositPOST: Write: %v", err)
		}
		return
	}

	if _, err = accountService.GetAccount(req.Id); err != nil {
		log.Printf("DepositPOST: GetAccount: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err = accountService.Deposit(req.Id, req.Amount); err != nil {
		log.Printf("DepositPOST: Deposit: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
