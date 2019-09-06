package main

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/maria-robobug/animal-api/internal/client"
	"github.com/maria-robobug/animal-api/server"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var (
	logger = logrus.New()
)

type appConfig struct {
	DogAPIKey     string `env:"DOG_API_KEY"`
	DogAPIBaseURI string `env:"DOG_API_BASE_URI"`
	ServerPort    string `env:"PORT" envDefault:"8080"`
}

func init() {
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.SetLevel(logrus.DebugLevel)
}

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Warnln("file .env not found, reading configuration from ENV")
	}

	var cfg appConfig
	if err := env.Parse(&cfg); err != nil {
		logger.Errorln("failed to parse ENV")
	}

	serv := setupServer(cfg)
	logger.Fatal(serv.Run())
}

func setupServer(cfg appConfig) *server.AnimalAPIServer {
	servConfig := &server.Config{
		DogAPIClient: setupClient(cfg),
		Addr:         ":" + cfg.ServerPort,
		Logger:       logger,
	}

	serv, err := server.New(servConfig)
	if err != nil {
		logrus.Errorf("could not initialise server: %s", err)
	}

	return serv
}

func setupClient(cfg appConfig) *client.DogAPIClient {
	// Client configuration
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client, err := client.New(cfg.DogAPIBaseURI, cfg.DogAPIKey, &http.Client{Transport: tr, Timeout: time.Second * 30})
	if err != nil {
		logrus.Errorf("could not create dog client: %s", err)
	}

	return client
}
