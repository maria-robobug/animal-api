package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/maria-robobug/dogfacts/server"
	"github.com/maria-robobug/dogfacts/server/client"
)

const (
	dogapiBaseURL = "https://api.thedogapi.com/v1"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	client, err := client.New(dogapiBaseURL, &http.Client{Timeout: time.Second * 30})
	if err != nil {
		errorLog.Fatalf("could not create dog client: %s", err)
	}

	cnfg := &server.Config{
		Client:   client,
		Addr:     *addr,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	serv, err := server.New(cnfg)
	if err != nil {
		errorLog.Fatalf("could not initialise server: %s", err)
	}

	errorLog.Fatal(serv.Run())
}
