package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gudcks0305/payments-apply/internal/test/integration"
	"github.com/gudcks0305/payments-apply/internal/test/mock"
	"github.com/gudcks0305/payments-apply/pkg/logger"

	"github.com/gudcks0305/payments-apply/internal/dto"
	"github.com/gudcks0305/payments-apply/internal/portone"
	"github.com/stretchr/testify/assert"
)

var Log = logger.Log

func TestPaymentFlow(t *testing.T) {
	engine, _, app := integration.SetupGinApp(t)
	app.RequireStart()
	defer app.RequireStop()

	t.Run("Initialize Payment", func(t *testing.T) {
		payload := dto.PaymentCreateRequest{
			Amount:      10000,
			ProductName: "테스트 상품",
			PayMethod:   "card",
		}

		jsonData, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/payments", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response dto.APIResponse[dto.IdResponse[string]]
		json.Unmarshal(w.Body.Bytes(), &response)

		merchantUid := response.Data.ID
		assert.NotEmpty(t, merchantUid)

		t.Run("Complete Payment", func(t *testing.T) {
			paidMock := mock.MockPayData[mock.PaidMock]
			completePayload := portone.PaymentClientResponse{
				ImpUid:        paidMock.ImpUID,
				MerchantUid:   merchantUid,
				PayMethod:     "card",
				PaidAmount:    paidMock.Amount,
				Status:        "paid",
				Name:          "당근 10kg",
				PgProvider:    "kcp",
				PgTid:         "22336466628585",
				BuyerName:     "포트원 기술지원팀",
				BuyerEmail:    "",
				BuyerTel:      "010-1234-5678",
				BuyerAddr:     "서울특별시 강남구 삼성동",
				BuyerPostcode: "123-456",
				PaidAt:        1648344363,
				Success:       true,
			}

			jsonData, _ := json.Marshal(completePayload)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPut, "/api/v1/payments/"+merchantUid+"/complete", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			engine.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			Log.Info(w.Body.String())

			t.Run("Get Payment By ImpUID", func(t *testing.T) {
				w := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodGet, "/api/v1/payments/imp/"+paidMock.ImpUID, nil)

				engine.ServeHTTP(w, req)

				assert.Equal(t, http.StatusOK, w.Code)

				var response dto.APIResponse[portone.PaymentData]
				json.Unmarshal(w.Body.Bytes(), &response)

				assert.Equal(t, paidMock.ImpUID, response.Data.ImpUID)

				Log.Info("결제 조회 성공: " + w.Body.String())
			})

		})
	})
}
func TestAdditionalPaymentCases(t *testing.T) {
	engine, _, app := integration.SetupGinApp(t)
	app.RequireStart()
	defer app.RequireStop()

	t.Run("Initialize Payment with Empty Request Body", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/payments", bytes.NewBuffer([]byte{}))
		req.Header.Set("Content-Type", "application/json")

		engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		Log.Info("빈 요청 본문으로 결제 실패: " + w.Body.String())
	})

	t.Run("Initialize Payment with Invalid JSON", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/payments", bytes.NewBuffer([]byte("{invalid json}")))
		req.Header.Set("Content-Type", "application/json")

		engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		Log.Info("잘못된 형식의 JSON으로 결제 실패: " + w.Body.String())
	})
	t.Run("Initialize Payment with cancel basic", func(t *testing.T) {
		payload := dto.PaymentCreateRequest{
			Amount:      10000,
			ProductName: "테스트 상품",
			PayMethod:   "card",
		}

		jsonData, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/payments", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response dto.APIResponse[dto.IdResponse[string]]
		json.Unmarshal(w.Body.Bytes(), &response)

		merchantUid := response.Data.ID
		assert.NotEmpty(t, merchantUid)

		t.Run("Cancel paid Payment By ImpUID", func(t *testing.T) {
			paidMock := mock.MockPayData[mock.PaidMock]
			w := httptest.NewRecorder()
			basicCompletePayload := dto.PaymentBasicConfirmRequest{
				ImpUID: paidMock.ImpUID,
			}
			jsonData, _ := json.Marshal(basicCompletePayload)
			req, _ := http.NewRequest(http.MethodPut, "/api/v1/payments/complete", bytes.NewBuffer(jsonData))

			engine.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)

			var response dto.APIResponse[portone.PaymentData]
			json.Unmarshal(w.Body.Bytes(), &response)

			assert.Equal(t, "cancelled", response.Data.Status)
			Log.Info("결제 취소 성공: " + w.Body.String())
		})
	})
}
