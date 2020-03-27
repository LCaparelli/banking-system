- [banking-system](#banking-system)
  * [REST API](#rest-api)
    + [/account](#-account)
      - [GET](#get)
      - [DELETE](#delete)
      - [POST](#post)
    + [/deposit](#-deposit)
    + [/withdraw](#-withdraw)
    + [REST API server installation](#rest-api-server-installation)
  * [Library API definition](#library-api-definition)
    + [Library import](#library-import)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>

# banking-system
[![Go Report Card](https://goreportcard.com/badge/github.com/LCaparelli/banking-system)](https://goreportcard.com/report/github.com/LCaparelli/banking-system)

banking-system is my [Go](https://golang.org/) playground project. It's meant to be a project where I can practice concepts I learned from the book "[The Go Programming Language](https://www.gopl.io/)". It's a *very* simple banking system with no intention of production use.

## REST API

The server REST API ran by `cmd/server/main.go` has the following endpoints:

  - /account
  - /deposit
  - /withdraw

Alternatively you may also use the account service library for Go present in `pkg/account/`. For more details about it refer to [its API definition](#library-api-definition).

###  /account
All requests to this endpoint must be performed with the "Content-type" header set to "application/json".

#### GET
To gather information about an existing account issue an HTTP GET.  The body itself is a JSON with a single integer field "id".

For example:
```json
{
    "id": 0
}
```

The "id" field must not be negative.

If the account with the informed id exists the response status will be OK (200) and the response body will contain the customer's name, address, current balance and the account id.

For example:

```json
{
    "Name": "Lucas Caparelli",
    "Address": "test, 321",
    "Balance": 0,
    "Id": 0
}
```

If the account does not exist the server responds with NotFound status (404) and the response body is empty.

#### DELETE
To delete an existing account issue an HTTP GET.  The body itself is a JSON with a single integer field "id".

For example:
```json
{
    "id": 0
}
```

The "id" field must not be negative.

If the account with the informed id exists the response status will be OK (200) and the body will be empty.

If the account does not exist the server responds with NotFound status (404) and the response body is empty.

#### POST
To create an account issue an HTTP POST. The body itself is a JSON with two strings ("name" and "address") and a decimal ("balance") field.

For example:

```json
{
     "name": "Lucas Caparelli",
     "address": "test, 321",
     "balance": 0
}
```
The "name" and "address" fields must not be empty. The "balance" field must not be negative.

This always return OK (200) with an empty body.

### /deposit
To make a deposit issue an HTTP POST. The  body itself is a JSON with an integer ("id") and a decimal ("amount").

For example:

```json
{
     "id": 0,
     "amount": 100.5
}
```
The "id" field must not be negative. The "balance" field must be positive.

If the account with the informed id exists the response status will be OK (200) and the body will be empty.

If the account does not exist the server responds with NotFound status (404) and the response body is empty.

### /withdraw
To make a withdraw issue an HTTP POST. The  body itself is a JSON with an integer ("id") and a decimal ("amount").

For example:

```json
{
     "id": 0,
     "amount": 100.5
}
```

The "id" field must not be negative. The "balance" field must be positive.

If the account with the informed id exists the response status will be OK (200) and the body will contain information regarding the operation's success. The body contains a boolean indicating whether the operation was successful or not ("ok") and a string describing the result ("msg").

For example, if there is enough balance to complete the withdraw:

```json
{
    "Ok": true,
    "Msg": "Successfully withdrew 100.50"
}
```

If there isn't enough balance:

```json
{
    "Ok": false,
    "Msg": "Not enough balance to withdraw 100.50"
}
```
If the account does not exist the server responds with NotFound status (404) and the response body is empty.

### REST API server installation

To install the HTTP server run:

```bash
$ go get github.com/LCaparelli/banking-system/cmd/server
```

## Library API definition

This project also contains a library so that the account service can be used by other Go programs. It is exported by the `github.com/LCaparelli/banking-system/pkg/account` package.

The following functions are part of the API:
- `Account(id int) (*domain.Account, error)`

  The "id" argument must not be negative.

- `DeleteAccount(id int) error`

  The "id" argument must not be negative.

- `CreateAccount(name, address string, balance float64) (int, error)`

 The "name" and "address" arguments must not be empty. The "balance" argument must not be negative.

- `Deposit(id int, amount float64) error`

  The "id" argument must not be negative. The "balance" argument must be positive.

- `Withdraw(id int, amount float64) error`

  The "id" argument must not be negative. The "balance" argument must be positive.

### Library import

To import the library in your program use the following import path:

```go
import "github.com/LCaparelli/banking-system/pkg/account"
```
