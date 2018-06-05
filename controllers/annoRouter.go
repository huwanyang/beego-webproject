package controllers

import "github.com/astaxie/beego"

type AnnoRouterController struct {
	beego.Controller
}

func (c *AnnoRouterController) URLMapping() {
	c.Mapping("ListAll", c.ListAll)
	c.Mapping("SaveInfo", c.SaveInfo)
}

// @router /router/list [get]
func (c *AnnoRouterController) ListAll() {
	// simple 类型
	c.Data["Title"] = "Beego Router List"
	c.Data["Content"] = "This is Beego Annotation Router list page."
	// slice 类型
	var bookList = make([]string, 0)
	bookList = append(bookList, "Golang Book")
	bookList = append(bookList, "Beego Book")
	bookList = append(bookList, "AutoTesting Book")
	c.Data["BookList"] = bookList
	c.TplName = "router/list.html"
}

// @router /router/save/*.* [*]
func (c *AnnoRouterController) SaveInfo() {
	param := make(map[string]interface{})
	var username = c.Ctx.Input.Param(":username")
	if len(username) != 0 {
		param["username"] = username
	}
	var id = c.Ctx.Input.Param(":id")
	if len(id) != 0 {
		param["id"] = id
	}
	var path = c.Ctx.Input.Param(":path")
	if len(path) != 0 {
		param["path"] = path
	}
	var ext = c.Ctx.Input.Param(":ext")
	if len(ext) != 0 {
		param["ext"] = ext
	}
	var splat = c.Ctx.Input.Param(":splat")
	if len(splat) != 0 {
		param["splat"] = splat
	}
	c.Data["Params"] = param
	c.TplName = "router/router.html"
}
