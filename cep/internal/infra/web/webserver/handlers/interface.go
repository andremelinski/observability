package handlers

import "net/http"

type CityWebHandlerInterface interface {
	CityTemperature(w http.ResponseWriter, r *http.Request)
}