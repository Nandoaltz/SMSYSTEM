package model

import "time"

type Registro struct{
	ID          uint64 `json:"id,omitempty"`
    UsuarioID  uint64 `json:"usuario_id,omitempty"` // ID do usuário que cadastrou o veículo
    UsuarioNome string	`json:"usuario_nome,omitempty"`
	VeiculoID	uint64	`json:"veiculo_id,omitempty"`
	VeiculoNome string	`json:"Veiculo_nome,omitempty"`
	Horario_Data time.Time  `json:"horario_data,omitempty"`
   	KM string `json:"km,omitempty"`
	Tipo string `json:"tipo,omitempty"`
	QuebraDeQuilometragem bool `json:"QuebraDeQuilometragem,omitempty"`
}