package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.Data["Title"] = "Beego Admin"
	c.Data["Content"] = "This is Beego Admin content. "
	mysql := make(map[string]interface{})
	mysql["MysqlUrl"] = beego.AppConfig.String("mysqlurl")
	mysql["MysqlUser"] = beego.AppConfig.String("mysqluser")
	mysql["MysqlPassword"] = beego.AppConfig.String("mysqlpassword")
	mysql["MysqlDB"] = beego.AppConfig.String("mysqldb")
	c.Data["Mysql"] = mysql
	c.TplName = "admin/index.tpl"
}

func (c *AdminController) ListBook() {
	// simple 类型
	c.Data["Title"] = "Beego Admin List"
	c.Data["Total"] = 3
	// slice 类型
	var bookList = make([]string, 0)
	bookList = append(bookList, "Golang Book")
	bookList = append(bookList, "Beego Book")
	bookList = append(bookList, "AutoTesting Book")
	c.Data["BookList"] = bookList
	// struct 类型
	user := new(User)
	user.Name = "Jack"
	user.Age = 20
	user.Info = "Beijing China."
	user.Sex = 1
	c.Data["UserInfo"] = &user
	// map 类型
	maps := make(map[string]interface{})
	maps["key1"] = "value1"
	maps["key2"] = 1234
	maps["key3"] = 12.22
	c.Data["mp"] = maps
	c.TplName = "admin/adminlist.tpl"
}

type User struct {
	Id   int    `form:"-"`
	Name string `form:"username,,姓名: "`
	Age  int    `form:",,年龄"`
	Sex  int    `form:",checkbox,性别"`
	Info string `form:",textarea"`
}

func (c *AdminController) FormPage() {
	c.Data["Form"] = &User{}
	c.TplName = "admin/save.html"
}

func (c *AdminController) AdminConf() {
	dev := make(map[string]interface{})
	dev["httpport"] = beego.AppConfig.String("httpport")
	dev["httpport_new"] = beego.AppConfig.String("dev::httpport")
	dev["mysqlurl"] = beego.AppConfig.String("mysqlurl")
	dev["mysqlurl_new"] = beego.AppConfig.String("dev::mysqlurl")
	dev["mysqluser"] = beego.AppConfig.String("dev::mysqluser")
	dev["mysqlpassword"] = beego.AppConfig.String("dev::mysqlpassword")
	dev["mysqldb"] = beego.AppConfig.String("dev::mysqldb")
	c.Data["dev"] = dev
	prod := make(map[string]interface{})
	prod["httpport"] = beego.AppConfig.String("prod::httpport")
	prod["mysqlurl"] = beego.AppConfig.String("prod::mysqlurl")
	prod["mysqlname"] = beego.AppConfig.String("prod::mysqlname")
	prod["mysqluser"] = beego.AppConfig.String("prod::mysqluser")
	prod["mysqlpassword"] = beego.AppConfig.String("prod::mysqlpassword")
	prod["mysqldb"] = beego.AppConfig.String("prod::mysqldb")
	c.Data["prod"] = prod
	test := make(map[string]interface{})
	test["httpport"] = beego.AppConfig.String("test::httpport")
	test["mysqlurl"] = beego.AppConfig.String("test::mysqlurl")
	test["mysqluser"] = beego.AppConfig.String("test::mysqluser")
	test["mysqlpassword"] = beego.AppConfig.String("test::mysqlpassword")
	test["mysqldb"] = beego.AppConfig.String("test::mysqldb")
	c.Data["test"] = test
	c.TplName = "admin/conf.html"
}

func (c *AdminController) AdminRouter() {
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
