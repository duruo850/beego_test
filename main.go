package main

import (
	_ "beego_test/lib/mysql"
	_ "beego_test/routers"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"os"
)

func main() {
	fmt.Println("hello beego test3")

	dir, _ := os.Getwd()
	web.SetViewsPath(dir + "/views2/")
	web.Run(":8080")
}
