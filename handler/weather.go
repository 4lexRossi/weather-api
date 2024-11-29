package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/4lexRossi/weather-api/config"
	"github.com/4lexRossi/weather-api/models"
)

func validateCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{5}-\d{3}$`)
	return re.MatchString(cep)
}

func getCityFromCEP(cep string) (string, error) {
	if !validateCEP(cep) {
		return "", fmt.Errorf("invalid zipcode format accepted 00000-00")
	}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("can not find zipcode")
	}

	var viaCEPResp models.ViaCEPResponse
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &viaCEPResp)

	if viaCEPResp.Cidade == "" {
		return "", fmt.Errorf("can not find zipcode")
	}

	return viaCEPResp.Cidade, nil
}

func getWeather(city string) (float64, error) {
	apiKey := config.GetWeatherAPIKey()
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("could not fetch weather data")
	}

	var weatherResp models.WeatherAPIResponse
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &weatherResp)

	return weatherResp.Current.TempC, nil
}

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	if !validateCEP(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	city, err := getCityFromCEP(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tempC, err := getWeather(city)
	if err != nil {
		http.Error(w, "could not fetch weather data", http.StatusInternalServerError)
		return
	}

	tempF := celsiusToFahrenheit(tempC)
	tempK := celsiusToKelvin(tempC)

	response := map[string]float64{
		"temp_C": tempC,
		"temp_F": tempF,
		"temp_K": tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
