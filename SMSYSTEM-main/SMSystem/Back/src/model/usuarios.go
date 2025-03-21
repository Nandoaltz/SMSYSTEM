package model

import (
	hashsenha "TCC/src/HashSenha"
	"errors"
	"regexp"
	"strings"
	"time"
)

type Usuarios struct{
	ID uint64		`json:"id,omitempty"`
	NOME string	`json:"nome,omitempty"`
	EMAIL string `json:"email,omitempty"`
	SENHA string `json:"senha,omitempty"`

	TIPO string `json:"tipos,omitempty"`

	DATA time.Time `json:"data,omitempty"`
}

func (u *Usuarios) Format(etapa string) error{
	
    valido, err := validarEmail(u.EMAIL)
	if err != nil {
		return err
	}
	if !valido {
		return errors.New("Email inválido")
	}
	
	if erro := u.valid(etapa); erro != nil{
		return erro
	}
	return nil
}
func (u *Usuarios) valid(etapa string)error{
	if u.NOME == ""{
		return errors.New("Preencha o campo Nome")
	}
	if etapa == "cadastro"{
		Hash, erro := hashsenha.ConvertHash(u.SENHA)
		if erro != nil{
			return erro
		}
		u.SENHA = string(Hash)
	}

	if etapa == "cadastro" && u.SENHA == ""{
		return errors.New("Preencha o campo Senha")
	}
	
	u.NOME = strings.TrimSpace(u.NOME)
	u.EMAIL = strings.TrimSpace(u.EMAIL)


	return nil
}

func validarEmail(email string) (bool, error) {
	
    regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    re, err := regexp.Compile(regex)
    if err != nil {
        return false, err
    }
	if email == ""{
		return false, errors.New("Campo obrigatório ")
	}
    return re.MatchString(email), nil
}
