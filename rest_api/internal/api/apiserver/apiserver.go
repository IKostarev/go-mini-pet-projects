package apiserver

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	cfg    *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(cfg *Config) *APIServer {
	return &APIServer{
		cfg:    cfg,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configLogger(); err != nil {
		return err
	}

	s.configRouter()

	s.logger.Info("start api server")

	return http.ListenAndServe(s.cfg.BindAddr, s.router)
}

func (s *APIServer) configLogger() error {
	level, err := logrus.ParseLevel(s.cfg.LogLevel)
	if err != nil {
		log.Fatal("Ошибка при парсинге уровня логирования: ", err)
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configRouter() {
	s.router.HandleFunc("/hell", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hi")
	}
}
