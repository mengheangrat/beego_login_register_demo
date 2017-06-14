package controllers

import(
	"fmt"
	"github.com/astaxie/beego"
	"bee_login_register_demo/models"
)

type RegistController struct {
	beego.Controller
}

func (this *RegistController) Get() {
	this.TplName = "regist.tpl"
}

func (this *RegistController) Post() {
	inputs := this.Input()
	//fmt.Println(inputs)
	username := inputs.Get("username")
	Pwd := inputs.Get("pwd")
	err :=models.AddUser(username,Pwd)

	if err == nil {
		this.TplName = "success.tpl"
	} else {
		fmt.Println(err)
		this.TplName = "error.tpl"
	}
}