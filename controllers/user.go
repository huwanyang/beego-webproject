package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"webproject/bean"
)

type UserController struct {
	beego.Controller
}

/*
	http://localhost:8080/api/v1/user/index?uuid=sdgg45678&user.id=11&user.Name=wanyang3&user.Age=33&user.Email=wanyang3@ss.cm
*/
func (this *UserController) Get() {
	var uuid string
	var user = bean.User{}
	this.Ctx.Input.Bind(&uuid, "uuid")
	this.Ctx.Input.Bind(&user, "user")
	this.Data["uuid"] = uuid
	this.Data["user"] = user
	this.TplName = "user/index.html"
}

func (this *UserController) Post() {
	method := this.Ctx.Request.Method
	name1 := this.Ctx.Request.PostForm.Get("username")
	username := this.GetString("username")
	age_str := this.Input().Get("age")
	age, _ := strconv.Atoi(age_str)
	email := this.GetString("email")
	userinfo := map[string]interface{}{
		"name":   username,
		"name1":  name1,
		"method": method,
		"age":    age,
		"email":  email,
	}
	this.Ctx.Output.JSON(userinfo, false, false)
}

func (this *UserController) Save() {
	user := bean.User{}
	err := this.ParseForm(&user)
	if err == nil {
		m, h, err := this.GetFile("uploadImage")
		if err != nil {
			this.Ctx.Output.JSON(err, false, false)
		} else {
			defer m.Close()
			this.SaveToFile("uploadImage", "static/download/"+h.Filename)
		}
		this.Ctx.Output.JSON(user, false, false)
	} else {
		this.Ctx.Output.JSON(err, false, false)
	}
}

func (this *UserController) JsonPrint() {
	var user = bean.User{}
	user.Name = "wanyang3"
	user.Age = 23
	user.Email = "wanyang3@staff.weibo.com"
	this.Data["json"] = &user
	this.ServeJSON()
}

func (this *UserController) XmlPrint() {
	var user = bean.User{}
	user.Name = "wanyang3"
	user.Age = 23
	user.Email = "wanyang3@staff.weibo.com"
	this.Data["xml"] = &user
	this.ServeXML()
}

func (this *UserController) JsonpPrint() {
	var user = bean.User{}
	user.Name = "wanyang3"
	user.Age = 23
	user.Email = "wanyang3@staff.weibo.com"
	this.Data["jsonp"] = &user
	this.ServeJSONP()
}

func (this *UserController) FormatPrint(){
	var user = bean.User{}
	user.Name = "wanyang3"
	user.Age = 23
	user.Email = "wanyang3@staff.weibo.com"
	this.Data["json"] = &user
	this.Data["xml"] = &user
	this.ServeFormatted()
}