package main

import (
	"log"
	"net/http"

	"github.com/4lexRossi/weather-api/handler"
)

func main() {
	http.HandleFunc("/weather", handler.WeatherHandler)

	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
