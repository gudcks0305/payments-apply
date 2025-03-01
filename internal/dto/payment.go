package dto

type PaymentCreateRequest struct {
	Amount      int    `json:"amount"`
	PayMethod   string `json:"pay_method"`
	ProductName string `json:"product_name"`
}

type PaymentData struct {
	ImpUID        string `json:"imp_uid"`
	MerchantUID   string `json:"merchant_uid"`
	PayMethod     string `json:"pay_method,omitempty"`
	Name          string `json:"name,omitempty"`
	PaidAmount    int    `json:"paid_amount,omitempty"`
	Currency      string `json:"currency,omitempty"`
	PgProvider    string `json:"pg_provider,omitempty"`
	PgType        string `json:"pg_type,omitempty"`
	PgTID         string `json:"pg_tid,omitempty"`
	ApplyNum      string `json:"apply_num,omitempty"`
	BuyerName     string `json:"buyer_name,omitempty"`
	BuyerEmail    string `json:"buyer_email,omitempty"`
	BuyerTel      string `json:"buyer_tel,omitempty"`
	BuyerAddr     string `json:"buyer_addr,omitempty"`
	BuyerPostcode string `json:"buyer_postcode,omitempty"`
	CustomData    string `json:"custom_data,omitempty"`
	Status        string `json:"status,omitempty"`
	PaidAt        int64  `json:"paid_at,omitempty"`
	ReceiptURL    string `json:"receipt_url,omitempty"`
	CardName      string `json:"card_name,omitempty"`
	BankName      string `json:"bank_name,omitempty"`
	CardQuota     int    `json:"card_quota,omitempty"`
	CardNumber    string `json:"card_number,omitempty"`
}
