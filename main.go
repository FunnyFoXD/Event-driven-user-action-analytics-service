package main

import (
	"fmt"
	"net/http"
	"os"

	"event-tracker/internal/handler"
	"event-tracker/internal/kafka"
	"event-tracker/internal/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func initConfig() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/app/config")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error while reading config: %w", err))
	}
}

func main() {
	// Viper initConfig
	initConfig()

	// Logger
	logLevel := viper.GetString("log.level")
	log, err := logger.NewLogger(logLevel)
	if err != nil {
		panic(fmt.Errorf("failed to init logger: %w", err))
	}
	defer log.Sync()

	// Kafka producer
	broker := viper.GetString("kafka.broker")
	topic := viper.GetString("kafka.topic")
	producer := kafka.NewProducer(broker, topic)
	defer producer.Close()

	// HTTP handlers
	http.HandleFunc("/event", handler.MakeEventHandler(log, producer))

	//Start server
	port := viper.GetString("server.port")
	log.Info("Starting server", zap.String("port", port))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server failed", zap.Error(err))
	}
}
