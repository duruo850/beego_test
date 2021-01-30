package controllers

import "github.com/beego/beego/v2/server/web"

func init() {
	_ = web.LoadAppConfig("ini", "conf/app.conf")
}
