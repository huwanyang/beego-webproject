package main

import (
	"github.com/astaxie/beego"
	_ "webproject/routers"
	"webproject/controllers"
)

func main() {
	// 设置显示该目录下所有的文件列表
	beego.BConfig.WebConfig.DirectoryIndex = true
	// 设置静态文件访问路径，也可以在 app.conf 中配置
	beego.SetStaticPath("/download", "download")
	// 设置 Controller 方式定义 Error 错误处理函数
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
