package usecases_dto

type LocationOutputDTO struct{
	Cep string `json:"cep"` 
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF string `json:"uf"`
	DDD string `json:"ddd"`
}

type TempDTOOutput struct{
	Celsius float64 `json:"temp_C"`;
	Fahrenheit float64 `json:"temp_F"`;
	Kelvin float64 `json:"temp_K"`;
}