package constants

import (
	"errors"
	"net/http"
)

type CodedError struct {
	err  error
	code int
}

func (ce *CodedError) Error() string {
	return ce.err.Error()
}

func (ce *CodedError) Code() int {
	return ce.code
}

var (
	// Unathorized
	ErrMissingAuthToken  = &CodedError{errors.New("missing authorization token"), http.StatusUnauthorized}
	ErrMissingAuthHeader = &CodedError{errors.New("missing authorization header"), http.StatusUnauthorized}

	ErrTokenInvalid     = &CodedError{errors.New("token is invalid"), http.StatusUnauthorized}
	ErrAuthTokenInvalid = &CodedError{errors.New("authorization token is invalid"), http.StatusUnauthorized}

	ErrPasswordMismatch = &CodedError{errors.New("password mismatch"), http.StatusUnauthorized}

	ErrUnexpectedSigningMethod = &CodedError{errors.New("unexpected signing method"), http.StatusUnauthorized}
	ErrEmailNotConfirmed       = &CodedError{errors.New("email is not confirmed"), http.StatusUnauthorized}
	ErrUserIDMismatch          = &CodedError{errors.New("user id mismatch"), http.StatusUnauthorized}

	// Forbidden
	ErrTokenExpired     = &CodedError{errors.New("token is expired"), http.StatusForbidden}
	ErrAuthTokenExpired = &CodedError{errors.New("authorization token is expired"), http.StatusForbidden}
	ErrUserMismatch     = &CodedError{errors.New("user mismatch"), http.StatusForbidden}

	// Internal
	ErrSignToken       = &CodedError{errors.New("failed to sign token"), http.StatusInternalServerError}
	ErrParseToken      = &CodedError{errors.New("failed to parse token"), http.StatusInternalServerError}
	ErrParseAuthToken  = &CodedError{errors.New("failed to parse authorization token"), http.StatusInternalServerError}
	ErrGenerateUUID    = &CodedError{errors.New("failed to generate UUID"), http.StatusInternalServerError}
	ErrGetFileFromForm = &CodedError{
		errors.New("error geting file from form"),
		http.StatusInternalServerError,
	}
	ErrOpenFile = &CodedError{errors.New("error opening file"), http.StatusInternalServerError}
	ErrReadFile = &CodedError{errors.New("error reading file"), http.StatusInternalServerError}

	// Bad Request
	ErrBindRequest     = &CodedError{errors.New("failed to bind request"), http.StatusBadRequest}
	ErrValidateRequest = &CodedError{errors.New("failed to validate request"), http.StatusBadRequest}

	ErrDBNotFound = &CodedError{errors.New("not found in the database"), http.StatusBadRequest}

	ErrEmptyType        = &CodedError{errors.New("type is empty"), http.StatusBadRequest}
	ErrEmptyToken       = &CodedError{errors.New("token is empty"), http.StatusBadRequest}
	ErrEmptySelector    = &CodedError{errors.New("selector is empty"), http.StatusBadRequest}
	ErrEmptyHandshakeID = &CodedError{errors.New("handshake id is empty"), http.StatusBadRequest}

	ErrInvalidQuerySort = &CodedError{errors.New("invalid query sort field"), http.StatusBadRequest}

	// Conflict
	ErrEmailAlreadyTaken = &CodedError{errors.New("email is taken already by other user"), http.StatusConflict}
)
