package portone

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gudcks0305/payments-apply/internal/config"
)

type AuthService struct {
	config     *config.Config
	httpClient *http.Client
	tokenCache *TokenCache
}

func NewAuthService(config *config.Config) *AuthService {
	return &AuthService{
		config:     config,
		httpClient: &http.Client{Timeout: 10 * time.Second},
		tokenCache: NewTokenCache(),
	}
}

func (s *AuthService) GetToken() (string, error) {
	if token, ok := s.tokenCache.Get(); ok {
		return token, nil
	}

	return s.refreshToken()
}

func (s *AuthService) refreshToken() (string, error) {
	url := fmt.Sprintf("%s/users/getToken", s.config.PortOne.BaseURL)

	tokenReq := TokenRequest{
		ImpKey:    s.config.PortOne.ImpKey,
		ImpSecret: s.config.PortOne.ImpSecret,
	}

	reqBody, err := json.Marshal(tokenReq)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", err
	}

	if tokenResp.Code != 0 {
		return "", errors.New(tokenResp.Message)
	}

	expireTime := time.Unix(tokenResp.Response.ExpiredAt, 0)
	safeExpireTime := expireTime.Add(-1 * time.Minute)

	s.tokenCache.Set(tokenResp.Response.AccessToken, safeExpireTime)

	return tokenResp.Response.AccessToken, nil
}

func (s *AuthService) InvalidateToken() {
	s.tokenCache.Clear()
}
