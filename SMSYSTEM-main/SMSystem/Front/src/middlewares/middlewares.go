package middlewares

import (
	"WEBTCC/src/cookies"
	"fmt"
	"net/http"
)

func Logger(p http.HandlerFunc)http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		fmt.Printf("\n %s %s %s", r.URL, r.Method, r.Host)
		p(w,r)
	}
}
func Autenticar(p http.HandlerFunc)http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		_, erro := cookies.Read(r)
		//Se der erro na hora de ler o cookie, ele retorna para a tela de login
		if erro != nil{
			http.Redirect(w,r, "/login", 301)
		}
		p(w,r)
	}
}