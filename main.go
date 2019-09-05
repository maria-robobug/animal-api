package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/maria-robobug/animal-api/internal/client"
	"github.com/maria-robobug/animal-api/server"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var (
	// Logging configuration
	infoLog  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
)

type appConfig struct {
	DogAPIKey     string `env:"DOG_API_KEY"`
	DogAPIBaseURI string `env:"DOG_API_BASE_URI"`
	ServerPort    string `env:"PORT" envDefault:"8080"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		errorLog.Println("file .env not found, reading configuration from ENV")
	}

	var cfg appConfig
	if err := env.Parse(&cfg); err != nil {
		errorLog.Fatalln("failed to parse ENV")
	}

	serv := setupServer(cfg)
	errorLog.Fatal(serv.Run())
}

func setupServer(cfg appConfig) *server.AnimalAPIServer {
	servConfig := &server.Config{
		DogAPIClient: setupClient(cfg),
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

func setupClient(cfg appConfig) *client.DogAPIClient {
	// Client configuration
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client, err := client.New(cfg.DogAPIBaseURI, cfg.DogAPIKey, &http.Client{Transport: tr, Timeout: time.Second * 30})
	if err != nil {
		errorLog.Fatalf("could not create dog client: %s", err)
	}

	return client
}
