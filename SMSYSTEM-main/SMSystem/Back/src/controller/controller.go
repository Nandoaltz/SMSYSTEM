package controller

import (
	aut "TCC/src/Auth"
	hashsenha "TCC/src/HashSenha"
	"TCC/src/banco"
	"TCC/src/model"
	"TCC/src/repositorio"
	"TCC/src/respostas"
	"fmt"
	"strconv"
	"strings"

	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func UsuarioPOST(w http.ResponseWriter, r *http.Request){
	corpo, erro :=  io.ReadAll(r.Body)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	var user model.Usuarios
	if erro = json.Unmarshal(corpo, &user); erro != nil{
		respostas.ERRO(w, http.StatusNotFound, erro)
		return
	}
	if erro := user.Format("cadastro"); erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.ConnectDB()
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	
	defer db.Close()

	repositorio := repositorio.Repositorio(db)
 	userID , erro := repositorio.CriaUsuarios(user)
	user.ID = userID
	if erro != nil{
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusCreated, user)
}

func UsuarioGET(w http.ResponseWriter, r *http.Request){
	Usuario := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.ConnectDB()
	if erro != nil{
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()
	repo := repositorio.Repositorio(db)
	user, erro := repo.BuscarUsuario(Usuario)

	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	
	respostas.JSON(w, http.StatusAccepted, user)
}	
func UsuariosGETID(w http.ResponseWriter, r *http.Request){
	Usuarios := mux.Vars(r)
	
	ID, erro := strconv.ParseUint(Usuarios["usrID"], 10, 64)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.ConnectDB()
	if erro != nil{
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}
	repo := repositorio.Repositorio(db)
	usr, erro := repo.BuscarUsuarioID(ID)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	ultimoVeiculo, erro := repo.BuscarUltimoVeiculo(usr.ID)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	type UsuarioComVeiculo struct {
		Usuario model.Usuarios `json:"usuario"`
		Veiculo model.Veiculo `json:"ultimoVeiculo"`
	}

	usuarioComVeiculo := UsuarioComVeiculo{
		Usuario: usr,
		Veiculo: ultimoVeiculo,
	}

	fmt.Println(usuarioComVeiculo)

	respostas.JSON(w, http.StatusAccepted, usuarioComVeiculo)
}
func UsuariosREAD(w http.ResponseWriter, r *http.Request){
	Usuarios := mux.Vars(r)

	usrID, err := strconv.ParseUint(Usuarios["usrID"], 10, 64)
	if err != nil{
	respostas.ERRO(w, http.StatusBadRequest, err)
	return
	}

	UsuarioToken, erro := aut.ExtrairIDDoToken(r)
	if erro != nil{
		respostas.ERRO(w, http.StatusUnauthorized, erro)
		return
	}
	if usrID != UsuarioToken{
		respostas.ERRO(w, http.StatusForbidden, erro)
		return
	}

	Request, err := io.ReadAll(r.Body)
	if err != nil{
	respostas.ERRO(w, http.StatusUnprocessableEntity, err)
	return
}
	var Usuario model.Usuarios

	if erro := json.Unmarshal(Request, &Usuario); erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	// if erro := Usuario.Format("edit"); erro != nil{
	// 	respostas.ERRO(w, http.StatusConflict, erro)
	// 	return
	// }
	db, erro := banco.ConnectDB()
	if erro != nil{
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositorio.Repositorio(db)
	repo.UpdateUsers(usrID, Usuario)
	
	respostas.JSON(w, http.StatusOK, Usuario)
}
func UsuariosDELETE(w http.ResponseWriter, r *http.Request){
	Usuario := mux.Vars(r)

	ID, err := strconv.ParseUint(Usuario["usrID"], 10, 64)
	if err != nil{
		respostas.ERRO(w, http.StatusBadRequest, err)
		return
	}

	IDEstract, erro := aut.ExtrairIDDoToken(r)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	if ID != IDEstract{
		respostas.ERRO(w, http.StatusUnauthorized, erro)
		return
	}

	db, err := banco.ConnectDB()
	if err != nil{
		respostas.ERRO(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repositorio.Repositorio(db)
	repo.DeleteUsers(ID)

	respostas.JSON(w, http.StatusAccepted, nil)
}
func Login(w http.ResponseWriter, r *http.Request){
		Usuario, erro := io.ReadAll(r.Body)
		if erro !=  nil{
			respostas.ERRO(w, http.StatusBadRequest, erro)
			return
		}
		var usuario model.Usuarios
		if erro := json.Unmarshal(Usuario, &usuario); erro != nil{
			respostas.ERRO(w, http.StatusInternalServerError, erro)
			return
		}

		db, err := banco.ConnectDB()
		if err != nil{
			respostas.ERRO(w, http.StatusInternalServerError, err)
			return
		}
		defer db.Close()

		repo := repositorio.Repositorio(db)
		u, err := repo.LOGIN(usuario.EMAIL)
		if err != nil{
			respostas.ERRO(w, http.StatusInternalServerError, err)
			return
		}
		if erro := hashsenha.CompareHash(usuario.SENHA, u.SENHA); erro != nil{
			respostas.ERRO(w, http.StatusUnauthorized, erro)
			return
		}
		token, erro := aut.JsonWebToken(u.ID)
		if erro != nil{
			respostas.ERRO(w, http.StatusInternalServerError, erro)
			return
		}

		userID := strconv.FormatUint(u.ID, 10)
	
	/*Os dados contendo o id e o token do usuário serão enviados para o front para serem armazenados no cookie*/
	respostas.JSON(w, http.StatusOK, model.Dados{Id: userID, Token: token})
}
func AtualizarSenha(w http.ResponseWriter, r *http.Request){
	IdDoToken, erro := aut.ExtrairIDDoToken(r)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	param := mux.Vars(r)
	userId, erro := strconv.ParseUint(param["usrID"], 10, 64)

	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	if IdDoToken != userId{
		respostas.ERRO(w, http.StatusConflict, erro)
		return
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	var Senha model.Senhas

	if erro := json.Unmarshal(corpoRequest, &Senha); erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}
	
	fmt.Print(Senha.NovaSenha, Senha.SenhaAtual)
	db, erro := banco.ConnectDB()
	if erro != nil{
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorio.Repositorio(db)

	SenhaDoBanco, erro := repositorio.BuscarSenha(IdDoToken)
	if erro != nil{
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}
	if erro := hashsenha.CompareHash(Senha.SenhaAtual, SenhaDoBanco); erro != nil{
		respostas.ERRO(w, http.StatusConflict, erro)
		return
	}
	senhaNova, erro := hashsenha.ConvertHash(Senha.NovaSenha)
	if erro != nil{
		respostas.ERRO(w, http.StatusUnauthorized, erro)
		return
	}

	if erro := repositorio.InserirSenhaNova(IdDoToken, string(senhaNova)); erro != nil{
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusAccepted, nil)
}

