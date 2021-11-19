//go:generate mockgen -source=user.go -destination=user_mock.go -package=service
package service

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"gitlab.com/procsy/attorney/api/internal/constants"
	"gitlab.com/procsy/attorney/api/internal/db"
	"gitlab.com/procsy/attorney/api/internal/model/core"
	"gitlab.com/procsy/attorney/api/internal/model/dto"
	"gitlab.com/procsy/attorney/api/internal/utils"
)

type AuthService interface {
	Register(ctx context.Context, request *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error)
	Login(ctx context.Context, request *dto.LoginUserRequest) (*dto.LoginUserResponse, error)
}

type AuthServiceImpl struct {
	log *logrus.Entry
	db  *db.Repository
}

func (svc *AuthServiceImpl) Register(
	ctx context.Context, request *dto.RegisterUserRequest,
) (*dto.RegisterUserResponse, error) {
	if exists, err := svc.db.UserRepo.CheckUserEmailExistence(ctx, request.Email); err != nil {
		return nil, fmt.Errorf("check user email existence: %w", err)
	} else if exists {
		return nil, constants.ErrEmailAlreadyTaken
	}

	user := &core.User{
		Email:    request.Email,
		Username: request.Username,
	}

	if err := user.Password.Init(request.Password); err != nil {
		return nil, err
	}

	if err := svc.db.UserRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	authToken, err := utils.GenerateAuthToken(&utils.AuthTokenWrapper{UserID: user.ID})
	if err != nil {
		return nil, err
	}

	return &dto.RegisterUserResponse{AuthToken: authToken}, nil
}

func (svc *AuthServiceImpl) Login(
	ctx context.Context,
	request *dto.LoginUserRequest,
) (*dto.LoginUserResponse, error) {
	user, err := svc.db.UserRepo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if err := user.Password.Validate(request.Password); err != nil {
		return nil, err
	}

	authToken, err := utils.GenerateAuthToken(&utils.AuthTokenWrapper{UserID: user.ID})
	if err != nil {
		return nil, err
	}

	return &dto.LoginUserResponse{AuthToken: authToken}, nil
}

func NewAuthService(log *logrus.Entry, db *db.Repository) AuthService {
	return &AuthServiceImpl{log: log, db: db}
}
