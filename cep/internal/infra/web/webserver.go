package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
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

	for _, handler := range s.Handlers {
		s.Router.MethodFunc(handler.Method, handler.Path, handler.HandlerFunc)
	}
	http.ListenAndServe(fmt.Sprintf(":%d", s.WebServerPort), s.Router)
}