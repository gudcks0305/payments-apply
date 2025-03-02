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

type APIResponse[T any] struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response T      `json:"response"`
}

type PaymentClientResponse struct {
	Success       bool    `json:"success"`
	ErrorCode     *string `json:"error_code"`
	ErrorMsg      *string `json:"error_msg"`
	ImpUid        string  `json:"imp_uid"`
	MerchantUid   string  `json:"merchant_uid"`
	PayMethod     string  `json:"pay_method"`
	PaidAmount    float64 `json:"paid_amount"`
	Status        string  `json:"status"`
	Name          string  `json:"name"`
	PgProvider    string  `json:"pg_provider"`
	EmbPgProvider *string `json:"emb_pg_provider"`
	PgTid         string  `json:"pg_tid"`
	BuyerName     string  `json:"buyer_name"`
	BuyerEmail    string  `json:"buyer_email"`
	BuyerTel      string  `json:"buyer_tel"`
	BuyerAddr     string  `json:"buyer_addr"`
	BuyerPostcode string  `json:"buyer_postcode"`
	CustomData    *string `json:"custom_data"`
	PaidAt        int     `json:"paid_at"` // Unix timestamp (int 타입으로 변경)
	ReceiptUrl    *string `json:"receipt_url"`
	ApplyNum      *string `json:"apply_num"`
	VbankNum      *string `json:"vbank_num"`
	VbankName     *string `json:"vbank_name"`
	VbankHolder   *string `json:"vbank_holder"`
	VbankDate     int     `json:"vbank_date"` // Unix timestamp (int 타입으로 변경)
}

type PaymentData struct {
	ImpUID            string                    `json:"imp_uid"`
	MerchantUID       string                    `json:"merchant_uid"`
	PayMethod         *string                   `json:"pay_method,omitempty"`
	Channel           *string                   `json:"channel,omitempty"`
	PgProvider        *string                   `json:"pg_provider,omitempty"`
	EmbPgProvider     *string                   `json:"emb_pg_provider,omitempty"`
	PgTid             *string                   `json:"pg_tid,omitempty"`
	PgID              *string                   `json:"pg_id,omitempty"`
	Escrow            *bool                     `json:"escrow,omitempty"`
	ApplyNum          *string                   `json:"apply_num,omitempty"`
	BankCode          *string                   `json:"bank_code,omitempty"`
	BankName          *string                   `json:"bank_name,omitempty"`
	CardCode          *string                   `json:"card_code,omitempty"`
	CardName          *string                   `json:"card_name,omitempty"`
	CardIssuerCode    *string                   `json:"card_issuer_code,omitempty"`
	CardIssuerName    *string                   `json:"card_issuer_name,omitempty"`
	CardPublisherCode *string                   `json:"card_publisher_code,omitempty"`
	CardPublisherName *string                   `json:"card_publisher_name,omitempty"`
	CardQuota         *int                      `json:"card_quota,omitempty"`
	CardNumber        *string                   `json:"card_number,omitempty"` // 결제에 사용된 마스킹된 카드번호
	CardType          *int                      `json:"card_type,omitempty"`   // 결제건에 사용된 카드 구분코드
	VbankCode         *string                   `json:"vbank_code,omitempty"`
	VbankName         *string                   `json:"vbank_name,omitempty"`
	VbankNum          *string                   `json:"vbank_num,omitempty"`
	VbankHolder       *string                   `json:"vbank_holder,omitempty"`
	VbankDate         *int                      `json:"vbank_date,omitempty"`
	VbankIssuedAt     *int                      `json:"vbank_issued_at,omitempty"`
	Name              *string                   `json:"name,omitempty"`
	Amount            float64                   `json:"amount"`
	CancelAmount      float64                   `json:"cancel_amount"`
	Currency          string                    `json:"currency"`
	BuyerName         *string                   `json:"buyer_name,omitempty"`
	BuyerEmail        *string                   `json:"buyer_email,omitempty"`
	BuyerTel          *string                   `json:"buyer_tel,omitempty"`
	BuyerAddr         *string                   `json:"buyer_addr,omitempty"`
	BuyerPostcode     *string                   `json:"buyer_postcode,omitempty"`
	CustomData        *string                   `json:"custom_data,omitempty"`
	UserAgent         *string                   `json:"user_agent,omitempty"`
	Status            string                    `json:"status"`
	StartedAt         *int                      `json:"started_at,omitempty"`
	PaidAt            *int                      `json:"paid_at,omitempty"`
	FailedAt          *int                      `json:"failed_at,omitempty"`
	CancelledAt       *int                      `json:"cancelled_at,omitempty"`
	FailReason        *string                   `json:"fail_reason,omitempty"`
	CancelReason      *string                   `json:"cancel_reason,omitempty"`
	ReceiptURL        *string                   `json:"receipt_url,omitempty"`
	CancelHistory     []PaymentCancelAnnotation `json:"cancel_history,omitempty"`
	CancelReceiptURLs []string                  `json:"cancel_receipt_urls,omitempty"`
	CashReceiptIssued *bool                     `json:"cash_receipt_issued,omitempty"`
	CustomerUID       *string                   `json:"customer_uid,omitempty"`
	CustomerUIDUsage  *string                   `json:"customer_uid_usage,omitempty"`
	Promotion         *interface{}              `json:"promotion,omitempty"`
}
type PaymentCancelAnnotation struct {
	PgTid          string  `json:"pg_tid"`
	Amount         float64 `json:"amount"`
	CancelledAt    int     `json:"cancelled_at"`
	Reason         string  `json:"reason"`
	CancellationID string  `json:"cancellation_id"`
	ReceiptURL     *string `json:"receipt_url,omitempty"`
}

type PaymentCancelRequest struct {
	ImpUID        string   `json:"imp_uid"`
	MerchantUID   *string  `json:"merchant_uid"`
	Amount        *float64 `json:"amount"`
	TaxFree       *int     `json:"tax_free"`
	VatAmount     *int     `json:"vat_amount"`
	Checksum      *int     `json:"checksum"`
	Reason        *string  `json:"reason"`
	RefundHolder  *string  `json:"refund_holder"`
	RefundBank    *string  `json:"refund_bank"`
	RefundAccount *string  `json:"refund_account"`
	RefundTel     *string  `json:"refund_tel"`
	RetainPromo   *bool    `json:"retain_promotion"`
	Extra         []struct {
	} `json:"extra"`
}
