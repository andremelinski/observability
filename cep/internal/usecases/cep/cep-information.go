package usecases

import (
	utils_interface "github.com/andremelinski/observability/cep/internal/pkg/utils/interface"
	usecases_dto "github.com/andremelinski/observability/cep/internal/usecases/dto"
)



type LocationUseCase struct {
	cepInfo utils_interface.ICepInfoAPI
}

func NewLocationUseCase(cepInfo utils_interface.ICepInfoAPI)*LocationUseCase{
	return &LocationUseCase{
		cepInfo,
	}
}

func (l *LocationUseCase)GetLocationInfo(cep string) (*usecases_dto.LocationOutputDTO, error){
	cepnfo, err := l.cepInfo.GetCEPInfo(cep)

	if err != nil {
		return nil, err
	}

	return &usecases_dto.LocationOutputDTO{
		Cep: cepnfo.Cep, 
		Logradouro: cepnfo.Logradouro, 
		Complemento: cepnfo.Complemento,
		Bairro: cepnfo.Bairro,
		Localidade: cepnfo.Localidade,
		UF: cepnfo.UF,
		DDD: cepnfo.DDD, 
	}, nil
}