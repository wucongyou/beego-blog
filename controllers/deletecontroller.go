package controllers

import (
	"strconv"

	"beego-blog/models"

	"github.com/astaxie/beego"
)

type DeleteController struct {
	beego.Controller
}

func (c *DeleteController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	blog := models.GetBlog(id)
	c.Data["Post"] = blog
	models.DelBlog(blog)
	c.Ctx.Redirect(302, "/")
}
