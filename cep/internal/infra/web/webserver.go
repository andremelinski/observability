package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type WebServer struct {
	Router        chi.Router
	Handlers      []RouteHandler
	WebServerPort int
}

func NewWebServer(
	serverPort int,
	handlers []RouteHandler,
) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      handlers,
		WebServerPort: serverPort,
	}
}

// loop through the handlers and add them to the router
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Timeout(60 * time.Second))
	// promhttp
	// gera um endpoint para que o prometheus pegue as infos
	s.Router.Handle("/metrics", promhttp.Handler())

	for _, handler := range s.Handlers {
		s.Router.MethodFunc(handler.Method, handler.Path, handler.HandlerFunc)
	}
	http.ListenAndServe(fmt.Sprintf(":%d", s.WebServerPort), s.Router)
}
