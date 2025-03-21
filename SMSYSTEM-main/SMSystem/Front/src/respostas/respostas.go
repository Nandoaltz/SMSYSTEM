package respostas

import (
	"encoding/json"
	"log"
	"net/http"
	
)

	type Erro struct{
		erro string	`json:"erro"`
	}

func JSON(w http.ResponseWriter, statuscode int, dados interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	if erro := json.NewEncoder(w).Encode(dados); erro != nil{
		log.Fatal(erro)
	}
}
/*Ele pega o valor do struct Erro, decodifica para depois ser codificado na funcção JSON*/
func ERROapi(w http.ResponseWriter, r *http.Response){
	var erro Erro //Tipo Erro é a descrição do erro em string
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}

func ERRO(w http.ResponseWriter, statuscode int, erro error){
	JSON(w, statuscode, struct{
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}