package main

import (
	"beego_test/lib/mysql"
	_ "beego_test/routers"
	"github.com/beego/beego/v2/server/web"
	"os"
)

func main() {
	mysql.Update()

	dir, _ := os.Getwd()
	web.SetViewsPath(dir + "/views2/")
	web.Run(":8080")
}
