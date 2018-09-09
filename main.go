package main

import (
	_ "asciinema-viewer/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.StaticDir["/dist"] = "static/dist"
	beego.BConfig.WebConfig.StaticDir["/upload"] = "upload"
	beego.Run()
}

