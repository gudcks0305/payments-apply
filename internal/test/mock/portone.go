package mock

import (
	"github.com/google/uuid"
	"github.com/gudcks0305/payments-apply/internal/portone"
	"github.com/gudcks0305/payments-apply/internal/utils"
	"time"
)

type MockClient struct {
	MockDo            func(method, path string, reqBody interface{}, respBody interface{}) error
	MockGet           func(path string, respBody interface{}) error
	MockPost          func(path string, reqBody interface{}, respBody interface{}) error
	MockGetPayment    func(id string, respBody *portone.APIResponse[portone.PaymentData]) error
	MockCancelPayment func(reqBody portone.PaymentCancelRequest, respBody *portone.APIResponse[portone.PaymentData]) error
}

func NewMockClient() portone.POClient {
	return &MockClient{
		MockDo:   nil,
		MockGet:  nil,
		MockPost: nil,
		MockGetPayment: func(id string, respBody *portone.APIResponse[portone.PaymentData]) error {
			respBody.Code = 0
			respBody.Message = "OK"
			respBody.Response = MockPayData[id]
			return nil
		},
		MockCancelPayment: func(reqBody portone.PaymentCancelRequest, respBody *portone.APIResponse[portone.PaymentData]) error {
			respBody.Code = 0
			respBody.Message = "OK"
			respBody.Response = MockPayData[reqBody.ImpUID]
			return nil
		},
	}
}

func (m *MockClient) Do(method, path string, reqBody interface{}, respBody interface{}) error {
	if m.MockDo != nil {
		return m.MockDo(method, path, reqBody, respBody)
	}
	return nil
}

func (m *MockClient) Get(path string, respBody interface{}) error {
	if m.MockGet != nil {
		return m.MockGet(path, respBody)
	}
	return nil
}

func (m *MockClient) Post(path string, reqBody interface{}, respBody interface{}) error {
	if m.MockPost != nil {
		return m.MockPost(path, reqBody, respBody)
	}
	return nil
}

func (m *MockClient) GetPayment(id string, respBody *portone.APIResponse[portone.PaymentData]) error {
	if m.MockGetPayment != nil {
		return m.MockGetPayment(id, respBody)
	}
	return nil
}

func (m *MockClient) CancelPayment(reqBody portone.PaymentCancelRequest, respBody *portone.APIResponse[portone.PaymentData]) error {
	if m.MockCancelPayment != nil {
		return m.MockCancelPayment(reqBody, respBody)
	}
	return nil
}

type PaymentData = portone.PaymentData
type PaymentCancelAnnotation = portone.PaymentCancelAnnotation

const (
	PaidMock     = "paid_mock"
	CanceledMock = "canceled_mock"
	ReadyMock    = "ready_mock"
)

