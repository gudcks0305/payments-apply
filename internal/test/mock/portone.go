package mock

import "github.com/gudcks0305/payments-apply/internal/portone"

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
			respBody.Response = portone.PaymentData{
				ImpUID:      id,
				MerchantUID: respBody.Response.MerchantUID,
				Amount:      10000,
				Status:      "paid",
			}
			return nil
		},
		MockCancelPayment: func(reqBody portone.PaymentCancelRequest, respBody *portone.APIResponse[portone.PaymentData]) error {
			respBody.Code = 0
			respBody.Message = "OK"
			respBody.Response = portone.PaymentData{
				ImpUID: reqBody.ImpUID,
				Amount: 10000,
			}
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
