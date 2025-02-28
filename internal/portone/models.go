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
