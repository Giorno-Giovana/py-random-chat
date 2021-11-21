package controllers

import (
	"context"
	"net/http"

	"junction-brella/internal/model/dto"
	"junction-brella/internal/service"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	log      *logrus.Entry
	registry *service.Registry
}

func (c *AuthController) RegisterUser(ctx echo.Context) error {
	request := new(dto.TinderRegisterUserRequest)
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
