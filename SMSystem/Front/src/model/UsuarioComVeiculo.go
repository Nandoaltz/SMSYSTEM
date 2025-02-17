package model

type UsuarioComVeiculo struct {
	Usuario Usuarios `json:"usuario"`
	Veiculo Veiculo `json:"ultimoVeiculo"`
}