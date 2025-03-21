package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statuscode int, info interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	if erro := json.NewEncoder(w).Encode(info); erro != nil{
		log.Fatal(erro)
	}

}

func ERRO(w http.ResponseWriter, statuscode int, erro error){
	JSON(w, statuscode, struct{
		Err string `json:"erro:, omitempty"`
	}{
		Err: erro.Error(),
	})
}