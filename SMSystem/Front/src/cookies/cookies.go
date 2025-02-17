package cookies

import (
	config "WEBTCC/Config"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

var cookie *securecookie.SecureCookie

func COOKIE(){
	cookie = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, id string, token string) error{
	dados := make(map[string]string)
	dados["Id"] = id
	dados["Token"] = token

	dadosDoCookie, erro := cookie.Encode("cookie", dados)
	if erro != nil{
		return erro
	}
	http.SetCookie(w, &http.Cookie{Name:"cookie", Value: dadosDoCookie, Path: "/", 
	HttpOnly: true,
})
return nil
}
func Read(r *http.Request)(map[string]string, error){
		c, erro := r.Cookie("cookie")
		if erro != nil{
			return nil, erro
		}
		valores := make(map[string]string)

		if erro := cookie.Decode("cookie", c.Value ,&valores); erro != nil{
			return nil, erro
		}
		return valores, nil
}
func PegarIDUsuario(r *http.Request) (string, error) {
	// Lê os dados do cookie
	dados, erro := Read(r)
	if erro != nil {
		return "", erro // Retorna erro se não conseguir ler o cookie
	}

	// Obtém o ID do usuário do mapa retornado pela função Read
	idUsuario, existe := dados["Id"]
	if !existe {
		return "", fmt.Errorf("ID do usuário não encontrado no cookie")
	}

	// Retorna o ID do usuário
	return idUsuario, nil
}
func DeletarCookie(w http.ResponseWriter){
	http.SetCookie(w, &http.Cookie{
		Name: "cookie",
		Value: "",
		Path: "/",
		HttpOnly: true,
		Expires: time.Unix(0,0),
	})
}