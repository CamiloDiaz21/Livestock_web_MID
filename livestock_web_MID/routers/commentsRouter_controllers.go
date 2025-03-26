package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:GanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:GanadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:GanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:GanadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:GanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:GanadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:GanadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:GanadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Hacienda_lotesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Hacienda_lotesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Hacienda_lotesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Hacienda_lotesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Hacienda_lotesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Hacienda_lotesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Hacienda_lotesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Hacienda_lotesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Hacienda_lotesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Hacienda_lotesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Historial_ventasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Historial_ventasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Historial_ventasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Historial_ventasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Historial_ventasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Historial_ventasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Historial_ventasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Historial_ventasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Historial_ventasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Historial_ventasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usucomunController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usucomunController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usucomunController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usucomunController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usucomunController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usucomunController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usucomunController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usucomunController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usucomunController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usucomunController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usvendedorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usvendedorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usvendedorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usvendedorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usvendedorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usvendedorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usvendedorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usvendedorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usvendedorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Registro_usvendedorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ganadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ganadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ganadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ganadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ganadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ganadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ganadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ganadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ganadoController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ganadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ventaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ventaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ventaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ventaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ventaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ventaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ventaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ventaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ventaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_web_MID/livestock_web_MID/controllers:Tipo_ventaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
