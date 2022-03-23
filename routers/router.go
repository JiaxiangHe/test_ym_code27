package routers

import (
	"github.com/astaxie/beego"
	"github.com/test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
