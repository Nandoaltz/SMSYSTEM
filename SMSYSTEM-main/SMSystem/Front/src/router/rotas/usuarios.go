package rotas

import (
	"WEBTCC/src/controller"
	"net/http"
)

var RotasUsuarios = []Rotas{
	{
		URI: "/cadastrar",
		Method: http.MethodGet,
		Func: controller.TelaDeCadastro,
		Authentic: false,
	},
	{
		URI: "/usuarios",
		Method: http.MethodPost,
		Func: controller.CriarUmUsuario,
		Authentic: false,
	},
	{
		URI: "/home",
		Method: http.MethodGet,
		Func: controller.Home,
		Authentic: true,
	},
	{
		URI: "/veiculos/{ID}/registros",
		Method: http.MethodGet,
		Func: controller.PagRegistros,
		Authentic: true,
	},
	{
		URI: "/veiculo/{ID}/registro",
		Method: http.MethodPost,
		Func: controller.RegistrarCorrida,
		Authentic: true,
	},
	{
		URI: "/veiculos",
		Method: http.MethodGet,
		Func: controller.TelaCadastroVeiculo,
		Authentic: true,
	},
	{
		URI: "/veiculoCadastrar",
		Method: http.MethodPost,
		Func: controller.AdicionarVeiculo,
		Authentic: true,
	},
	{
		URI: "/buscarUsuarios",
		Method: http.MethodGet,
		Func: controller.MostrarUsuarios,
		Authentic: true,
	},
	{
		URI: "/usuarios/{ID}",
		Method: http.MethodGet,
		Func: controller.MostrarPerfilDeUsuario,
		Authentic: true,
	},
	{
		URI: "/perfilUsuarioLogado",
		Method: http.MethodGet,
		Func: controller.MostrarPerfilDeUsuarioLogado,
		Authentic: true,
	},
	{
		URI: "/EditarPerfil",
		Method: http.MethodGet,
		Func: controller.TelaAlterarInformacoesUser,
		Authentic: true,
	},
	{
		URI: "/SalvarAlteracao",
		Method: http.MethodPut,
		Func: controller.SalvarAlteracao,
		Authentic: true,
	},
	{
		URI: "/AtualizarSenha",
		Method: http.MethodGet,
		Func: controller.AtualizarSenha,
		Authentic: true,
	},
	{
		URI: "/AlterarSenha",
		Method: http.MethodPost,
		Func: controller.AlterarSenha,
		Authentic: true,
	},
	{
		URI: "/logout",
		Method: http.MethodGet,
		Func: controller.Logout,
		Authentic: true,
	},
	{
		URI: "/deletarUsuario",
		Method: http.MethodDelete,
		Func: controller.DeletarUser,
		Authentic: true,
	},
}
