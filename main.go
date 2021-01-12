package main

import (
	_ "beego_test/routers"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	fmt.Println("hello beego test3")
	web.Run(":8080")
}
