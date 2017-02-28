package controllers

import (
	"time"

	"beego-blog/models"

	"github.com/astaxie/beego"
)

type NewController struct {
	beego.Controller
}

func (c *NewController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "new.tpl"
}

func (c *NewController) Post() {
	params := c.Input()
	var blog models.Blog
	blog.Title = params.Get("title")
	blog.Content = params.Get("content")
	blog.Ctime = time.Now()
	blog.Mtime = time.Now()
	models.SaveBlog(blog)
	c.Ctx.Redirect(302, "/")
}
