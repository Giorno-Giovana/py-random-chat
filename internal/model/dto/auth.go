package dto

type RegisterUserRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterUserResponse AuthTokenResponse

type LoginUserRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginUserResponse AuthTokenResponse
