package config
import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var(
    URLdaAPI string
    Porta int
    HashKey []byte
    BlockKey []byte
    erro error
)

func CarregarVariaveisDeAmbiente(){
    
    if erro = godotenv.Load(); erro != nil{
        log.Fatal(erro)
    }
    
    URLdaAPI = os.Getenv("URLAPI")

    Porta, erro = strconv.Atoi(os.Getenv("Porta"))
    if erro != nil{
        log.Fatal(erro)
    }

    HashKey = []byte(os.Getenv("HASHKEY"))
    BlockKey = []byte(os.Getenv("BLOCKKEY"))
}