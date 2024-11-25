package config

import "os"

// Função para obter a chave da API do WeatherAPI
func GetWeatherAPIKey() string {
	return os.Getenv("WEATHER_API_KEY")
}
