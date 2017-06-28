package controllers

import (
	"github.com/astaxie/beego"
)

type ShowloadingController struct {
	beego.Controller
}

func (index *ShowloadingController) Get() {
	sess := index.StartSession()
	username := sess.Get("username")
	if username == nil || username == "" {
		index.Redirect("/login",302)
	} else {
		index.Data["username"]=username
		index.TplName = "showloading.html"
	}
}
