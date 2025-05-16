package main

import (
	"fmt"
	"net/http"

	"event-tracker/internal/handler"
	"event-tracker/internal/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error while reading config: %w", err))
	}
}

func main() {
	initConfig()

	logLevel := viper.GetString("log.level")
	log, err := logger.NewLogger(logLevel)
	if err != nil {
		panic(fmt.Errorf("failed to init logger: %w", err))
	}
	defer log.Sync()

	http.HandleFunc("/event", handler.MakeEventHandler(log))

	port := viper.GetString("server.port")
	log.Info("Starting server", zap.String("port", port))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server failed", zap.Error(err))
	}
}
