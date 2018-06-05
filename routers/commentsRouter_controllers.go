package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["webproject/controllers:AnnoRouterController"] = append(beego.GlobalControllerRouter["webproject/controllers:AnnoRouterController"],
		beego.ControllerComments{
			Method:           "ListAll",
			Router:           `/router/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["webproject/controllers:AnnoRouterController"] = append(beego.GlobalControllerRouter["webproject/controllers:AnnoRouterController"],
		beego.ControllerComments{
			Method:           "SaveInfo",
			Router:           `/router/save/*.*`,
			AllowHTTPMethods: []string{"*"},
			MethodParams:     param.Make(),
			Params:           nil})

}
