package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/maria-robobug/dogfacts/pkg/api"
	"github.com/maria-robobug/dogfacts/pkg/dog"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	client, err := dog.NewClient("https://api.thedogapi.com/v1", &http.Client{Timeout: time.Second * 30})
	if err != nil {
		errorLog.Fatalf("could not create dog client:\n%s", err)
	}

	app := api.New(client, *addr, infoLog, errorLog)
	app.RegisterRoutes()

	errorLog.Fatal(app.Run())
}
