package main

import (
	_ "quickstart/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/down", "download")  
	beego.Run()
}

