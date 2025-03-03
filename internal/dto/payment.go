package dto

type PaymentCreateRequest struct {
	Amount      uint   `json:"amount"`
	PayMethod   string `json:"pay_method"`
	ProductName string `json:"product_name"`
}

type PaymentBasicConfirmRequest struct {
	ImpUID string `json:"imp_uid"`
}
