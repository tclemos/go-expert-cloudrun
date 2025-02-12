package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	via_cep "github.com/tclemos/go-expert-cloudrun/internal/infra/providers/via-cep"
	weather_api "github.com/tclemos/go-expert-cloudrun/internal/infra/providers/weather-api"
	"github.com/tclemos/go-expert-cloudrun/internal/infra/webserver/handlers"
	"github.com/tclemos/go-expert-cloudrun/internal/services/weather"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	cepProvider := via_cep.NewCepProvider()
	weatherProvider := weather_api.NewWeatherProvider("fe6af96660ea4d8497302401251202")
	weatherService := weather.NewService(cepProvider, weatherProvider)
	cepHandlers := handlers.NewCepHandlers(weatherService)

	r.Get("/cep/{cep}", cepHandlers.GetWeather)

	http.ListenAndServe(":8080", r)
}
