package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Configs struct {
	WeatherAPIKey string `mapstructure:"WEATHERAPI_KEY"`
}

func LoadConfig() (*Configs, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Arquivo .env não encontrado, tentando variáveis de ambiente")
	}

	var cfg Configs
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
