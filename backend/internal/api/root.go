package api

import (
	"context"

	"junction-brella/internal/api/controllers"
	"junction-brella/internal/db"
	"junction-brella/internal/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIService struct {
	log    *logrus.Entry
	router *echo.Echo
	debug  bool
}

func (svc *APIService) Serve() {
	svc.log.Info("Starting HTTP server")
	listenAddr := viper.GetString("service.bind.address") + ":" + viper.GetString("service.bind.port")
	svc.log.Fatal(svc.router.Start(listenAddr))
}

func (svc *APIService) Shutdown(ctx context.Context) error {
	if err := svc.router.Shutdown(ctx); err != nil {
		svc.log.Fatal(err)
	}

	return nil
}

func NewAPIService(log *logrus.Entry, dbConn *mongo.Database, debug bool) (*APIService, error) {
	svc := &APIService{
		log:    log,
		router: echo.New(),
		debug:  debug,
	}

	svc.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	svc.router.Validator = NewValidator()
	svc.router.Binder = NewBinder()

	repository, err := db.NewRepository(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	registry := service.NewRegistry(log, repository)
	authCtrl := controllers.NewAuthController(log, registry)
	userCtrl := controllers.NewUserController(log, registry)
	roomCtrl := controllers.NewRoomController(log, registry)
	tinderCtrs := controllers.NewTinderController(log, registry)

	svc.router.HTTPErrorHandler = svc.httpErrorHandler
	svc.router.Use(svc.XRequestIDMiddleware(), svc.LoggingMiddleware())
	api := svc.router.Group("/api")

	authAPI := api.Group("/auth")

	authAPI.POST("/register", authCtrl.RegisterUser)
	authAPI.POST("/login", authCtrl.LoginUser)

	userAPI := api.Group("/user", svc.AuthMiddleware())

	userAPI.GET("/search", userCtrl.SearchUsers)

	roomAPI := api.Group("/room")

	roomAPI.GET("/join", roomCtrl.Join)
	roomAPI.GET("/list", roomCtrl.List)

	tinderAPI := api.Group("/tinder")

	tinderAPI.GET("/next", tinderCtrs.Next)
	tinderAPI.POST("/register", authCtrl.RegisterUser)

	return svc, nil
}
