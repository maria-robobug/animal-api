package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/maria-robobug/animal-api/server"
	"github.com/maria-robobug/animal-api/server/client"
)

const (
	dogapiBaseURL = "https://api.thedogapi.com/v1"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	client, err := client.New(dogapiBaseURL, &http.Client{Timeout: time.Second * 30})
	if err != nil {
		errorLog.Fatalf("could not create dog client: %s", err)
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
