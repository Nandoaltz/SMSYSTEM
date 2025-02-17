package controller

import (
	"WEBTCC/Config"
	"WEBTCC/src/model"
	"WEBTCC/src/request"
	"WEBTCC/src/respostas"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/gorilla/mux"
)


func RegistrarCorrida(w http.ResponseWriter, r *http.Request) {
	
	body, err := io.ReadAll(r.Body)
	if err != nil {
    respostas.ERRO(w, http.StatusBadRequest, fmt.Errorf("erro ao ler o corpo da requisição: %v", err))
    return
	}

	fmt.Println("JSON recebido:", string(body))
	
	var data model.Registro

	if erro := json.Unmarshal(body, &data); erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	vars := mux.Vars(r)
	idVeiculo := vars["ID"]
	
	url := fmt.Sprintf("%s/veiculo/%s/registro", config.URLdaAPI, idVeiculo)

	if data.KM == "" || data.Tipo == "" {
		respostas.ERRO(w, http.StatusBadRequest, fmt.Errorf("Campos 'km' e 'tipoRegistro' são obrigatórios"))
		return
	}

	dataToSend := map[string]interface{}{
		"km": data.KM ,
		"tipo" : data.Tipo,
		"QuebraDeQuilometragem": data.QuebraDeQuilometragem,
	}
	
	d, erro := json.Marshal(dataToSend)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	
	
	// Faz a requisição POST para a API
	response, erro := request.RequestAut(r, http.MethodPost, url, bytes.NewBuffer(d))
	if erro != nil {
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	defer response.Body.Close()

	// Verifica o status da resposta da API
	if response.StatusCode >= 400 {
		respostas.ERROapi(w, response)
		return
	}

	// Se tudo deu certo, envia uma resposta de sucesso
	respostas.JSON(w, http.StatusCreated, map[string]string{"mensagem": "Registro efetuado"})
}