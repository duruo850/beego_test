package main

import (
	_ "beego_test/controllers"
	"beego_test/lib/mysql"
	_ "beego_test/routers"
	"github.com/beego/beego/v2/server/web"
	"os"
)

func main() {
	mysql.Update()

	dir, _ := os.Getwd()
	web.SetViewsPath(dir + "/views2/")

	port, _ := web.AppConfig.String("Port")
	web.Run(":" + port)
}
