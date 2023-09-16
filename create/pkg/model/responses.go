package model

type Response struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int16  `json:"expires_in"`
}

type ErrorResponse struct {
	Code       string `json:"code"`
	Date       string `json:"date"`
	Identifier string `json:"identifier"`
}
