package controllers

import (
	"beego_test/models"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	fmt.Println("beego 欢迎你")
	c.ViewPath = ""
	c.TplName = "index.html"
}

func (c *MainController) AddUser() {
	var user models.User
	_ = c.ParseForm(&user)
	addUser, err := models.AddUser(user)
	if err != nil {
		fmt.Println("添加失败")
		return
	}
	c.Data["ID"] = addUser.ID
	c.Data["Name"] = addUser.Name
	c.Data["Password"] = addUser.Password
	c.TplName = "query.html"
}

func (c *MainController) QueryUser() {
	name := c.GetString("name")
	usermsg, _ := models.GetUser(name)
	c.Data["ID"] = usermsg.ID
	c.Data["Name"] = usermsg.Name
	c.Data["Password"] = usermsg.Password
	c.TplName = "query.html"
}
