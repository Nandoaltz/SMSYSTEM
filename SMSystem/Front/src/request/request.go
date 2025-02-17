package request

import (
	"WEBTCC/src/cookies"
	"fmt"
	"io"
	"net/http"
)
//Essa dunção adiciona o token no campo de autenticação do bearer token da requisição
func RequestAut(r *http.Request, metodo string, url string, dados io.Reader)(*http.Response, error){
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil{
		return nil, erro
	}
	cookie, _ := cookies.Read(r)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cookie["Token"]))

	client := &http.Client{}
	response, erro := client.Do(request)
	if erro != nil{
		return nil, erro
	}
	return response, nil
}