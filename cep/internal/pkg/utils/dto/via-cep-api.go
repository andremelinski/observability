package utils_dto

type ViaCepDTO struct {
	Cep string `json:"cep"` 
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade string `json:"unidade"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF string `json:"uf"`
	IBGE string `json:"ibge"`
	Gia string `json:"gia"`
	DDD string `json:"ddd"`
	Siafi string `json:"siafi"`
}