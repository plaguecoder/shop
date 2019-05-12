package contracts

import (
	"net/http"
)

type Customer struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Area        string `json:"area,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Description string `json:"description,omitempty"`
}

type Transaction struct {
	ID          int64  `json:"id,omitempty"`
	CustomerID  int64  `json:"customer_id,omitempty" db:"customer_id"`
	Amount      int64  `json:"amount,omitempty"`
	Date        string `json:"date,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

type Error struct {
	Title   string
	Message string
}

type AddCustomerResponse struct {
	StatusCode int    `json:"-"`
	Data       string `json:"data,omitempty"`
	Error      *Error `json:"error,omitempty"`
}

func (acr *AddCustomerResponse) Success() {
	acr.StatusCode = http.StatusOK
	acr.Data = "Successfully added given customer."
}

func (acr *AddCustomerResponse) BadRequest(title, message string) {
	acr.StatusCode = http.StatusBadRequest
	acr.Error = &Error{
		Title:   title,
		Message: message,
	}
}

func (acr *AddCustomerResponse) ServerError(message string) {
	acr.StatusCode = http.StatusInternalServerError
	acr.Error = &Error{
		Title:   "Internal Server Error",
		Message: message,
	}
}

type GetCustomerResponseData struct {
	*Customer
	Transactions []Transaction `json:"transactions"`
}

type GetCustomerResponse struct {
	StatusCode int                      `json:"-"`
	Data       *GetCustomerResponseData `json:"data,omitempty"`
	Error      *Error                   `json:"error,omitempty"`
}

func (gcr *GetCustomerResponse) Success(customer *Customer, transaction []Transaction) {
	gcr.StatusCode = http.StatusOK
	gcr.Data = &GetCustomerResponseData{Transactions: transaction, Customer: customer}
}

func (gcr *GetCustomerResponse) BadRequest(title, message string) {
	gcr.StatusCode = http.StatusBadRequest
	gcr.Error = &Error{
		Title:   title,
		Message: message,
	}
}

func (gcr *GetCustomerResponse) ServerError(message string) {
	gcr.StatusCode = http.StatusInternalServerError
	gcr.Error = &Error{
		Title:   "Internal Server Error",
		Message: message,
	}
}

type GetAllCustomersResponse struct {
	StatusCode int        `json:"-"`
	Data       []Customer `json:"data"`
	Error      *Error     `json:"error,omitempty"`
}

func (gacr *GetAllCustomersResponse) Success(customers []Customer) {
	gacr.StatusCode = http.StatusOK
	gacr.Data = customers
}

func (gacr *GetAllCustomersResponse) ServerError(message string) {
	gacr.StatusCode = http.StatusInternalServerError
	gacr.Data = []Customer{}
	gacr.Error = &Error{
		Title:   "Internal Server Error",
		Message: message,
	}
}

type AddTransactionResponse struct {
	StatusCode int
	Data       string
	Error      *Error
}

func (atr *AddTransactionResponse) Success() {
	atr.StatusCode = http.StatusOK
	atr.Data = "Successfully added transaction for given user."
}

func (atr *AddTransactionResponse) BadRequest(title, message string) {
	atr.StatusCode = http.StatusBadRequest
	atr.Error = &Error{
		Title:   title,
		Message: message,
	}
}

func (atr *AddTransactionResponse) ServerError(message string) {
	atr.StatusCode = http.StatusInternalServerError
	atr.Error = &Error{
		Title:   "Internal Server Error",
		Message: message,
	}
}
