package portone

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gudcks0305/payments-apply/internal/config"
	"github.com/gudcks0305/payments-apply/pkg/logger"
)

type Client struct {
	config      *config.Config
	httpClient  *http.Client
	authService *AuthService
}

// NewClient 는 새 Client 인스턴스를 생성합니다
func NewClient(config *config.Config) *Client {
	authService := NewAuthService(config)
	return &Client{
		config:      config,
		httpClient:  &http.Client{Timeout: 10 * time.Second},
		authService: authService,
	}
}

func (c *Client) Do(method, path string, reqBody interface{}, respBody interface{}) error {
	return c.doWithRetry(method, path, reqBody, respBody, true)
}

func (c *Client) doWithRetry(method, path string, reqBody interface{}, respBody interface{}, allowRetry bool) error {
	token, err := c.authService.GetToken()
	if err != nil {
		return fmt.Errorf("토큰 획득 실패: %w", err)
	}

	url := fmt.Sprintf("%s%s", c.config.PortOne.BaseURL, path)

	var bodyBytes []byte
	if reqBody != nil {
		bodyBytes, err = json.Marshal(reqBody)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	logger.Log.Info("[PortOne API 요청] %s %s", method, path)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized && allowRetry {
		logger.Log.Warn("[PortOne API] 토큰 만료됨, 재발급 시도")
		c.authService.InvalidateToken()
		return c.doWithRetry(method, path, reqBody, respBody, false) // 재시도는 한 번만
	}

	// 응답 디코딩
	if respBody != nil {
		if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
			return err
		}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errResp := ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("API 오류 응답 (상태 코드: %d)", resp.Status+err.Error())
		}
		return errors.New(errResp.Message)
	}

	return nil
}

func (c *Client) Get(path string, respBody interface{}) error {
	return c.Do(http.MethodGet, path, nil, respBody)
}

func (c *Client) Post(path string, reqBody interface{}, respBody interface{}) error {
	return c.Do(http.MethodPost, path, reqBody, respBody)
}

func (c *Client) GetPayment(id string, respBody interface{}) error {
	return c.Get("/payments/"+id, respBody)
}
