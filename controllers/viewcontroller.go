package controllers

import (
	"strconv"

	"beego-blog/models"

	"github.com/astaxie/beego"
)

type ViewController struct {
	beego.Controller
}

func (c *ViewController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	c.Data["Post"] = models.GetBlog(id)
	c.Layout = "layout.tpl"
	c.TplName = "view.tpl"
}
