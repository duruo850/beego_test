package routers

import (
	"beego_test/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/add", &controllers.MainController{}, "Post:AddUser")
	web.Router("/query", &controllers.MainController{}, "Post:QueryUser")
	web.Router("/get", &controllers.MainController{}, "Get:GetUser")
	web.Router("/consul", &controllers.ConsulController{}, "Get:GetConsul")
}
