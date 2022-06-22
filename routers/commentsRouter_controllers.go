package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["aceniubi/controllers:JwtController"] = append(beego.GlobalControllerRouter["aceniubi/controllers:JwtController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/api/login/[post]",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
