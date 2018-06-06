package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/ronow2cn/go-admin/controllers"
)

//增加拦截器。
var filterAdmin = func(ctx *context.Context) {
	url := ctx.Input.URL()
	logs.Info("##### filter url : %s, %v", url, ctx.Input.Session("userinfo"))
	//TODO 如果判断用户未登录。
	_, ok := ctx.Input.Session("userinfo").(string)
	if !ok && url != "/admin" {
		logs.Info("##### Redirect url : %s", url)
		ctx.Redirect(302, "/login")
		return
	}

}

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

	beego.InsertFilter("/admin/*", beego.BeforeRouter, filterAdmin)
}
