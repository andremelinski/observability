package cep_repository

import (
	"context"

	utils_interface "github.com/andremelinski/observability/cep/internal/pkg/utils/interface"
)

type LocationOutputDTO struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	DDD         string `json:"ddd"`
}

type LocationRepository struct {
	cepInfo utils_interface.ICepInfoAPI
}

func NewLocationRepository(cepInfo utils_interface.ICepInfoAPI) *LocationRepository {
	return &LocationRepository{
		cepInfo,
	}
}

func (l *LocationRepository) GetLocationInfo(ctx context.Context, cep string) (*LocationOutputDTO, error) {
	cepnfo, err := l.cepInfo.GetCEPInfo(ctx, cep)

	if err != nil {
		return nil, err
	}

	return &LocationOutputDTO{
		Cep:         cepnfo.Cep,
		Logradouro:  cepnfo.Logradouro,
		Complemento: cepnfo.Complemento,
		Bairro:      cepnfo.Bairro,
		Localidade:  cepnfo.Localidade,
		UF:          cepnfo.UF,
		DDD:         cepnfo.DDD,
	}, nil
}
