package rotas

import (
	"TCC/src/controller"
	"net/http"
)

var rotas = []Rotas{
	{
		URI: "/usuarios",
		Method: http.MethodPost,
		Func: controller.UsuarioPOST,
		Authentic: false,
	},
	{
		URI: "/usuarios",
		Method: http.MethodGet,
		Func: controller.UsuarioGET,
		Authentic: true,
	},
	{
		URI: "/usuarios/{usrID}",
		Method: http.MethodGet,
		Func: controller.UsuariosGETID,
		Authentic: true,
	},
	{
		URI: "/usuarios/{usrID}",
		Method: http.MethodPut,
		Func: controller.UsuariosREAD,
		Authentic: true,
	},
	{
		URI: "/usuarios/{usrID}/atualizarsenha",
		Method: http.MethodPost,
		Func: controller.AtualizarSenha,
		Authentic: true,
	},
	{
		URI: "/usuarios/{usrID}",
		Method: http.MethodDelete,
		Func: controller.UsuariosDELETE,
		Authentic: true,
	},
	{
		URI: "/login",
		Method: http.MethodPost,
		Func: controller.Login,
		Authentic: false,
	},
	{
		URI: "/veiculoCadastrar",
		Method: http.MethodPost,
		Func: controller.CadastrarVeiculo,
		Authentic: true,
	},
	{
		URI: "/veiculos",
		Method: http.MethodGet,
		Func: controller.BuscarVeiculoNome,
		Authentic: true,
	},
	{
		URI: "/veiculo/{ID}/registro",
		Method: http.MethodPost,
		Func: controller.CadastrarRegistro,
		Authentic: true,
	},
	{
		URI: "/veiculos/{ID}/registros",
		Method: http.MethodGet,
		Func: controller.BuscarRegistros,
		Authentic: true,
	},
}