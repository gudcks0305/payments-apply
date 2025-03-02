package integration_test

import (
	"bytes"
	"encoding/json"
	"github.com/gudcks0305/payments-apply/internal/test/integration"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gudcks0305/payments-apply/internal/dto"
	"github.com/gudcks0305/payments-apply/internal/portone"
	"github.com/stretchr/testify/assert"
)

func TestPaymentFlow(t *testing.T) {
	engine, _, app := integration.SetupGinApp(t)

	app.RequireStart()
	defer app.RequireStop()

	// 1. 결제 초기화 테스트
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

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)

		merchantUid, exists := response["merchantUid"]
		assert.True(t, exists)
		assert.NotEmpty(t, merchantUid)

		// 다음 테스트에서 사용할 merchantUid 저장
		merchantUidStr := merchantUid.(string)

		// 2. 결제 완료 테스트
		t.Run("Complete Payment", func(t *testing.T) {
			completePayload := portone.PaymentClientResponse{
				ImpUid:        "imp_347242536261",
				MerchantUid:   merchantUidStr,
				PayMethod:     "card",
				PaidAmount:    1004,
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
			req, _ := http.NewRequest("POST", "/api/v1/payments/complete", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			engine.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)

			var response map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &response)

			success, exists := response["success"]
			assert.True(t, exists)
			assert.True(t, success.(bool))
		})
	})
}
