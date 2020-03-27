package handler

import (
	"encoding/json"
	"github.com/LCaparelli/banking-system/internal/service"
	"github.com/LCaparelli/banking-system/internal/web/request"
	"github.com/LCaparelli/banking-system/internal/web/response"
	"log"
	"net/http"
)

var (
	accountService = service.AccountServiceFactory()
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	body, err := reqBody(r)
	if err != nil {
		log.Printf("reqBody: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpMethod := r.Method
	switch httpMethod {
	case "GET":
		accountGET(w, body)
	case "DELETE":
		accountDELETE(w, body)
	case "POST":
		accountPOST(w, body)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		// TODO: decide whether should this should be logged and how.
	}
}

func accountGET(w http.ResponseWriter, body []byte) {
	var req request.AccountGET
	err, status, respBody := initReq(&req, body)
	if err != nil {
		w.WriteHeader(status)
		if _, err = w.Write(respBody); err != nil {
			log.Printf("accountGET: Write: %v", err)
		}
		return
	}

	account, err := accountService.GetAccount(req.Id)
	if err != nil {
		// TODO: decide whether this should be 404. If we were getting the ID via URL instead of request body, it would.
		w.WriteHeader(http.StatusNotFound)
		log.Printf("accountGET: GetAccount: %v", err)
		return
	}

	respBody, err = json.Marshal(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("accountGET: Marshal: %v", err)
		return
	}
}

func accountDELETE(w http.ResponseWriter, body []byte) {
	var req request.AccountDELETE
	err, status, respBody := initReq(&req, body)
	if err != nil {
		w.WriteHeader(status)
		if _, err = w.Write(respBody); err != nil {
			log.Printf("accountDELETE: Write: %v", err)
		}
		return
	}

	err = accountService.DeleteAccount(req.Id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("accountDELETE: DeleteAccount: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func accountPOST(w http.ResponseWriter, body []byte) {
	var req request.AccountPOST
	err, status, respBody := initReq(&req, body)
	if err != nil {
		w.WriteHeader(status)
		if _, err = w.Write(respBody); err != nil {
			log.Printf("accountDELETE: Write: %v", err)
		}
		return
	}

	id, err := accountService.CreateAccount(req.Name, req.Address, req.Balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("accountDELETE: CreateAccount: %v", err)
		return
	}

	respBody, err = json.Marshal(response.AccountGET{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("accountDELETE: Marshal: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(respBody))
	if err != nil {
		log.Printf("accountDELETE: Write: %v", err)
	}
}
