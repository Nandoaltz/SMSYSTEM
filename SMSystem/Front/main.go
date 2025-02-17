package main

import (
	config "WEBTCC/Config"
	"WEBTCC/src/cookies"
	"WEBTCC/src/router"
	"WEBTCC/utils"
	// "encoding/hex"
	"fmt"
	"log"
	"net/http"

	// "github.com/gorilla/securecookie"
)
// func init(){
// 	Randonkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(Randonkey)
// 	sla := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(sla)
// }
func main(){
	utils.CarregarTemplates()
	config.CarregarVariaveisDeAmbiente()
	r := router.GerarNovoRouter()
	cookies.COOKIE()
	porta := fmt.Sprintf(":%d", config.Porta)
	fmt.Printf("Rodando na porta: %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(porta, r))
}