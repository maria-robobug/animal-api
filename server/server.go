package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/maria-robobug/animal-api/internal/client"
)

var (
	errConfigNilClient   = errors.New("invalid config: nil client")
	errConfigMissingPort = errors.New("invalid config: missing address port")
	errLoggerMissing     = errors.New("invalid config: logger missing")
)

// Server interface
type Server interface {
	New(cnfg *Config) (*Server, error)
	GetRandomDog(w http.ResponseWriter, r *http.Request)
}

// AnimalAPIServer holds service client and server information
type AnimalAPIServer struct {
	DogAPIClient      client.DogAPI
	Server            *http.Server
	InfoLog, ErrorLog *log.Logger
}

// Config holds service config information
type Config struct {
	DogAPIClient      client.DogAPI
	Addr              string
	InfoLog, ErrorLog *log.Logger
}

// New returns a new service
func New(cnfg *Config) (*AnimalAPIServer, error) {
	if cnfg.DogAPIClient == nil {
		return &AnimalAPIServer{}, errConfigNilClient
	}

	if cnfg.Addr == "" {
		return &AnimalAPIServer{}, errConfigMissingPort
	}

	if cnfg.InfoLog == nil || cnfg.ErrorLog == nil {
		return &AnimalAPIServer{}, errLoggerMissing
	}

	return &AnimalAPIServer{
		DogAPIClient: cnfg.DogAPIClient,
		Server: &http.Server{
			Addr:     cnfg.Addr,
			ErrorLog: cnfg.ErrorLog,
		},
		InfoLog:  cnfg.InfoLog,
		ErrorLog: cnfg.ErrorLog,
	}, nil
}

// Run starts the server
func (s *AnimalAPIServer) Run() error {
	s.registerRoutes()

	s.InfoLog.Printf("Starting server on %s", s.Server.Addr)
	return s.Server.ListenAndServe()
}

func (s *AnimalAPIServer) registerRoutes() {
	r := chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	r.Get("/api/v1/dogs/random", s.GetRandomDog)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		s.InfoLog.Printf("Route -> %s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		s.ErrorLog.Panicf("Logging err: %s\n", err.Error())
	}

	s.Server.Handler = r
}