var MockPayData = map[string]PaymentData{
	"paid_mock": {
		ImpUID:            "paid_mock",
		MerchantUID:       uuid.New().String(),
		Amount:            25000,
		Status:            "paid",
		PayMethod:         utils.ToPointer("card"),
		Channel:           utils.ToPointer("pc"),
		PgProvider:        utils.ToPointer("kcp"),
		EmbPgProvider:     utils.ToPointer("kcp"),
		PgTid:             utils.ToPointer("txid_paid_mock"),
		PgID:              utils.ToPointer("pgid_paid_mock"),
		Escrow:            utils.ToPointer(false),
		ApplyNum:          utils.ToPointer("12345678"),
		BankCode:          utils.ToPointer("004"), // 국민은행 코드
		BankName:          utils.ToPointer("국민은행"),
		CardCode:          utils.ToPointer("008"), // BC카드 코드
		CardName:          utils.ToPointer("BC카드"),
		CardIssuerCode:    utils.ToPointer("008"),
		CardIssuerName:    utils.ToPointer("BC카드"),
		CardPublisherCode: utils.ToPointer("008"),
		CardPublisherName: utils.ToPointer("BC카드"),
		CardQuota:         utils.ToPointer(0),                     // 할부 0개월
		CardNumber:        utils.ToPointer("4123-4567-8901-2345"), // 마스킹된 카드번호
		CardType:          utils.ToPointer(1),                     // 신용카드
		Name:              utils.ToPointer("mockData-paid"),
		Currency:          "KRW",
		BuyerName:         utils.ToPointer("mock 구매자"),
		BuyerEmail:        utils.ToPointer("mock@example.com"),
		BuyerTel:          utils.ToPointer("010-1234-5678"),
		BuyerAddr:         utils.ToPointer("mock address"),
		BuyerPostcode:     utils.ToPointer("12345"),
		CustomData:        utils.ToPointer("custom data for paid"),
		UserAgent:         utils.ToPointer("Mozilla/5.0"),
		StartedAt:         utils.ToPointer(int(time.Now().Add(-time.Minute * 5).Unix())), // 결제 시작 시간 (5분 전)
		PaidAt:            utils.ToPointer(int(time.Now().Unix())),                       // 결제 완료 시간
		ReceiptURL:        utils.ToPointer("https://example.com/receipt/paid_mock"),
		CashReceiptIssued: utils.ToPointer(false),
		CustomerUID:       utils.ToPointer("customer_paid_mock"),
		CustomerUIDUsage:  utils.ToPointer("issue"),
	},
	"canceled_mock": {
		ImpUID:        "canceled_mock",
		MerchantUID:   uuid.New().String(),
		Amount:        15000,
		CancelAmount:  15000, // 취소 금액
		Status:        "canceled",
		PayMethod:     utils.ToPointer("vbank"),
		Channel:       utils.ToPointer("mobile"),
		PgProvider:    utils.ToPointer("html5_inicis"),
		EmbPgProvider: utils.ToPointer("html5_inicis"),
		PgTid:         utils.ToPointer("txid_canceled_mock"),
		PgID:          utils.ToPointer("pgid_canceled_mock"),
		Escrow:        utils.ToPointer(false),
		VbankCode:     utils.ToPointer("088"), // 신한은행 가상계좌 코드
		VbankName:     utils.ToPointer("신한은행"),
		VbankNum:      utils.ToPointer("110-123-456789"),
		VbankHolder:   utils.ToPointer("mock 입금자"),
		VbankDate:     utils.ToPointer(int(time.Now().Add(time.Hour * 24).Unix())), // 가상계좌 입금 만료일 (24시간 후)
		VbankIssuedAt: utils.ToPointer(int(time.Now().Add(-time.Hour).Unix())),     // 가상계좌 발급 시간 (1시간 전)
		CancelledAt:   utils.ToPointer(int(time.Now().Unix())),                     // 취소 완료 시간
		CancelReason:  utils.ToPointer("mock data cancel"),
		FailReason:    utils.ToPointer("mock fail reason"), // 취소 시 FailReason 도 추가 (실패 원인)
		Name:          utils.ToPointer("mockData-canceled"),
		Currency:      "KRW",
		BuyerName:     utils.ToPointer("mock 구매자"),
		BuyerTel:      utils.ToPointer("010-1234-5678"),
		BuyerAddr:     utils.ToPointer("mock address"),
		BuyerPostcode: utils.ToPointer("12345"),
		CustomData:    utils.ToPointer("custom data for canceled"),
		UserAgent:     utils.ToPointer("Mozilla/5.0 (Mobile)"),
		StartedAt:     utils.ToPointer(int(time.Now().Add(-time.Minute * 10).Unix())), // 결제 시작 시간 (10분 전)
		CancelHistory: []PaymentCancelAnnotation{ // 취소 history mock data
			{
				Amount:      15000,
				Reason:      "mock cancel reason",
				ReceiptURL:  utils.ToPointer("https://example.com/cancel_receipt"),
				CancelledAt: int(time.Now().Unix()),
			},
		},
		CancelReceiptURLs: []string{"https://example.com/cancel_receipt_url_1", "https://example.com/cancel_receipt_url_2"}, // 복수 취소 영수증 URL
		CashReceiptIssued: utils.ToPointer(true),                                                                            // 현금영수증 발급 여부
		CustomerUID:       utils.ToPointer("customer_canceled_mock"),
		CustomerUIDUsage:  utils.ToPointer("payment"),
	},
	"ready_mock": {
		ImpUID:           "ready_mock",
		MerchantUID:      uuid.New().String(),
		Amount:           30000,
		Status:           "ready", // ready 상태
		PayMethod:        utils.ToPointer("card"),
		Channel:          utils.ToPointer("app"),
		PgProvider:       utils.ToPointer("danal_tpay"),
		EmbPgProvider:    utils.ToPointer("danal_tpay"),
		PgTid:            utils.ToPointer("txid_ready_mock"),
		PgID:             utils.ToPointer("pgid_ready_mock"),
		Escrow:           utils.ToPointer(false),
		Name:             utils.ToPointer("mockData-ready"),
		Currency:         "KRW",
		BuyerName:        utils.ToPointer("mock 구매자"),
		BuyerAddr:        utils.ToPointer("mock address"),
		BuyerPostcode:    utils.ToPointer("12345"),
		CustomData:       utils.ToPointer("custom data for ready"),
		UserAgent:        utils.ToPointer("iOS App 1.0"),
		StartedAt:        utils.ToPointer(int(time.Now().Add(-time.Minute * 20).Unix())), // 결제 시작 시간 (20분 전)
		CustomerUID:      utils.ToPointer("customer_ready_mock"),
		CustomerUIDUsage: utils.ToPointer("none"),
	},
}
