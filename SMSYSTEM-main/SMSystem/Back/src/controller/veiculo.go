package controller

import (
	aut "TCC/src/Auth"

	"TCC/src/banco"
	"TCC/src/model"
	"TCC/src/repositorio"
	"TCC/src/respostas"
	"encoding/json"

	"fmt"

	"io"
	"net/http"
	"strings"
)

func CadastrarVeiculo(w http.ResponseWriter, r *http.Request) {
    corpo, erro :=  io.ReadAll(r.Body)

	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	var veiculo model.Veiculo
   
	if erro = json.Unmarshal(corpo, &veiculo); erro != nil{
		respostas.ERRO(w, http.StatusNotFound, erro)
		return
	}


    db, err := banco.ConnectDB()
    if err != nil {
        respostas.ERRO(w, http.StatusInternalServerError, err)
        return
    }
    defer db.Close()

    repo := repositorio.Repositorio(db)
	idUser, _ := aut.ExtrairIDDoToken(r)

	fmt.Println("ID no token", idUser)

	tipo, erro := repo.TipoDeMotorista(idUser);
	fmt.Println(tipo)
	if erro != nil{
		respostas.ERRO(w, http.StatusInternalServerError,  erro)
		return
	}
	if tipo != "gestor"{
		respostas.ERRO(w, http.StatusUnauthorized, erro)
		return
	}


    if id, err := repo.CadastrarVeiculo(veiculo); err != nil {
       veiculo.ID = id
		respostas.ERRO(w, http.StatusInternalServerError, err)
        return
    }

    respostas.JSON(w, http.StatusCreated, veiculo)
}
func BuscarVeiculoNome(w http.ResponseWriter, r *http.Request) {
    Veiculo := strings.ToLower(r.URL.Query().Get("veiculos"))

	db, erro := banco.ConnectDB()
	if erro != nil{
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositorio.Repositorio(db)
	vcl, erro := repo.BuscarVeiculo(Veiculo)

	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	respostas.JSON(w, http.StatusAccepted, vcl)
}