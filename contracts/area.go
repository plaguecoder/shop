package contracts

import "net/http"

type Area struct {
	ID   int64  `json:"-"`
	Name string `json:"name,omitempty"`
}

type AddAreaResponse struct {
	StatusCode int    `json:"-"`
	Data       string `json:"data,omitempty"`
	Error      *Error `json:"error,omitempty"`
}

func (acr *AddAreaResponse) Success() {
	acr.StatusCode = http.StatusOK
	acr.Data = "Successfully added given customer."
}

func (acr *AddAreaResponse) BadRequest(title, message string) {
	acr.StatusCode = http.StatusBadRequest
	acr.Error = &Error{
		Title:   title,
		Message: message,
	}
}

func (acr *AddAreaResponse) ServerError(message string) {
	acr.StatusCode = http.StatusInternalServerError
	acr.Error = &Error{
		Title:   "Internal Server Error",
		Message: message,
	}
}

type GetAllAreasResponse struct {
	StatusCode int      `json:"-"`
	Data       []string `json:"data"`
	Error      *Error   `json:"error,omitempty"`
}

func (acr *GetAllAreasResponse) Success(names []string) {
	acr.StatusCode = http.StatusOK
	acr.Data = names
}

func (acr *GetAllAreasResponse) BadRequest(title, message string) {
	acr.StatusCode = http.StatusBadRequest
	acr.Error = &Error{
		Title:   title,
		Message: message,
	}
}

func (acr *GetAllAreasResponse) ServerError(message string) {
	acr.StatusCode = http.StatusInternalServerError
	acr.Error = &Error{
		Title:   "Internal Server Error",
		Message: message,
	}
}
