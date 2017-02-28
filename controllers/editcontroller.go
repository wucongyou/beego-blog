package controllers

import (
	"strconv"

	"beego-blog/models"

	"github.com/astaxie/beego"
)

type EditController struct {
	beego.Controller
}

func (c *EditController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	c.Data["Post"] = models.GetBlog(id)
	c.Layout = "layout.tpl"
	c.TplName = "edit.tpl"
}

func (c *EditController) Post() {
	params := c.Input()
	var blog models.Blog
	blog.Id, _ = strconv.Atoi(params.Get("id"))
	blog.Title = params.Get("title")
	blog.Content = params.Get("content")
	models.UpdateBlog(blog)
	c.Ctx.Redirect(302, "/")
}
