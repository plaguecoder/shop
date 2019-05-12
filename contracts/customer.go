package contracts

type Customer struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Area        string `json:"area,omitempty"`
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

type GetCustomerResponse struct {
	StatusCode int
	Data       *Customer `json:"data,omitempty"`
	Error      *Error    `json:"error,omitempty"`
}

type GetAllCustomersResponse struct {
	StatusCode int
	Data       []Customer `json:"data"`
	Error      *Error     `json:"error,omitempty"`
}
