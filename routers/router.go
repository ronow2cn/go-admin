package routers

import (
	"github.com/astaxie/beego"
	"github.com/ronow2cn/go-admin/controllers"
)

func init() {

	MainController := &controllers.MainController{}
	beego.Router("/", MainController, "*:Index")
	beego.Router("/index", MainController, "*:Index")
	beego.Router("/login", MainController, "*:Login")
	beego.Router("/loginreq", MainController, "*:LoginReq")
	beego.Router("/logoutreq", MainController, "*:LogoutReq")

	userInfoController := &controllers.UserInfoController{}
	beego.Router("/admin/userInfo/edit", userInfoController, "get:Edit")
	beego.Router("/admin/userInfo/delete", userInfoController, "post:Delete")
	beego.Router("/admin/userInfo/save", userInfoController, "post:Save")
	beego.Router("/admin/userInfo/list", userInfoController, "get:List")
}
