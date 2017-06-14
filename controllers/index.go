package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"bee_login_register_demo/models"
)

type IndexController struct {
	beego.Controller
}

func (index *IndexController) Get() {
	sess := index.StartSession()
	username := sess.Get("username")
	fmt.Println(username)
	if username == nil || username == "" {
		index.TplName = "index.tpl"
	} else {
		index.TplName = "success.tpl"
	}
}

func (index *IndexController) Post() {

	var user models.Users
	inputs := index.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("username")
	password  := inputs.Get("pwd")
	err := models.ValidateUser(user,password)
	if err == nil {
		u ,_ :=models.FindUser(user.Username )
		index.SetSession("Loginuser", fmt.Sprintf("%d", u.Id)+"||"+u.Username)
		fmt.Println("Loginuser:", index.GetSession("Loginuser"))
		index.TplName = "success.tpl"
	} else {
		fmt.Println(err)
		index.TplName = "error.tpl"
	}
}

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
		this.TplName = "index.tpl"
}