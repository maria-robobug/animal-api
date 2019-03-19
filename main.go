package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/maria-robobug/animal-api/client"
	"github.com/maria-robobug/animal-api/server"
)

const (
	dogapiBaseURL = "https://api.thedogapi.com/v1"
)

func main() {
	// Logging configuration
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Client configuration
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	apiKey := os.Getenv("DOG_API_KEY")
	if apiKey == "" {
		errorLog.Fatalf("config error: please provide DOG_API_KEY")
	}

	client, err := client.New(dogapiBaseURL, apiKey, &http.Client{Transport: tr, Timeout: time.Second * 30})
	if err != nil {
		errorLog.Fatalf("could not create dog client: %s", err)
	}

	// Server configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cnfg := &server.Config{
		Client:   client,
		Addr:     ":" + port,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	serv, err := server.New(cnfg)
	if err != nil {
		errorLog.Fatalf("could not initialise server: %s", err)
	}

	errorLog.Fatal(serv.Run())
}
