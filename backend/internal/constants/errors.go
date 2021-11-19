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
	ErrPhoneCodeMismach = &CodedError{errors.New("phone code mismatch"), http.StatusUnauthorized}

	ErrUnexpectedSigningMethod = &CodedError{errors.New("unexpected signing method"), http.StatusUnauthorized}
	ErrEmailNotConfirmed       = &CodedError{errors.New("email is not confirmed"), http.StatusUnauthorized}
	ErrCertificateNotConfirmed = &CodedError{errors.New("certificate is not confirmed"), http.StatusUnauthorized}
	ErrUserIDMismatch          = &CodedError{errors.New("user id mismatch"), http.StatusUnauthorized}

	// Forbidden
	ErrTokenExpired        = &CodedError{errors.New("token is expired"), http.StatusForbidden}
	ErrAuthTokenExpired    = &CodedError{errors.New("authorization token is expired"), http.StatusForbidden}
	ErrAuthorityMismatch   = &CodedError{errors.New("authority mismatch"), http.StatusForbidden}
	ErrUserMismatch        = &CodedError{errors.New("user mismatch"), http.StatusForbidden}
	ErrInvalidUserAuthType = &CodedError{errors.New("invalid user auth type"), http.StatusForbidden}

	// Internal
	ErrSignToken      = &CodedError{errors.New("failed to sign token"), http.StatusInternalServerError}
	ErrParseToken     = &CodedError{errors.New("failed to parse token"), http.StatusInternalServerError}
	ErrParseAuthToken = &CodedError{errors.New("failed to parse authorization token"), http.StatusInternalServerError}
	ErrWritePOAXML    = &CodedError{
		errors.New("failed to write POA contents to the response"),
		http.StatusInternalServerError,
	}
	ErrGenerateUUID    = &CodedError{errors.New("failed to generate UUID"), http.StatusInternalServerError}
	ErrGetFileFromForm = &CodedError{
		errors.New("ошибка при получении файла из формы"),
		http.StatusInternalServerError,
	}
	ErrOpenFile                = &CodedError{errors.New("ошибка при открытии файла"), http.StatusInternalServerError}
	ErrReadFile                = &CodedError{errors.New("ошибка при чтении файла"), http.StatusInternalServerError}
	ErrUnimplementedListOption = &CodedError{
		errors.New("given /poa/list option is not implemented"),
		http.StatusInternalServerError,
	}

	// Bad Request
	ErrBindRequest     = &CodedError{errors.New("failed to bind request"), http.StatusBadRequest}
	ErrValidateRequest = &CodedError{errors.New("failed to validate request"), http.StatusBadRequest}

	ErrDBNotFound = &CodedError{errors.New("not found in the database"), http.StatusBadRequest}

	ErrEmptyType        = &CodedError{errors.New("type is empty"), http.StatusBadRequest}
	ErrEmptyToken       = &CodedError{errors.New("token is empty"), http.StatusBadRequest}
	ErrEmptySelector    = &CodedError{errors.New("selector is empty"), http.StatusBadRequest}
	ErrEmptyHandshakeID = &CodedError{errors.New("handshake id is empty"), http.StatusBadRequest}

	ErrInvalidRelationType = &CodedError{errors.New("invalid relation type"), http.StatusBadRequest}
	ErrInvalidQuerySort    = &CodedError{errors.New("invalid query sort field"), http.StatusBadRequest}

	ErrRecipientDoesNotExist      = &CodedError{errors.New("recipient does not exist"), http.StatusBadRequest}
	ErrRepresentativeDoesNotExist = &CodedError{errors.New("representative does not exist"), http.StatusBadRequest}

	ErrUnsupportedAcceptMIMEType = &CodedError{
		errors.New("unsupported mime type in the accept header"),
		http.StatusNotAcceptable,
	}
	ErrParsePOAFileName    = &CodedError{errors.New("failed to parse poa file name"), http.StatusBadRequest}
	ErrValidatePOAFileName = &CodedError{errors.New("failed to validate poa file name"), http.StatusBadRequest}
	ErrPOAXMLFileTooBig    = &CodedError{
		errors.New("XML-файл доверенности имеет размер больше мегабайта"),
		http.StatusBadRequest,
	}

	// Conflict
	ErrDuplicateCertificateThumbprint = &CodedError{errors.New("duplicate certificate thumbprint"), http.StatusConflict}
	ErrDuplicateSignedPOA             = &CodedError{errors.New("duplicate signed poa"), http.StatusConflict}
	ErrDuplicatePOARelation           = &CodedError{errors.New("duplicate poa relation"), http.StatusConflict}
	ErrDuplicateRelation              = &CodedError{errors.New("duplicate relation"), http.StatusConflict}
	ErrEmailAlreadyConfirmed          = &CodedError{errors.New("email is already confirmed"), http.StatusConflict}
	ErrEmailLocked                    = &CodedError{
		errors.New("Ваш работодатель запретил вам менять свой email"),
		http.StatusConflict,
	}
	ErrEmailAlreadyTaken = &CodedError{
		errors.New("email is taken already by other user"),
		http.StatusConflict,
	}
	ErrPOAAlreadySigned        = &CodedError{errors.New("poa is already signed"), http.StatusConflict}
	ErrPOARepresentativeDiffer = &CodedError{errors.New("poa representative is different"), http.StatusConflict}
	ErrPOARecipientDiffer      = &CodedError{errors.New("poa recipient is different"), http.StatusConflict}
	ErrEmptyUserAuthType       = &CodedError{errors.New("empty user auth type"), http.StatusConflict}
	ErrUninitializedLogin      = &CodedError{errors.New("login is not initialized"), http.StatusBadRequest}
)
