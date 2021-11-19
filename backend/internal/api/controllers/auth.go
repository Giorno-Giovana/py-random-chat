package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"junction-brella/internal/model/dto"
	"junction-brella/internal/service"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

func (c *AuthController) InitOAuthLoginUserRequest(ctx echo.Context) error {
	clientID := viper.GetString("oauth.vk.client_id")
	redirectURI := "http://localhost:8080/me"
	scopeTemp := strings.Join(viper.GetStringSlice("oauth.vk.scope"), "+")
	state := viper.GetString("oauth.vk.state")

	url := fmt.Sprintf(
		"https://oauth.vk.com/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=%s",
		clientID, redirectURI, scopeTemp, state,
	)

	response := &dto.InitOAuthLoginUserResponse{OAuthURL: url}
	return ctx.JSON(http.StatusOK, response)
}

func (c *AuthController) ConfirmOAuthLoginUserRequest(ctx echo.Context) error {
	state_temp := ctx.QueryParam("state")
	if state_temp == "" {
		return fmt.Errorf("state query param is not provided")
	} else if state_temp != viper.GetString("oauth.vk.state") {
		return fmt.Errorf("state query param do not match original one, got=%s", state_temp)
	}

	code := ctx.QueryParam("code")
	if code == "" {
		return fmt.Errorf("code query param is not provided")
	}

	redirectURI := "http://localhost:8080/me"
	clientID := viper.GetString("oauth.vk.client_id")
	clientSecret := viper.GetString("oauth.vk.client_secret")

	url := fmt.Sprintf("https://oauth.vk.com/access_token?grant_type=authorization_code&code=%s&redirect_uri=%s&client_id=%s&client_secret=%s",
		code, redirectURI, clientID, clientSecret)

	req, _ := http.NewRequest("POST", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	token := struct {
		AccessToken string `json:"access_token"`
	}{}

	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &token)
	url = fmt.Sprintf("https://api.vk.com/method/%s?v=5.124&access_token=%s", "users.get", token.AccessToken)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	response := &dto.ConfirmOAuthLoginUserResponse{Data: string(bytes)}
	return ctx.JSON(http.StatusOK, response)
}

func NewAuthController(log *logrus.Entry, registry *service.Registry) *AuthController {
	return &AuthController{log: log, registry: registry}
}
