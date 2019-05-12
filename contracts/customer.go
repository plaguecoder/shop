package contracts

type Customer struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Area        string `json:"data,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Description string `json:"description,omitempty"`
}

type Error struct {
	Title   string
	Message string
}

type AddCustomerResponse struct {
	StatusCode int
	Data       string `json:"data,omitempty"`
	Error      *Error `json:"error,omitempty"`
}

type GetCustomersResponse struct {
	StatusCode int
	Data       []Customer `json:"data,omitempty"`
	Error      *Error     `json:"error,omitempty"`
}
