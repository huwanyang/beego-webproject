package controllers

import "github.com/astaxie/beego"

type AutoRouterController struct {
	beego.Controller
}

func (c *AutoRouterController) Login() {
	paramlen := c.Ctx.Input.ParamsLen()
	var params = make(map[string]string)
	if paramlen != 0 {
		params = c.Ctx.Input.Params()
		splat := c.Ctx.Input.Param(":splat")
		if len(splat) != 0 {
			params["splat"] = splat
		}
		ext := c.Ctx.Input.Param("ext")
		if len(ext) != 0 {
			params["ext"] = ext
		}
	} else {
		c.Data["Msg"] = "No Params Get."
	}
	c.Data["Params"] = params
	c.TplName = "router/autoRouter.html"
}
