package service

import (
	"junction-brella/internal/db"

	"github.com/sirupsen/logrus"
)

type Registry struct {
	AuthService   AuthService
	UserService   UserService
	RoomService   RoomService
	TinderService TinderService
}

func NewRegistry(log *logrus.Entry, repository *db.Repository) *Registry {
	registry := new(Registry)

	registry.AuthService = NewAuthService(log, repository)
	registry.UserService = NewUserService(log, repository)
	registry.RoomService = NewRoomService(log)
	registry.TinderService = NewTinderService(log)

	return registry
}
