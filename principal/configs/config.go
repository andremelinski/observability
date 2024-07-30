package configs

import "github.com/spf13/viper"

type conf struct {
	HTTP_PORT                   int    `mapstructure:"HTTP_PORT"`
	EXTERNAL_URL                string `mapstructure:"EXTERNAL_URL"`
	REQUEST_NAME_OTEL           string `mapstructure:"REQUEST_NAME_OTEL"`
	OTEL_SERVICE_NAME           string `mapstructure:"OTEL_SERVICE_NAME"`
	OTEL_EXPORTER_OTLP_ENDPOINT string `mapstructure:"OTEL_EXPORTER_OTLP_ENDPOINT"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
