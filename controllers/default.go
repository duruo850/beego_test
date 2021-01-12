package controllers

import (
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
