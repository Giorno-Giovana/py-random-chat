package api

import (
	"fmt"

	"junction-brella/internal/constants"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type validatorImpl struct {
	validator *validator.Validate
}

func (v *validatorImpl) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func NewValidator() echo.Validator {
	return &validatorImpl{validator: validator.New()}
}

type binderImpl struct{}

func (b *binderImpl) Bind(i interface{}, ctx echo.Context) error {
	db := new(echo.DefaultBinder)
	if err := db.Bind(i, ctx); err != nil {
		return fmt.Errorf("%w: %v", constants.ErrBindRequest, err)
	}

	if err := ctx.Validate(i); err != nil {
		return fmt.Errorf("%w: %v", constants.ErrValidateRequest, err)
	}

	return nil
}

func NewBinder() echo.Binder {
	return &binderImpl{}
}
