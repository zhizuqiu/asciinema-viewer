package controllers

import (
	"fmt"
)

type FileUploadController struct {
	BaseController
}

//文件下载方法
func (this *FileUploadController) Download() {
	filename := this.GetString("filename")
	fmt.Println("filename:", filename)

	this.Ctx.Output.Download("./upload/demo.txt")
}

func (this *FileUploadController) Upload() {
	name := "myfile"
	fmt.Println("Upload")
	f, h, _ := this.GetFile(name)
	path := "./upload/" + h.Filename
	f.Close()
	this.SaveToFile(name, path)
}
