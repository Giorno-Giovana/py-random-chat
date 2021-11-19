package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"junction-brella/internal/api"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	viper.AutomaticEnv()

	viper.SetConfigFile(viper.GetString("CONFIG_PATH"))
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("fatal error config file: %s \n", err)
		os.Exit(1)
	}

	viper.SetDefault("service.bind.address", "0.0.0.0")
	viper.SetDefault("service.bind.port", "8080")

	log := logrus.New()

	formatter := logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	}

	var debug bool
	switch viper.GetString("logging.level") {
	case "warning":
		log.SetLevel(logrus.WarnLevel)
	case "notice":
		log.SetLevel(logrus.InfoLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
		debug = true
		formatter.PrettyPrint = true
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	log.SetFormatter(&formatter)

	log.Infof("log level: %s", log.Level.String())

	clientOptions := options.Client().ApplyURI(viper.GetString("db.connection_string"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("failed to establish mongo db connection: %s", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("failed to check (ping) mongo db connection: %s", err)
	}

	log.Info("Connected to MongoDB!")
	mongoDB := client.Database(viper.GetString("db.database"))

	svc, err := api.NewAPIService(
		logrus.NewEntry(log),
		mongoDB,
		debug,
	)
	if err != nil {
		log.Fatalf("error creating service instance: %s", err)
	}

	go svc.Serve()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("service.shutdown_timeout"))*time.Second,
	)
	defer cancel()

	if err := svc.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
