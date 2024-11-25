# Weather API

Este é um sistema em Go que recebe um CEP, consulta o clima na cidade associada e retorna a temperatura em Celsius, Fahrenheit e Kelvin.

## Como rodar a aplicação

1. Clone o repositório.
2. Defina a variável de ambiente `WEATHER_API_KEY` com sua chave da API WeatherAPI.
3. Execute os comandos:

```bash
# Rodar localmente
go run main.go

# Ou, se preferir, use o Docker
docker build -t weather-api .
docker run -p 8080:8080 weather-api
