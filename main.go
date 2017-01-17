package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beedb"
)

func main() {
	beedb.OnDebug = true
	beego.Run()
}

