package contracts

import "net/http"

type Transaction struct {
	ID          int64  `json:"id,omitempty"`
	CustomerID  int64  `json:"customer_id,omitempty" db:"customer_id"`
	Amount      int64  `json:"amount,omitempty"`
	Date        string `json:"date,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
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
