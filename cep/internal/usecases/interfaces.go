package usecases

type ILocationInfo interface {
	GetLocationInfo(cep string) (*LocationOutputDTO, error)
}

type IWeatherInfo interface{
	GetTempByPlaceName(name string) (*TempDTO, error)
}