package models

type ViaCEPResponse struct {
	Bairro     string `json:"bairro"`
	Cidade     string `json:"localidade"`
	Estado     string `json:"uf"`
	Logradouro string `json:"logradouro"`
}
