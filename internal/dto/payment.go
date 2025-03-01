package dto

type PaymentCreateRequest struct {
	Amount      int    `json:"amount"`
	PayMethod   string `json:"pay_method"`
	ProductName string `json:"product_name"`
}
