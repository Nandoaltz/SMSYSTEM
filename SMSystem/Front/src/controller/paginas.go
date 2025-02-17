package controller

import (
	config "WEBTCC/Config"
	"WEBTCC/src/cookies"
	"WEBTCC/src/model"
	"WEBTCC/src/request"
	"WEBTCC/src/respostas"
	"WEBTCC/utils"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func TelaDeLogin(w http.ResponseWriter, r *http.Request){
	utils.Exec(w, "login.html", nil)
}
func TelaDeCadastro(w http.ResponseWriter, r *http.Request){
	utils.Exec(w, "cadastrar.html", nil)
}
func TelaCadastroVeiculo(w http.ResponseWriter, r *http.Request){
	utils.Exec(w, "Veiculos.html", nil)
}
func TelaDeRegistro(w http.ResponseWriter, r *http.Request){
	utils.Exec(w, "registros.html", nil)
}
func MostrarUsuarios(w http.ResponseWriter, r *http.Request){
	usuario := strings.ToLower(r.URL.Query().Get("usuario"))
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.URLdaAPI, usuario)

	response, erro := request.RequestAut(r, http.MethodGet, url, nil)
	fmt.Println()
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	var usuarios []model.Usuarios
	
	if erro := json.NewDecoder(response.Body).Decode(&usuarios); erro != nil{
		respostas.ERRO(w, http.StatusConflict, erro)
		return
	}

    // Renderiza o template com a lista de usuários
    utils.Exec(w, "usuarios.html", usuarios)
}
func MostrarPerfilDeUsuario(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	ID, erro := strconv.ParseUint(vars["ID"], 10, 64)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	fmt.Println("ID do usuario fora do cookie:", ID)
	IDCOOKIE, _ := cookies.PegarIDUsuario(r)
	IDcookie, _ := strconv.ParseUint(IDCOOKIE, 10, 64)
	if ID == IDcookie{
		http.Redirect(w, r, "/perfilUsuarioLogado", 302)
		return
	}
	url := fmt.Sprintf("%s/usuarios/%d", config.URLdaAPI, ID)
	response, erro := request.RequestAut(r, http.MethodGet, url, nil)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	var usuarios model.UsuarioComVeiculo

	if erro := json.NewDecoder(response.Body).Decode(&usuarios); erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	
	fmt.Println(usuarios.Usuario)
	fmt.Println(usuarios.Veiculo)

	utils.Exec(w, "perfil.html", usuarios)
}
func MostrarPerfilDeUsuarioLogado(w http.ResponseWriter, r *http.Request){
	IDdoCookie, erro := cookies.PegarIDUsuario(r)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	ID, erro := strconv.ParseUint(IDdoCookie, 10, 64)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	fmt.Println("Usuário logado", ID)

	url := fmt.Sprintf("%s/usuarios/%d", config.URLdaAPI, ID)
	response, erro := request.RequestAut(r, http.MethodGet, url, nil)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	var usuarios model.UsuarioComVeiculo

	if erro := json.NewDecoder(response.Body).Decode(&usuarios); erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	

	utils.Exec(w, "perfilDoUsuarioLogado.html", usuarios)
}
func TelaAlterarInformacoesUser(w http.ResponseWriter, r *http.Request){
	utils.Exec(w, "EditarPerfil.html", nil)
}
func AtualizarSenha(w http.ResponseWriter, r *http.Request){
	utils.Exec(w, "AtualizarSenha.html", nil)
}