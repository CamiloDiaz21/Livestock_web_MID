// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
<<<<<<< HEAD
		beego.NSNamespace("/usu_comun",
			beego.NSInclude(
				&controllers.Registro_usucomunController{},
			),
		),
		beego.NSNamespace("/usu_vendedor",
			beego.NSInclude(
				&controllers.Registro_usvendedorController{},
=======
		beego.NSNamespace("/ganado",
			beego.NSInclude(
				&controllers.GanadoController{},
>>>>>>> 911f526bea3dc169ba2b364fedb1f598aa457833
			),
		),
	)
	beego.AddNamespace(ns)
}
