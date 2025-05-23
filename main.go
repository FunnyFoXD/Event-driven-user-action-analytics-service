package main

import (
	"fmt"
	"net/http"

	"event-tracker/internal/db"
	"event-tracker/internal/handler"
	"event-tracker/internal/kafka"
	"event-tracker/internal/logger"
	"event-tracker/internal/router"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("No .env file found or failed to load it...")
	}

	viper.AutomaticEnv()
}

func main() {
	// Viper init
	initEnv()

	// Logger
	logLevel := viper.GetString("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	log, err := logger.NewLogger(logLevel)
	if err != nil {
		panic(fmt.Errorf("failed to init logger: %w", err))
	}
	defer log.Sync()

	// Kafka producer
	broker := viper.GetString("KAFKA_BROKER")
	topic := viper.GetString("KAFKA_TOPIC")
	producer := kafka.NewProducer(broker, topic)
	defer producer.Close()

	// Database init
	dbHost := viper.GetString("POSTGRES_HOST")
	dbPort := viper.GetString("POSTGRES_PORT")
	dbUser := viper.GetString("POSTGRES_USER")
	dbPass := viper.GetString("POSTGRES_PASSWORD")
	dbName := viper.GetString("POSTGRES_DB")

	database, err := db.InitDB(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		log.Fatal("Failed to connect to DB", zap.Error(err))
	}

	// Kafka consumer
	kafka.StartConsumer(broker, topic, database, log)

	// HTTP handlers
	http.HandleFunc("/event", handler.MakeEventHandler(log, producer))

	// Router init
	r := router.NewRouter(log, producer, database)

	//Start server
	port := viper.GetString("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Info("Starting server", zap.String("port", port))
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Server failed", zap.Error(err))
	}
}
