package main

import (
	"flag"

	_ "beego-blog/routers"

	"github.com/astaxie/beedb"
	"github.com/astaxie/beego"
)

var (
	debug bool
)

func init() {
	flag.BoolVar(&debug, "debug", true, "debug switch")
}

func main() {
	flag.Parse()
	beedb.OnDebug = debug
	beego.Run()
}
