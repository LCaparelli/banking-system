package handler

import (
	"encoding/json"
	"fmt"
	"github.com/LCaparelli/banking-system/internal/web/request"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	invalidJSON      = "invalid JSON syntax or fields in request body"
	notEnoughBalance = "Not enough balance to withdraw %.2f"
	successWithdraw  = "Successfully withdrew %.2f"
	successDeposit   = "Deposit successful"
)

func unmarshalReq(body []byte, req request.Request) error {
	err := json.Unmarshal(body, req)
	if err != nil {
		return fmt.Errorf("%s Unmarshal: %v", invalidJSON, err)
	}
	return nil
}

func initReq(req request.Request, body []byte) (error, int, []byte) {
	if err := unmarshalReq(body, req); err != nil {
		err = fmt.Errorf("request: %v", err)
		log.Println(err)
		return err, http.StatusBadRequest, []byte(invalidJSON)
	}

	if err := req.Validate(); err != nil {
		err = fmt.Errorf("validate: %v", err)
		log.Println(err)
		return err, http.StatusBadRequest, []byte(err.Error())
	}
	return nil, 0, nil
}

func reqBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("readAll: %v", err)
	}
	return body, nil
}
