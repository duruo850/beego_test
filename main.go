package main

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	fmt.Println("hello beego test3")
	web.Run(":8080")
}
