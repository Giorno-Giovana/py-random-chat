package controllers

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gitlab.com/procsy/attorney/api/internal/model/dto"
	"gitlab.com/procsy/attorney/api/internal/service"
)

type AuthController struct {
	log      *logrus.Entry
	registry *service.Registry
}

func (c *AuthController) RegisterUser(ctx echo.Context) error {
	request := new(dto.RegisterUserRequest)
	if err := ctx.Bind(request); err != nil {
		return err
	}

	response, err := c.registry.AuthService.Register(context.Background(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *AuthController) LoginUser(ctx echo.Context) error {
	request := new(dto.LoginUserRequest)
	if err := ctx.Bind(request); err != nil {
		return err
	}

	response, err := c.registry.AuthService.Login(context.Background(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

func NewAuthController(log *logrus.Entry, registry *service.Registry) *AuthController {
	return &AuthController{log: log, registry: registry}
}
