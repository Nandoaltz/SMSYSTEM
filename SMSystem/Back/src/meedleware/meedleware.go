package me

import (
	aut "TCC/src/Auth"
	"TCC/src/respostas"
	"fmt"
	"net/http"
)


func Logger(n http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("\n %s %s %s \n", r.URL,r.Method, r.Host)
		n(w,r)
	}
}

func Autenticar(n http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if erro := aut.ValidToken(r); erro != nil{
			respostas.ERRO(w,http.StatusUnauthorized, erro)
			return
		}
		n(w,r)
	}
}