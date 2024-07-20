package web

import (
	"net/http"

	"github.com/andremelinski/observability/cep/internal/infra/web/handlers"
)

type RouteHandler struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

type WebRouter struct {
	CityWebHandler       handlers.ICityWebHandler
}

func NewWebRouter(
	cityHandlers handlers.ICityWebHandler,
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