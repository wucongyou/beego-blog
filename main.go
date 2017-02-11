package main

import (
	_ "beego-blog/routers"
	"github.com/astaxie/beedb"
	"github.com/astaxie/beego"
)

func main() {
	beedb.OnDebug = true
	beego.Run()
}
