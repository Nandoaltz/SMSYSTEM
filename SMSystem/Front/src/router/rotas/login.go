package rotas

import (
	"WEBTCC/src/controller"
	"net/http"
)

var RotasSliceLogin = []Rotas{
	{
		URI: "/login",
		Method: http.MethodGet,
		Func: controller.TelaDeLogin,
		Authentic: false,
	},
	{
		URI: "/login",
		Method: http.MethodPost,
		Func: controller.FazerLogin,
		Authentic: false,
	},
}