package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Title"] = "Welcome to Beego. This is Title."
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Login() {
	c.Ctx.Output.Body([]byte("Please Login."))
}

func (c *MainController) ErrorPage() {
	var code = c.Ctx.Input.Param(":code")
	fmt.Println("code: ", code)
	switch code {
	case "401":
		c.Abort(code)
	case "403":
		c.Abort(code)
	case "404":
		c.Abort(code)
	case "500":
		c.Abort(code)
	case "503":
		c.Abort(code)
	case "db":
		c.Abort(code)
	case "custom":
		c.Abort(code)
	case "unknow":
		c.Abort(code)
	default:
		c.Ctx.Output.Body([]byte("This is " + code + " Page."))
	}
}
