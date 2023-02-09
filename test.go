package main

import (
	_ "beeblog/models"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
