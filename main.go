package main

import (
	"flag"

	_ "beego-blog/routers"

	"fmt"
	"github.com/astaxie/beedb"
	"github.com/astaxie/beego"
)

var (
	debug bool
)

func init() {
	flag.BoolVar(&debug, "debug", true, "beedb debug switch")
}

func main() {
	flag.Parse()
	fmt.Println("debug: ", debug)
	beedb.OnDebug = debug
	beego.Run()
}
