package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type ConsulController struct {
	web.Controller
}

// 定义返回的结构体，并转为json格式
type Result struct {
	Status int    `json:"status"` // 首字母大写
	Msg    string `json:"msg"`
}

func (c *ConsulController) GetConsul() {
	result := Result{0, "test"} // 赋值
	c.Data["json"] = &result
	_ = c.ServeJSON() // 返回json
	c.ServeFormatted()
}
