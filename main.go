package main

import (
	"github.com/maria-robobug/animal-api/internal/rest/dog"
	"log"
	"net/http"
	"os"
	"time"

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

type config struct {
	DogAPIKey  string `env:"DOG_API_KEY"`
	ServerPort string `env:"PORT" envDefault:"8080"`
}

func main() {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		errorLog.Fatalln("Failed to parse ENV")
	}

	if err := godotenv.Load(); err != nil {
		errorLog.Println("File .env not found, reading configuration from ENV")
	}

	serv := setupServer(cfg)

	errorLog.Fatal(serv.Run())
}

func setupServer(cfg config) *dog.Service {
	servConfig := &dog.Config{
		Client:   setupClient(cfg),
		Addr:     ":" + cfg.ServerPort,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	serv, err := dog.NewServer(servConfig)
	if err != nil {
		errorLog.Fatalf("could not initialise server: %s", err)
	}

	return serv
}

func setupClient(cfg config) *dog.Client {
	// Client configuration
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client, err := dog.NewClient(dogapiBaseURL, cfg.DogAPIKey, &http.Client{Transport: tr, Timeout: time.Second * 30})
	if err != nil {
		errorLog.Fatalf("could not create dog client: %s", err)
	}

	return client
}
