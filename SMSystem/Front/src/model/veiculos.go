package model

import "time"

type Veiculo struct {
    ID          uint64 `json:"id,omitempty"`
    UsuarioID  uint64 `json:"usuario_id,omitempty"`
    Nome        string  `json:"nome,omitempty"`
    Placa       string `json:"placa,omitempty"`
    DATA         time.Time `json:"data,omitempty"`
}