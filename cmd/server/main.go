package main

import (
	"banking-system/internal/web/handler"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/account", accountHandler)
	http.HandleFunc("/withdraw", withdrawHandler)
	http.HandleFunc("/deposit", depositHandler)
	http.ListenAndServe(":8080", nil)
}

func reqBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("readAll: %v", err)
	}
	return body, nil
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	body, err := reqBody(r)
	if err != nil {
		log.Printf("reqBody: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpMethod := r.Method
	switch httpMethod {
	case "GET":
		handler.AccountGET(w, body)
	case "DELETE":
		handler.AccountDELETE(w, body)
	case "POST":
		handler.AccountPOST(w, body)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		// TODO: decide whether should this should be logged and how.
	}
}

func withdrawHandler(w http.ResponseWriter, r *http.Request) {
	body, err := reqBody(r)
	if err != nil {
		log.Printf("reqBody: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpMethod := r.Method
	switch httpMethod {
	case "POST":
		handler.WithdrawPOST(w, body)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		// TODO: decide whether should this should be logged and how.
	}
}

func depositHandler(w http.ResponseWriter, r *http.Request) {
	body, err := reqBody(r)
	if err != nil {
		log.Printf("reqBody: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpMethod := r.Method
	switch httpMethod {
	case "POST":
		handler.DepositPOST(w, body)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		// TODO: decide whether should this should be logged and how.
	}
}
