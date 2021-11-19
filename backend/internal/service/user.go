//go:generate mockgen -source=user.go -destination=user_mock.go -package=service
package service

import (
	"context"

	"github.com/sirupsen/logrus"

	"junction-brella/internal/db"
	"junction-brella/internal/model/convert"
	"junction-brella/internal/model/core"
	"junction-brella/internal/model/dto"
)

type UserService interface {
	GetUserData(ctx context.Context, userID core.UserID) (*dto.GetUserDataResponse, error)

	SearchUsers(ctx context.Context, selector string) (*dto.SearchUsersResponse, error)
}

type userServiceImpl struct {
	log *logrus.Entry
	db  *db.Repository
}

func (svc *userServiceImpl) GetUserData(ctx context.Context, userID core.UserID) (*dto.GetUserDataResponse, error) {
	user, err := svc.db.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &dto.GetUserDataResponse{User: convert.User2DTO(user)}, nil
}
func (svc *userServiceImpl) SearchUsers(ctx context.Context, selector string) (*dto.SearchUsersResponse, error) {
	var usersCore []core.User

	usersCore, err := svc.db.UserRepo.SelectUsers(ctx, selector)
	if err != nil {
		return nil, err
	}

	users := []dto.User{}
	for _, user := range usersCore {
		users = append(users, convert.User2DTO(&user))
	}

	return &dto.SearchUsersResponse{Users: users}, nil
}

func NewUserService(log *logrus.Entry, db *db.Repository) UserService {
	return &userServiceImpl{log: log, db: db}
}
