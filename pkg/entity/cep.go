package entity

import (
	"fmt"
	"regexp"
	"strings"
)

type Cep struct {
	prefix string
	suffix string
}

const cepPattern = `^\d{5}-*\d{3}$`

func NewCep(cep string) (*Cep, error) {
	prefix, suffix, err := parse(cep)
	if err != nil {
		return nil, err
	}
	return &Cep{prefix, suffix}, nil
}

func (c *Cep) String(formatted bool) string {
	if formatted {
		return fmt.Sprintf("%s-%s", c.prefix, c.suffix)
	}

	return fmt.Sprintf("%s%s", c.prefix, c.suffix)
}

func parse(cep string) (prefix string, suffix string, err error) {
	if len(cep) < 8 || len(cep) > 9 {
		err = fmt.Errorf("tamanho do cep inválido")
		return
	}

	if !regexp.MustCompile(cepPattern).MatchString(cep) {
		err = fmt.Errorf("formato do cep inválido")
		return
	}

	cepLimpo := strings.ReplaceAll(cep, "-", "")
	prefix = cepLimpo[:5]
	suffix = cepLimpo[5:]
	return
}
