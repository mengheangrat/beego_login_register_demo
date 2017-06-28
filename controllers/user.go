package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"opscenter/models"
)

type LoginController struct {
	beego.Controller
}

func (index *LoginController) Get() {
	index.TplName = "login.tpl"
}


func (index *LoginController) Post() {
	var user models.Users
	inputs := index.Input()
	user.Username = inputs.Get("username")
	password  := inputs.Get("password")
	err := models.ValidateUser(user,password)
	if err == nil {
		u ,_ :=models.FindUser(user.Username )
		index.SetSession("userLogin", fmt.Sprintf("%d", u.Id))
		index.SetSession("username", fmt.Sprintf("%s", u.Username))
		index.Redirect("/",302)
	} else {
		fmt.Println(err)
		index.TplName = "error.tpl"
	}
}




type RegistController struct {
	beego.Controller
}

func (index *RegistController) Get() {
	index.TplName = "regist.tpl"
}

func (index *RegistController) Post() {
	inputs := index.Input()
	//fmt.Println(inputs)
	username := inputs.Get("username")
	Pwd := inputs.Get("pwd")
	err :=models.AddUser(username,Pwd)
	if err == nil {
		index.TplName = "success.tpl"
	} else {
		fmt.Println(err)
		index.TplName = "error.tpl"
	}
}

type LogoutUserController struct {
	beego.Controller
}

func (index *LogoutUserController) Get() {
	index.DelSession("userLogin")
	index.DelSession("username")
	index.Redirect("/login", 302)
}