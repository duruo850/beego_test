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
