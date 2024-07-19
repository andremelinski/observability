package web

import (
	"net/http"

	"github.com/andremelinski/go-gcp/internal/infra/web/webserver/handlers"
)

type RouteHandler struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

// struct recebe a interface que possui os endpoints desse usecase + middlewares
type WebRouter struct {
	CityWebHandler       handlers.CityWebHandlerInterface
}

func NewWebRouter(
	cityHandlers handlers.CityWebHandlerInterface,
) *WebRouter {
	return &WebRouter{
		cityHandlers,
	}
}


// metodo para cadastrar todas as rotas
func (s *WebRouter) BuildHandlers() []RouteHandler {
	return []RouteHandler{
		{
			Path:        "/",
			Method:      "GET",
			HandlerFunc: s.CityWebHandler.CityTemperature,
		},
	}
}