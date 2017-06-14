package routers

import (
	"bee_login_register_demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
	beego.Router("/regist",&controllers.RegistController{})
	beego.Router("/login",&controllers.LoginController{})
}
