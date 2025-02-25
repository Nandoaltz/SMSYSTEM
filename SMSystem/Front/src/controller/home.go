package controller

import (
	config "WEBTCC/Config"
	"WEBTCC/src/cookies"
	"WEBTCC/src/model"
	"WEBTCC/src/request"
	"WEBTCC/src/respostas"
	"WEBTCC/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request){
	url := fmt.Sprintf("%s/veiculos", config.URLdaAPI)
	response, erro := request.RequestAut(r, http.MethodGet, url, nil)
	if erro != nil{
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400{
		respostas.ERROapi(w, response)
		return
	}
	var veiculos []model.Veiculo

	if erro := json.NewDecoder(response.Body).Decode(&veiculos); erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	utils.Exec(w, "home.html", veiculos)
}

func IdRequestPerfil(w http.ResponseWriter, r *http.Request){
	url := fmt.Sprintf("%s/veiculos", config.URLdaAPI)
	response, erro := request.RequestAut(r, http.MethodGet, url, nil)
	if erro != nil{
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400{
		respostas.ERROapi(w, response)
		return
	}
	var veiculos []model.Veiculo

	if erro := json.NewDecoder(response.Body).Decode(&veiculos); erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var User model.Usuarios

	IDCookie, erro := cookies.PegarIDUsuario(r)
	if erro != nil{
		respostas.ERRO(w, http.StatusConflict, erro)
		return
	}

	User.ID, erro = strconv.ParseUint(IDCookie, 10, 64)


	utils.Exec(w, "home.html", User.ID)
}