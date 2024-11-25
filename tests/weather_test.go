package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/4lexRossi/weather-api/handler"
)

// Testa o manipulador de requisições
func TestWeatherHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/weather?cep=01001-000", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.WeatherHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, but got %v", status)
	}

	// Verificar se a resposta é válida (opcional, pode ser ajustada)
	expected := `{"temp_C":28.5,"temp_F":83.3,"temp_K":301.65}`
	if rr.Body.String() != expected {
		t.Errorf("Expected body %v, but got %v", expected, rr.Body.String())
	}
}
