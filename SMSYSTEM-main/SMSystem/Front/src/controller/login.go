package controller

import (
	config "WEBTCC/Config"
	"WEBTCC/src/cookies"
	"WEBTCC/src/model"
	"WEBTCC/src/respostas"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func FazerLogin(w http.ResponseWriter, r *http.Request){
	if erro := r.ParseForm(); erro != nil{
        respostas.ERRO(w, http.StatusBadRequest, erro)
        return
    }

    user, erro := json.Marshal(map[string]string{
        "email" : r.FormValue("email"),
        "senha" : r.FormValue("senha"),
    })
    if erro != nil{
        respostas.ERRO(w, http.StatusInternalServerError, erro)
        return
    }
    
    url := fmt.Sprintf("%s/login", config.URLdaAPI)
    resposta, erro := http.Post(url, "application/json", bytes.NewBuffer(user))
    if erro != nil{
        respostas.ERRO(w, http.StatusInternalServerError, erro)
        return
    }

    if resposta.StatusCode >= 400{
        respostas.ERROapi(w, resposta)
        return
    }
    defer resposta.Body.Close()

    var dados model.Dados

    if erro := json.NewDecoder(resposta.Body).Decode(&dados); erro != nil{
        respostas.ERRO(w, http.StatusBadRequest, erro)
        return
    }
    if erro = cookies.Save(w, dados.Id, dados.Token); erro != nil{
        respostas.ERRO(w, http.StatusUnauthorized, erro)
        return
    }
}
