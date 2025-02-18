package router

import (
	"TCC/src/router/rotas"

	"github.com/gorilla/mux"
)

func GererNovoRouter()*mux.Router{
	return rotas.ConfigurarRotas(mux.NewRouter())
}