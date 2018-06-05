package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

//func (c *ErrorController) Error404() {
//	c.Data["code"] = "404"
//	c.Data["content"] = "Page Not Found"
//	c.TplName = "error.html"
//}

func (c *ErrorController) Error500() {
	c.Data["code"] = "500"
	c.Data["content"] = "Internal Server Error"
	c.TplName = "error.html"
}

func (c *ErrorController) Errordb() {
	c.Data["code"] = "DBError"
	c.Data["content"] = "Database Server Unavailable"
	c.TplName = "error.html"
}

func (c *ErrorController) Errorcustom() {
	c.Data["code"] = "CustomError"
	c.Data["content"] = "用户自定义错误信息。"
	c.TplName = "error.html"
}

func (c *ErrorController) Errorunknow() {
c.Data["code"] = "UnknowError"
c.Data["content"] = "系统未知异常，请与系统管理员联系。"
c.TplName = "error.html"
}

