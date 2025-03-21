package router

import (
	"WEBTCC/src/router/rotas"

	"github.com/gorilla/mux"
)

func GerarNovoRouter() *mux.Router{
	return rotas.ConfigRotas(mux.NewRouter())
}