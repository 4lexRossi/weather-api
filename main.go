package main

import (
	"log"
	"net/http"

	"github.com/4lexRossi/weather-api/handler"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}
	http.HandleFunc("/weather", handler.WeatherHandler)

	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
