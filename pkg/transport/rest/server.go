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

type Server struct {
	Client            DogAPI
	Server            *http.Server
	InfoLog, ErrorLog *log.Logger
}

type Config struct {
	Client            DogAPI
	Addr              string
	InfoLog, ErrorLog *log.Logger
}

func NewServer(cnfg *Config) (*Server, error) {
	if cnfg.Client == nil {
		return &Server{}, errConfigNilClient
	}

	if cnfg.Addr == "" {
		return &Server{}, errConfigMissingPort
	}

	if cnfg.InfoLog == nil || cnfg.ErrorLog == nil {
		return &Server{}, errLoggerMissing
	}

	return &Server{
		Client: cnfg.Client,
		Server: &http.Server{
			Addr:     cnfg.Addr,
			ErrorLog: cnfg.ErrorLog,
		},
		InfoLog:  cnfg.InfoLog,
		ErrorLog: cnfg.ErrorLog,
	}, nil
}

func (s *Server) Run() error {
	s.registerRoutes()

	s.InfoLog.Printf("Starting server on %s", s.Server.Addr)
	return s.Server.ListenAndServe()
}

func (s *Server) registerRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/dog", s.GetRandomDog)
	s.Server.Handler = r
}
