package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"beego-blog/models"
)

type EditController struct {
	beego.Controller
}

func (this *EditController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.tpl"
	this.TplName = "edit.tpl"
}

func (this *EditController) Post() {
	inputs := this.Input()
	var blog models.Blog
	blog.Id, _ = strconv.Atoi(inputs.Get("id"))
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	models.UpdateBlog(blog)
	this.Ctx.Redirect(302, "/")
}
