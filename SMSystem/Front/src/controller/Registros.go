package controller

import (
	config "WEBTCC/Config"
	"WEBTCC/src/model"
	"WEBTCC/src/request"
	"WEBTCC/src/respostas"
	"WEBTCC/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func PagRegistros(w http.ResponseWriter, r *http.Request) {
	// Pega o ID da URL
	vars := mux.Vars(r)
	idVeiculo := vars["ID"]

	// Gera a URL da API para obter os registros
	url := fmt.Sprintf("%s/veiculos/%s/registros", config.URLdaAPI, idVeiculo)

	if r.Method == http.MethodPost {
		RegistrarCorrida(w, r)
	}
	// Faz a requisição
	response, erro := request.RequestAut(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	defer response.Body.Close()

	// Verifica o status da resposta
	if response.StatusCode >= 400 {
		respostas.ERROapi(w, response)
		return
	}

	// Cria o slice para armazenar os registros
	var registros []model.Registro

	// Decodifica o corpo da resposta JSON em registros
	if erro := json.NewDecoder(response.Body).Decode(&registros); erro != nil {
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	

	// Define o cabeçalho de conteúdo HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Renderiza o template passando os registros
	utils.Exec(w, "registros.html", registros)
}
