package controller

import (
	config "WEBTCC/Config"
	"WEBTCC/src/request"
	"WEBTCC/src/respostas"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func AdicionarVeiculo(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		respostas.ERRO(w, http.StatusBadRequest, err)
		return
	}
	veiculo, erro := json.Marshal(map[string]string{
		"nome":r.FormValue("nome"),
		"placa":r.FormValue("placa"),
	})
	if erro != nil{
		respostas.ERRO(w, http.StatusConflict, erro)
		return
	}
	url := fmt.Sprintf("%s/veiculoCadastrar", config.URLdaAPI)

	response, erro := request.RequestAut(r, http.MethodPost, url, bytes.NewBuffer(veiculo))
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	defer response.Body.Close()

if response.StatusCode >= 400 {
    respostas.ERROapi(w, response)
    return
}

respostas.JSON(w, http.StatusCreated, map[string]string{"mensagem": "Registro efetuado"})
}