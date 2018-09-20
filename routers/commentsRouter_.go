package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["blog-api/controllers:ErrorController"] = append(beego.GlobalControllerRouter["blog-api/controllers:ErrorController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/error`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-api/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog-api/controllers:LoginController"],
		beego.ControllerComments{
			Method: "ToLogin",
			Router: `/comm/toLogin`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-api/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog-api/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/comm/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-api/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog-api/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-api/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog-api/controllers:LoginController"],
		beego.ControllerComments{
			Method: "ToReg",
			Router: `/comm/toReg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog-api/controllers:LoginController"] = append(beego.GlobalControllerRouter["blog-api/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Reg",
			Router: `/comm/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
