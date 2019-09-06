package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"

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
	DogAPIClient client.DogAPI
	Server       *http.Server
	Logger       *logrus.Logger
}

// Config holds service config information
type Config struct {
	DogAPIClient client.DogAPI
	Addr         string
	Logger       *logrus.Logger
}

// New returns a new service
func New(cnfg *Config) (*AnimalAPIServer, error) {
	if cnfg.DogAPIClient == nil {
		return &AnimalAPIServer{}, errConfigNilClient
	}

	if cnfg.Addr == "" {
		return &AnimalAPIServer{}, errConfigMissingPort
	}

	if cnfg.Logger == nil {
		return &AnimalAPIServer{}, errLoggerMissing
	}

	return &AnimalAPIServer{
		DogAPIClient: cnfg.DogAPIClient,
		Server:       &http.Server{Addr: cnfg.Addr},
		Logger:       cnfg.Logger,
	}, nil
}

// Run starts the server
func (s *AnimalAPIServer) Run() error {
	s.registerRoutes()

	s.Logger.Infof("starting server on %s", s.Server.Addr)
	return s.Server.ListenAndServe()
}

func (s *AnimalAPIServer) registerRoutes() {
	r := chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RequestID,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second), // Sets a timeout value on the request context (ctx)
	)

	r.Get("/", s.GetHealth)
	r.Get("/health", s.GetHealth)
	r.Get("/api/v1/dogs/random", s.GetRandomDog)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		s.Logger.Infof("route -> %s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		s.Logger.Errorf("logging err: %s\n", err.Error())
	}

	s.Server.Handler = r
}
