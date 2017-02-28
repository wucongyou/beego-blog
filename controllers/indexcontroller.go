package controllers

import (
	"beego-blog/models"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Data["blogs"] = models.GetAll()
	c.Layout = "layout.tpl"
	c.TplName = "index.tpl"
}
