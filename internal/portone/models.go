package portone

import "time"

type TokenRequest struct {
	ImpKey    string `json:"imp_key"`
	ImpSecret string `json:"imp_secret"`
}

type TokenResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response struct {
		AccessToken string    `json:"access_token"`
		ExpiredAt   int64     `json:"expired_at"`
		Now         int64     `json:"now"`
		ExpireTime  time.Time `json:"-"`
	} `json:"response"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type APIResponse struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Response interface{} `json:"response"`
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

/*
{
  // "imp_uid": "imp_uid",
  // "merchant_uid": "merchant_uid",
  // "amount": 0,
  // "tax_free": 0,
  // "vat_amount": 0,
  // "checksum": 0,
  // "reason": "reason",
  // "refund_holder": "refund_holder",
  // "refund_bank": "refund_bank",
  // "refund_account": "refund_account",
  // "refund_tel": "refund_tel",
  // "retain_promotion": false,
  // "extra": [],
}
*/

type PaymentCancelRequest struct {
	ImpUID        string `json:"imp_uid"`
	MerchantUID   string `json:"merchant_uid"`
	Amount        int    `json:"amount"`
	TaxFree       int    `json:"tax_free"`
	VatAmount     int    `json:"vat_amount"`
	Checksum      int    `json:"checksum"`
	Reason        string `json:"reason"`
	RefundHolder  string `json:"refund_holder"`
	RefundBank    string `json:"refund_bank"`
	RefundAccount string `json:"refund_account"`
	RefundTel     string `json:"refund_tel"`
	RetainPromo   bool   `json:"retain_promotion"`
	Extra         []struct {
	} `json:"extra"`
}
