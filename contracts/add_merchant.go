package contracts

type AddMerchantRequest struct {
	Area      string `json:"area,omitempty"`
	Name      string `json:"name"`
	Phone     string `json:"phone,omitempty"`
	AmountDue string `json:"amount_due,omitempty"`
}
