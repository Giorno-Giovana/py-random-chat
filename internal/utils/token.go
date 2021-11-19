package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"gitlab.com/procsy/attorney/api/internal/constants"
	"gitlab.com/procsy/attorney/api/internal/model/core"
)

// TokenWrapper wraps the jwt library making it easy and clean
// to generate, validate and parse jwt tokens
type TokenWrapper struct {
	Email  string      `json:"email,omitempty"`
	UserID core.UserID `json:"user_id,omitempty"`

	jwt.StandardClaims
}

func GenerateToken(tw *TokenWrapper) (string, error) {
	if tw.ExpiresAt == 0 {
		t := time.Second * time.Duration(viper.GetInt64(constants.ViperJWTTTLKey))
		tw.ExpiresAt = time.Now().Add(t).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tw)
	secret, err := token.SignedString([]byte(viper.GetString(constants.ViperJWTSecretKey)))
	if err != nil {
		return "", fmt.Errorf("%w: %v", constants.ErrSignToken, err)
	}

	return base64.URLEncoding.EncodeToString([]byte(secret)), nil
}

func ParseToken(base64Secret string) (*TokenWrapper, error) {
	buffer, err := base64.URLEncoding.DecodeString(base64Secret)
	if err != nil {
		return nil, fmt.Errorf("error decoding from base64: %w", err)
	}

	return parseToken(string(buffer))
}

func parseToken(secret string) (*TokenWrapper, error) {
	t, err := jwt.ParseWithClaims(
		secret,
		&TokenWrapper{},
		keyFunc([]byte(viper.GetString(constants.ViperJWTSecretKey))),
	)

	if ve, ok := err.(*jwt.ValidationError); ok {
		// check if Expiration error was set
		if ve.Errors&jwt.ValidationErrorExpired == jwt.ValidationErrorExpired {
			return nil, constants.ErrTokenExpired
		} else {
			return nil, constants.ErrTokenInvalid
		}
	} else if err != nil {
		return nil, fmt.Errorf("%w: %v", constants.ErrParseToken, err)
	}

	tw, ok := t.Claims.(*TokenWrapper)
	if !ok {
		return nil, fmt.Errorf("%w: %v", constants.ErrTokenInvalid, errors.New("failed casting claims"))
	}

	return tw, nil
}
