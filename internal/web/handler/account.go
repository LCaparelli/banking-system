package handler

import (
	"banking-system/internal/service"
	"banking-system/internal/web/request"
	"banking-system/internal/web/response"
	"encoding/json"
	"log"
	"net/http"
)

var (
	accountService = service.AccountServiceFactory()
)

func AccountGET(w http.ResponseWriter, body []byte) {
	var req request.AccountGET
	err, status, respBody := initReq(&req, body)
	if err != nil {
		w.WriteHeader(status)
		if _, err = w.Write(respBody); err != nil {
			log.Printf("AccountGET: Write: %v", err)
		}
		return
	}

	account, err := accountService.GetAccount(req.Id)
	if err != nil {
		// TODO: decide whether this should be 404. If we were getting the ID via URL instead of request body, it would.
		w.WriteHeader(http.StatusNotFound)
		log.Printf("AccountGET: GetAccount: %v", err)
		return
	}

	respBody, err = json.Marshal(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("AccountGET: Marshal: %v", err)
		return
	}

	_, err = w.Write([]byte(respBody))
	if err != nil {
		log.Printf("AccountGET: Write: %v", err)
	}
}

func AccountDELETE(w http.ResponseWriter, body []byte) {
	var req request.AccountDELETE
	err, status, respBody := initReq(&req, body)
	if err != nil {
		w.WriteHeader(status)
		if _, err = w.Write(respBody); err != nil {
			log.Printf("AccountDELETE: Write: %v", err)
		}
		return
	}

	err = accountService.DeleteAccount(req.Id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("AccountDELETE: DeleteAccount: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AccountPOST(w http.ResponseWriter, body []byte) {
	var req request.AccountPOST
	err, status, respBody := initReq(&req, body)
	if err != nil {
		w.WriteHeader(status)
		if _, err = w.Write(respBody); err != nil {
			log.Printf("AccountDELETE: Write: %v", err)
		}
		return
	}

	id, err := accountService.CreateAccount(req.Name, req.Address, req.Balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("AccountDELETE: CreateAccount: %v", err)
		return
	}

	respBody, err = json.Marshal(response.AccountGET{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("AccountDELETE: Marshal: %v", err)
		return
	}
	_, err = w.Write([]byte(respBody))
	if err != nil {
		log.Printf("AccountDELETE: Write: %v", err)
	}
}
