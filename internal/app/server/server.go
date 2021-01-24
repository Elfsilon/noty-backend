package server

import (
	"github.com/Elfsilon/noty-backend/internal/app/store"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Server ...
type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New ...
func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *Server) Start() error {
	if err := s.configLogger(); err != nil {
		return err
	}

	s.configRouter()

	if err := s.configStore(); err != nil {
		return err
	}

	s.logger.Info("server is starting")

	return http.ListenAndServe(s.config.Addr, s.router)
}

func (s *Server) configLogger() error {
	lvl, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(lvl)

	return nil
}

func (s *Server) configRouter() {
	s.router.HandleFunc("/hello", s.helloFunc())
}

func (s *Server) configStore() error {
	store := store.New(s.config.Store)
	if err := store.Open(); err != nil {
		return err
	}

	s.store = store

	return nil
}
