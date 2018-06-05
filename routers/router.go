package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"webproject/controllers"
	"net/http"
	"html/template"
)

// 自定义错误页面跳转页面
func pageNotFound(w http.ResponseWriter, r *http.Request){
	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "Sorry Page Not Found."
	data["code"] = "404"
	t.Execute(w, data)
}

func init() {
	beego.ErrorHandler("404", pageNotFound)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "get:Login")
	beego.Router("/:code:string", &controllers.MainController{}, "*:ErrorPage")
	//beego.Router("/admin",&controllers.AdminController{})
	//beego.Router("/admin/listbook", &controllers.AdminController{}, "get:ListBook")
	//beego.Router("/admin/form", &controllers.AdminController{}, "get:FormPage")
	//beego.Router("/admin/conf", &controllers.AdminController{}, "get:AdminConf")
	//beego.Router("/admin/user/:username:string", &controllers.AdminController{}, "*:AdminRouter")
	//beego.Router("/admin/api/:id:int", &controllers.AdminController{}, "*:AdminRouter")
	//beego.Router("/admin/download/*.*", &controllers.AdminController{}, "*:AdminRouter")
	//beego.Router("/admin/download/test/*", &controllers.AdminController{}, "get,post:AdminRouter")

	// 自动路由配置
	//beego.AutoRouter(&controllers.AutoRouterController{})
	// 使用注解配置
	//beego.Include(&controllers.AnnoRouterController{})

	// namespace 命令空间写法
	ns := beego.NewNamespace("/api",
		beego.NSCond(func(ctx *context.Context) bool {
			domain := ctx.Input.Domain()
			if domain == "127.0.0.1" || domain == "localhost" {
				return true
			}
			return false
		}),
		beego.NSNamespace("/v1",
			beego.NSGet("/", func(ctx *context.Context) {
				result := map[string]interface{}{"success": false, "message": "Not Allowed."}
				ctx.Output.JSON(result, true, false)
			}),
			beego.NSNamespace("/admin",
				beego.NSRouter("/", &controllers.AdminController{}),
				beego.NSRouter("/list", &controllers.AdminController{}, "get:ListBook"),
				beego.NSRouter("/form", &controllers.AdminController{}, "get,post:FormPage"),
				beego.NSRouter("/conf", &controllers.AdminController{}, "*:AdminConf"),
			),
			beego.NSNamespace("/user",
				beego.NSGet("/", func(ctx *context.Context) {
					ctx.Output.Body([]byte("notAllowed."))
				}),
				beego.NSRouter("/index", &controllers.UserController{}),
				beego.NSRouter("/save", &controllers.UserController{}),
				beego.NSRouter("/saveuser", &controllers.UserController{}, "post:Save"),
				beego.NSRouter("/:username:string", &controllers.AdminController{}, "*:AdminRouter"),
				beego.NSRouter("/id/:id:int", &controllers.AdminController{}, "*:AdminRouter"),
				beego.NSNamespace("/download",
					beego.NSRouter("/*.*", &controllers.AdminController{}, "*:AdminRouter"),
					beego.NSRouter("/test/*", &controllers.AdminController{}, "get,post:AdminRouter"),
				),
				beego.NSRouter("/json", &controllers.UserController{}, "*:JsonPrint"),
				beego.NSRouter("/xml", &controllers.UserController{}, "*:XmlPrint"),
				beego.NSRouter("/jsonp", &controllers.UserController{}, "*:JsonpPrint"),
				beego.NSRouter("/header", &controllers.UserController{}, "*:FormatPrint"),
			),
			beego.NSNamespace("/flash",
				beego.NSRouter("/post", &controllers.FlashController{}),
				beego.NSRouter("/get", &controllers.FlashController{}),
			),
			beego.NSInclude(&controllers.AnnoRouterController{}),
			beego.NSAutoRouter(&controllers.AutoRouterController{}),
			beego.NSAutoRouter(&controllers.PrepareController{}),
		),
	)
	beego.AddNamespace(ns)

	// 自定义 filter 过滤
	var filterLogin = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("uid").(int)
		if !ok && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/login")
		}
	}
	beego.InsertFilter("/api/v1/admin/list", beego.BeforeRouter, filterLogin)

}
