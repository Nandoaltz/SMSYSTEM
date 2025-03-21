package rotas

import (
	"WEBTCC/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Rotas struct{
	URI string
	Method string
	Func func(w http.ResponseWriter, r *http.Request)
	Authentic bool
}

func ConfigRotas(r *mux.Router) *mux.Router{
		rotas := RotasSliceLogin
		rotas = append(rotas, RotasUsuarios...)
		

		for _, router := range rotas{
			if router.Authentic{
				r.HandleFunc(router.URI, middlewares.Logger(middlewares.Autenticar(router.Func))).Methods(router.Method)
			}else{
				r.HandleFunc(router.URI, middlewares.Logger(router.Func)).Methods(router.Method)
			}
		}

		fileServer := http.FileServer(http.Dir("./assets/"))
		r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

		return r
}

