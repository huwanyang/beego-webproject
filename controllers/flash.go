package controllers

import "github.com/astaxie/beego"

type FlashController struct {
	beego.Controller
}

func (c *FlashController) Get() {
	flash := beego.ReadFromRequest(&c.Controller)
	if _, ok := flash.Data["error"]; ok {
		c.Data["error"] = true
	}
	if _, ok := flash.Data["notice"]; ok {
		c.Data["notice"] = true
	}
	if _, ok := flash.Data["warning"]; ok {
		c.Data["warning"] = true
	}
	if _, ok := flash.Data["success"]; ok {
		c.Data["success"] = true
	}
	c.TplName = "flash/index.html"
}

func (c *FlashController) Post() {
	flash := beego.NewFlash()
	number, err := c.GetInt("number")
	if err != nil {
		flash.Error("Error: Get Number error.")
	} else if number < 0 {
		flash.Warning("Warning: Number < 0")
	} else if number > 0 && number < 10 {
		flash.Notice("Notice: Number > 0 && Number < 10.")
	} else {
		flash.Success("Success: Number > 10.", number)
	}
	flash.Store(&c.Controller)
	c.Redirect("/api/v1/flash/get", 302)
}
