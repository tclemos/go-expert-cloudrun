package via_cep

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tclemos/go-expert-cloudrun/pkg/entity"
)

var (
	ErrNotFound = fmt.Errorf("cep não encontrado")
)

type CepProvider struct{}

func NewCepProvider() *CepProvider {
	return &CepProvider{}
}

func (p *CepProvider) Get(ctx context.Context, cep entity.Cep) (string, error) {
	cepInput := cep.String(false)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://viacep.com.br/ws/"+cepInput+"/json/", nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		return "", err
	}

	if _, found := m["erro"]; found {
		return "", ErrNotFound
	}

	cidade := m["localidade"].(string)

	return cidade, nil
}
