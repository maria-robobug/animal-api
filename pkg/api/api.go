package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maria-robobug/dogfacts/pkg/dog"
)

type Application struct {
	Client            dog.DogApi
	Server            *http.Server
	InfoLog, ErrorLog *log.Logger
}

func New(cli dog.DogApi, addr string, infoLog, errorLog *log.Logger) *Application {
	return &Application{
		Client: cli,
		Server: &http.Server{
			Addr:     addr,
			ErrorLog: errorLog,
		},
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
}

func (app *Application) Run() error {
	app.InfoLog.Printf("Starting server on %s", app.Server.Addr)
	return app.Server.ListenAndServe()
}

func (app *Application) RegisterRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/dog", app.GetRandomDog)
	app.Server.Handler = r
}
