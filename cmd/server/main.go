package main

import (
	"github.com/LCaparelli/banking-system/internal/web/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/account", handler.AccountHandler)
	http.HandleFunc("/withdraw", handler.WithdrawHandler)
	http.HandleFunc("/deposit", handler.DepositHandler)
	http.ListenAndServe(":8080", nil)
}
