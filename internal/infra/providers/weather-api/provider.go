package weather_api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tclemos/go-expert-cloudrun/pkg/entity"
)

var (
	ErrNotFound = fmt.Errorf("local não encontrado")
)

type WeatherProvider struct {
	apiKey string
}

func NewWeatherProvider(apiKey string) *WeatherProvider {
	return &WeatherProvider{
		apiKey: apiKey,
	}
}

func (wp *WeatherProvider) Get(ctx context.Context, city string) (*entity.Weather, error) {
	urlAddr := "http://api.weatherapi.com/v1/current.json?key=" + wp.apiKey + "&q=" + url.QueryEscape(city)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlAddr, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	if code, found := m["code"]; found && code.(int) == 1006 {
		return nil, ErrNotFound
	}

	tempC := m["current"].(map[string]interface{})["temp_c"].(float64)

	return &entity.Weather{TempC: tempC}, nil
}
