package main

import (
	// aut "TCC/src/Auth"
	"TCC/src/config"
	"TCC/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	// secretKey, err := aut.GenerateSecretKey()
	// if err != nil {
	// 	fmt.Println("Erro ao gerar chave secreta:", err)
	// 	return
	// }
	// fmt.Println(secretKey)
	
	router := router.GererNovoRouter()
	config.Config()
	porta := fmt.Sprintf(":%d", config.Porta)
	fmt.Printf("Rodando na porta:%d", config.Porta)
	log.Fatal(http.ListenAndServe(porta, router))
}