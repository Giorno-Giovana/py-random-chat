package controllers

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gitlab.com/procsy/attorney/api/internal/constants"
	"gitlab.com/procsy/attorney/api/internal/model/core"
	"gitlab.com/procsy/attorney/api/internal/service"
)

type UserController struct {
	log      *logrus.Entry
	registry *service.Registry
}

func (c *UserController) GetUserData(ctx echo.Context) error {
	userID := core.UserID(ctx.QueryParam("user_id"))
	if len(userID) == 0 {
		userID = core.UserID(ctx.Request().Header.Get("User-Id"))
	}

	response, err := c.registry.UserService.GetUserData(context.Background(), userID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *UserController) SearchUsers(ctx echo.Context) error {
	selector := ctx.QueryParam("selector")
	if len(selector) == 0 {
		return constants.ErrEmptySelector
	}

	response, err := c.registry.UserService.SearchUsers(context.Background(), selector)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

func NewUserController(log *logrus.Entry, registry *service.Registry) *UserController {
	return &UserController{log: log, registry: registry}
}
