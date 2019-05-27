package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/maria-robobug/animal-api/internal/storage"

	"github.com/maria-robobug/animal-api/internal/client"
	"github.com/maria-robobug/animal-api/server"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const (
	dogapiBaseURL = "https://api.thedogapi.com/v1"
)

var (
	// Logging configuration
	infoLog  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
)

type appConfig struct {
	DogAPIKey  string `env:"DOG_API_KEY"`
	ServerPort string `env:"PORT" envDefault:"8080"`
	RedisHost  string `env:"REDIS_HOST" envDefault:"localhost:6379"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		errorLog.Println("File .env not found, reading configuration from ENV")
	}

	var cfg appConfig
	if err := env.Parse(&cfg); err != nil {
		errorLog.Fatalln("Failed to parse ENV")
	}

	serv := setupServer(cfg)
	errorLog.Fatal(serv.Run())
}

func setupServer(cfg appConfig) *server.AnimalAPIServer {
	servConfig := &server.Config{
		Cache:        setupCache(cfg.RedisHost),
		DogAPIClient: setupClient(cfg.DogAPIKey),
		Addr:         ":" + cfg.ServerPort,
		InfoLog:      infoLog,
		ErrorLog:     errorLog,
	}

	serv, err := server.New(servConfig)
	if err != nil {
		errorLog.Fatalf("could not initialise server: %s", err)
	}

	return serv
}

func setupCache(host string) *storage.RedisCache {
	rc, err := storage.NewCache(host)
	if err != nil {
		errorLog.Fatalf("could not create redis cache: %s", err)
	}

	return rc
}

func setupClient(apiKey string) *client.DogAPIClient {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client, err := client.New(
		dogapiBaseURL,
		apiKey,
		&http.Client{
			Transport: tr,
			Timeout:   time.Second * 30,
		},
	)
	if err != nil {
		errorLog.Fatalf("could not create dog client: %s", err)
	}

	return client
}
