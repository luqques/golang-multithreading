package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BrasilApiResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	channelBrasilApi := make(chan BrasilApiResponse)
	channelViaCep := make(chan ViaCepResponse)

	var cep string
	fmt.Print("Digite um CEP: ")
	fmt.Scan(&cep)

	go requestBrasilApi(cep, channelBrasilApi)
	go requestViaCep(cep, channelViaCep)

	select {
	case viaCepResp := <-channelViaCep:
		fmt.Println("ViaCEP Response:", viaCepResp)
	case brasilApiResp := <-channelBrasilApi:
		fmt.Println("Brasil API Response:", brasilApiResp)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: Sem respostas recebidas dentro do tempo limite")
	}
}

func requestBrasilApi(cep string, channel chan<- BrasilApiResponse) {
	resp, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	defer resp.Body.Close()

	var brasilApiResponse BrasilApiResponse
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &brasilApiResponse)
	if err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		return
	}

	channel <- brasilApiResponse
}

func requestViaCep(cep string, channel chan<- ViaCepResponse) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	defer resp.Body.Close()

	var viaCepResponse ViaCepResponse
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &viaCepResponse)
	if err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		return
	}

	channel <- viaCepResponse
}
