package controller

import (
	aut "TCC/src/Auth"
	"TCC/src/banco"
	"TCC/src/model"
	"TCC/src/repositorio"
	"TCC/src/respostas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


func CadastrarRegistro(w http.ResponseWriter, r *http.Request) {
    // Obtém o ID do veículo a partir da URL
    veiculo := mux.Vars(r)
    IDVeiculo, erro := strconv.ParseUint(veiculo["ID"], 10, 64)
    if erro != nil {
        respostas.ERRO(w, http.StatusBadRequest, erro)
        return
    }

    // Extrai o ID do usuário do token de autenticação
    usuarioID, erro := aut.ExtrairIDDoToken(r)
    if erro != nil {
        respostas.ERRO(w, http.StatusBadRequest, erro)
        return
    }

    // Lê o corpo da requisição
    corpo, erro := io.ReadAll(r.Body)
    if erro != nil {
        respostas.ERRO(w, http.StatusBadRequest, erro)
        return
    }
    fmt.Println(string(corpo))

    // Mapeia os dados do corpo para a estrutura de registro
    var publi model.Registro
    if erro := json.Unmarshal(corpo, &publi); erro != nil {
        respostas.ERRO(w, http.StatusBadRequest, erro)
        return
    }

    // Validações básicas
    if publi.KM == "" || publi.Tipo == "" {
        respostas.ERRO(w, http.StatusBadRequest, fmt.Errorf("os campos km e tipoRegistro são obrigatórios"))
        return
    }

    // Atribui os IDs ao registro
    publi.UsuarioID = usuarioID
    publi.VeiculoID = IDVeiculo

    // Conecta ao banco de dados
    db, erro := banco.ConnectDB()
    if erro != nil {
        respostas.ERRO(w, http.StatusInternalServerError, erro)
        return
    }
    defer db.Close()

    // Inicializa o repositório e tenta inserir o registro no banco
    repositorio := repositorio.Repositorio(db)

    publi.ID, erro = repositorio.PostarRegistro(publi.VeiculoID,publi.QuebraDeQuilometragem,publi)
    if erro != nil {
        respostas.ERRO(w, http.StatusInternalServerError, erro)
        publi.ID = 0
        return
    }

    respostas.JSON(w, http.StatusAccepted, publi)
}

func BuscarRegistros(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)
 
    IDVeiculo, erro := strconv.ParseUint(id["ID"], 10, 64)
    if erro != nil{
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
     registros, erro := repositorio.BuscarRegistros(IDVeiculo)

    if erro != nil{
    respostas.ERRO(w, http.StatusConflict, erro)
    return
    }
    
     respostas.JSON(w, http.StatusAccepted, registros)
}