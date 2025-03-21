package controller

import (
	config "WEBTCC/Config"
	"WEBTCC/src/cookies"
	"WEBTCC/src/request"
	"WEBTCC/src/respostas"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func CriarUmUsuario(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		respostas.ERRO(w, http.StatusBadRequest, err)
		return
	}

	user, erro := json.Marshal(map[string]string{
		"nome": r.FormValue("nome"),
		"email" : r.FormValue("email"),
		"senha" : r.FormValue("senha"),
		"tipos": r.FormValue("tipos"),

	})


	fmt.Println(bytes.NewBuffer(user))

	
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	//O tatamento de erro é alternativo, pois não expecificamente o erro vai dar no trecho desse codigo
	//A requisição dessa pagina com a API pode retornrar um status code de erro
	//A requisição foi bem sucedida, mas o erro pode ocorrer no processo interno e a variavel de erro desse codigo ainda retorna o erro como nil
	url := fmt.Sprintf("%s/usuarios", config.URLdaAPI)
	resposta, erro := http.Post(url, "application/json", bytes.NewBuffer(user))
	if erro != nil{
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}

	//Com a resposta retornada da requisição da API, aqui é feito o tratamento do numero do status code
	//Se for maior de 400 isso indica um status de erro
	if resposta.StatusCode >= 400{
		respostas.ERROapi(w, resposta)
		return
	}
	
	defer resposta.Body.Close()
	respostas.JSON(w, http.StatusCreated, map[string]string{"mensagem": "Usuário criado"})
}

func SalvarAlteracao(w http.ResponseWriter, r *http.Request){
	erro := r.ParseForm()
	if erro != nil {
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	user, _ := json.Marshal(map[string]string{
		"nome": r.FormValue("nome"),
		"email" : r.FormValue("email"),
	})
	fmt.Println(bytes.NewBuffer(user))
	Id, _ := cookies.PegarIDUsuario(r)
	ID, _ := strconv.ParseUint(Id, 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d", config.URLdaAPI, ID)
	fmt.Println(url)

	resposta, erro := request.RequestAut(r, http.MethodPut, url, bytes.NewBuffer(user))
	if erro != nil{
		respostas.ERRO(w, http.StatusConflict, erro)
		return
	}

	//Com a resposta retornada da requisição da API, aqui é feito o tratamento do numero do status code
	//Se for maior de 400 isso indica um status de erro
	if resposta.StatusCode >= 400{
		respostas.ERROapi(w, resposta)
		return
	}
	
	defer resposta.Body.Close()

	respostas.JSON(w, http.StatusAccepted, map[string]string{"mensagem": "Usuário Atualizado"})
}

func AlterarSenha(w http.ResponseWriter, r *http.Request){
	erro := r.ParseForm()
	if erro != nil {
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	user, _ := json.Marshal(map[string]string{
		"senha": r.FormValue("senha"),
		"senhanova" : r.FormValue("senhanova"),
	})

	Id, _ := cookies.PegarIDUsuario(r)
	ID, _ := strconv.ParseUint(Id, 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d/atualizarsenha", config.URLdaAPI, ID)
	fmt.Println(url)

	resposta, erro := request.RequestAut(r, http.MethodPost, url, bytes.NewBuffer(user))
	if erro != nil{
		respostas.ERRO(w, http.StatusConflict, erro)
		return
	}

	//Com a resposta retornada da requisição da API, aqui é feito o tratamento do numero do status code
	//Se for maior de 400 isso indica um status de erro
	if resposta.StatusCode >= 400{
		respostas.ERROapi(w, resposta)
		return
	}
	
	defer resposta.Body.Close()

	respostas.JSON(w, http.StatusAccepted, map[string]string{"mensagem": "Usuário Atualizado"})
}
func Logout(w http.ResponseWriter, r *http.Request){
	cookies.DeletarCookie(w)
	http.Redirect(w, r, "/login", 302)
}
func DeletarUser(w http.ResponseWriter, r *http.Request){
	cookieID, _ := cookies.PegarIDUsuario(r)
	ID, _ := strconv.ParseUint(cookieID, 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d", config.URLdaAPI, ID)

	resposta, erro := request.RequestAut(r, http.MethodDelete, url, nil)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	if resposta.StatusCode >= 400{
		respostas.ERROapi(w, resposta)
		return
	}
	
	defer resposta.Body.Close()
	respostas.JSON(w, http.StatusAccepted,nil)
}