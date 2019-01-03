package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/maria-robobug/dogfacts/pkg/dogapi"
)

type application struct {
	client   *dogapi.DogClient
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	client, err := dogapi.NewDogClient("https://api.thedogapi.com/v1", &http.Client{Timeout: time.Second * 30})
	if err != nil {
		errorLog.Fatalf("could not create dog client:\n%s", err)
	}

	app := &application{
		client,
		errorLog,
		infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", srv.Addr)
	errorLog.Fatal(srv.ListenAndServe())
}
