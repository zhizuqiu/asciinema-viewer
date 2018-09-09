package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"fmt"
	"asciinema-viewer/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) List() {
	var lists []models.Path
	rd, err := ioutil.ReadDir("upload")
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = lists
		c.ServeJSON()
		return
	}

	index := 0
	for _, f := range rd {
		if !f.IsDir() {
			lists = append(lists, models.Path{Id: index, Path: "upload/" + f.Name()})
			index++
		}
	}
	c.Data["json"] = lists
	c.ServeJSON()
}

func (c *MainController) Viewer() {
	path := c.GetString("path")
	c.Data["path"] = path
	c.TplName = "viewer.html"
}
