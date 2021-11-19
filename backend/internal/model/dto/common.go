package dto

type BasicResponse struct{}

type TokenResponse struct {
	Token string `json:"token"`
}

type AuthTokenResponse struct {
	AuthToken string `json:"auth_token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
