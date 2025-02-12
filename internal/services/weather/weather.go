package weather

import (
	"context"
	"fmt"

	"github.com/tclemos/go-expert-cloudrun/pkg/entity"
)

var (
	ErrNotFound = fmt.Errorf("não encontrado")
)

type CepProvider interface {
	Get(context.Context, entity.Cep) (string, error)
}

type WeatherProvider interface {
	Get(context.Context, string) (*entity.Weather, error)
}

type Service struct {
	cepProvider     CepProvider
	weatherProvider WeatherProvider
}

func NewService(cepProvider CepProvider, weatherProvider WeatherProvider) *Service {
	return &Service{
		cepProvider:     cepProvider,
		weatherProvider: weatherProvider,
	}
}

func (s *Service) Get(ctx context.Context, cep entity.Cep) (*entity.Weather, error) {
	resp, err := s.cepProvider.Get(ctx, cep)
	if err != nil {
		return nil, err
	}

	weather, err := s.weatherProvider.Get(ctx, resp)
	if err != nil {
		return nil, err
	}

	return weather, nil
}
