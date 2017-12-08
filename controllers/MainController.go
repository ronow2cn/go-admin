package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	userinfo := c.GetSession("userinfo")
	if userinfo == nil {
		c.Ctx.Redirect(302, "/login")
		return
	}

	c.TplName = "index.html"
}

func (c *MainController) Login() {
	c.TplName = "login/login.html"
}

func (c *MainController) LoginReq() {
	username := c.GetString("username")
	passwd := c.GetString("password")

	if username == "admin" && passwd == "admin" {
		c.SetSession("userinfo", username)

		c.Data["json"] = &map[string]interface{}{"status": true, "info": "success"}
		c.ServeJSON()
		return
	} else {

		c.Data["json"] = &map[string]interface{}{"status": false, "info": "failed"}
		c.ServeJSON()
		return
	}

}

func (c *MainController) LogoutReq() {
	c.DelSession("userinfo")
	c.Ctx.Redirect(302, "/login")
}
