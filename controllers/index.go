package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (index *IndexController) Get() {
	sess := index.StartSession()
	username := sess.Get("username")
	if username == nil || username == "" {
		index.Redirect("/login",302)
	} else {
		index.Data["username"]=username
		index.TplName = "index.tpl"
	}
}
type HomeController struct {
	beego.Controller
}

func (index *HomeController) Get() {
	sess := index.StartSession()
	username := sess.Get("username")
	if username == nil || username == "" {
		index.Redirect("/login",302)
	} else {
		index.Data["username"]=username
		index.TplName = "home.html"
	}
}