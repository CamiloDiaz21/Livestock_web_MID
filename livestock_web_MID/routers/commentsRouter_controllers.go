package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

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

}
