package weather_api

import (
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type conf struct {
	WeatherAPIKey string `mapstructure:"WEATHER_API_KEY"`
}

func LoadConfig() (*conf, error) {
	// Descobre o caminho absoluto para o diretório pkg/weather_api/
	_, currentFile, _, _ := runtime.Caller(0)
	configPath := filepath.Dir(currentFile)

	viper.SetConfigFile(filepath.Join(configPath, ".env"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg conf
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
