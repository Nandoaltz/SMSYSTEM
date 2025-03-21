package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)
var(
	Porta int
	erro error
	StringBanco string
	Signiture []byte
)
func Config(){
	
	if erro = godotenv.Load(); erro != nil{
		log.Fatal()
	}
	
	strport := os.Getenv("PORTA")
	Porta, erro = strconv.Atoi(strport)

	StringBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("USUARIO_DB"), os.Getenv("SENHA_DB"), os.Getenv("NOME_DB"))
	
	Signiture = []byte(os.Getenv("SECRET_KEY"))
}