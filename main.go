package main

import (
	"github.com/astaxie/beego"
	"github.com/test/models"
	_ "github.com/test/routers"
)

func main() {
	models.LinkedDome()
	beego.Run()
}
