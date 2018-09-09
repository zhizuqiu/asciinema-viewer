package routers

import (
	"asciinema-viewer/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
    beego.Router("/", &controllers.MainController{})
	beego.Router("/list", &controllers.MainController{},"GET:List")
	beego.Router("/viewer", &controllers.MainController{},"GET:Viewer")
	beego.Router("/file/Upload", &controllers.FileUploadController{}, "POST:Upload")
}
