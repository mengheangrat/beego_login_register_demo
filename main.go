package main

import (
	_ "bee_login_register_demo/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:mypassword@tcp(192.168.31.128:3306)/bee_test?charset=utf8")
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Run()
	//orm.RunCommand()
}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("userLogin").(string)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}

