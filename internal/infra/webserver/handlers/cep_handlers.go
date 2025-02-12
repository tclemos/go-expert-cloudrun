package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	via_cep "github.com/tclemos/go-expert-cloudrun/internal/infra/providers/via-cep"
	weather_api "github.com/tclemos/go-expert-cloudrun/internal/infra/providers/weather-api"
	"github.com/tclemos/go-expert-cloudrun/internal/services/weather"
	"github.com/tclemos/go-expert-cloudrun/pkg/entity"
)

type WeatherHandlers struct {
	WeatherService *weather.Service
}

func NewCepHandlers(weatherService *weather.Service) *WeatherHandlers {
	return &WeatherHandlers{
		WeatherService: weatherService,
	}
}

func (h *WeatherHandlers) GetWeather(w http.ResponseWriter, r *http.Request) {
	cepParam := chi.URLParam(r, "cep")
	cep, erro := entity.NewCep(cepParam)
	if erro != nil {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	weather, err := h.WeatherService.Get(r.Context(), *cep)
	if errors.Is(err, via_cep.ErrNotFound) || errors.Is(err, weather_api.ErrNotFound) {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		TempC float64 `json:"temp_C"`
		TempF float64 `json:"temp_F"`
		TempK float64 `json:"temp_K"`
	}{
		TempC: weather.TempC,
		TempF: weather.TempF(),
		TempK: weather.TempK(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
