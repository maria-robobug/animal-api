package rest

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	errConfigNilClient   = errors.New("invalid config: nil client")
	errConfigMissingPort = errors.New("invalid config: missing address port")
	errLoggerMissing     = errors.New("invalid config: logger missing")
)

// Server interface
type Server interface {
	NewServer(cnfg *Config) (*Server, error)
	GetRandomDog(w http.ResponseWriter, r *http.Request)
}

// Service holds service client and server information
type Service struct {
	Client            DogAPI
	Server            *http.Server
	InfoLog, ErrorLog *log.Logger
}

// Config holds service config information
type Config struct {
	Client            DogAPI
	Addr              string
	InfoLog, ErrorLog *log.Logger
}

// NewServer returns a new service
func NewServer(cnfg *Config) (*Service, error) {
	if cnfg.Client == nil {
		return &Service{}, errConfigNilClient
	}

	if cnfg.Addr == "" {
		return &Service{}, errConfigMissingPort
	}

	if cnfg.InfoLog == nil || cnfg.ErrorLog == nil {
		return &Service{}, errLoggerMissing
	}

	return &Service{
		Client: cnfg.Client,
		Server: &http.Server{
			Addr:     cnfg.Addr,
			ErrorLog: cnfg.ErrorLog,
		},
		InfoLog:  cnfg.InfoLog,
		ErrorLog: cnfg.ErrorLog,
	}, nil
}

// Run starts the server
func (s *Service) Run() error {
	s.registerRoutes()

	s.InfoLog.Printf("Starting server on %s", s.Server.Addr)
	return s.Server.ListenAndServe()
}

func (s *Service) registerRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/dog", s.GetRandomDog)
	s.Server.Handler = r
}
