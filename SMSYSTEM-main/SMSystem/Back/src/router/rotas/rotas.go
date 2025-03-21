package rotas

import (
	me "TCC/src/meedleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Rotas struct{
	URI string
	Method string
	Func func(w http.ResponseWriter, r *http.Request)
	Authentic bool
}

func ConfigurarRotas(r *mux.Router)*mux.Router{
		r = mux.NewRouter()
		for _, rout := range rotas{
			if rout.Authentic == true{
				r.HandleFunc(rout.URI, me.Logger(me.Autenticar(rout.Func))).Methods(rout.Method)
			}else{
				r.HandleFunc(rout.URI, me.Logger(rout.Func)).Methods(rout.Method)
			}
		}
		return r
}