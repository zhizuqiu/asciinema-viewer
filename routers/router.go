package routers

import (
	"asciinema-viewer/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/upload", &controllers.MainController{},"GET:Upload")
	beego.Router("/file/Upload", &controllers.FileUploadController{}, "POST:Upload")
}
