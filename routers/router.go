package routers

import (
	"opscenter/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
	beego.Router("/regist",&controllers.RegistController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/logout",&controllers.LogoutUserController{})
	beego.Router("/home",&controllers.HomeController{})
	beego.Router("/show/loading",&controllers.ShowloadingController{})
}
