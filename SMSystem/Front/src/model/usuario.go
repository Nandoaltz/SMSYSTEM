package model

import (
	
	"time"
)

type Usuarios struct{
	ID uint64		`json:"id,omitempty"`
	NOME string	`json:"nome,omitempty"`
	EMAIL string `json:"email,omitempty"`
	SENHA string `json:"senha,omitempty"`
<<<<<<< HEAD
	DATA time.Time `json:"data,omitempty"`
	VEICULO Veiculo `json:"veiculo,omitempty"`
=======
	TIPO string	`json:"tipos,omitempty"`
	DATA time.Time `json:"data,omitempty"`
>>>>>>> b308f24 (Novas Funcionalidades)
}