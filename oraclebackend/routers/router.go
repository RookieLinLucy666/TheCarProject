package routers

import (
	"github.com/astaxie/beego"
	"oraclebackend/controllers"
)

func init() {
	beego.Router("/datashare/createdata", &controllers.DataShareController{}, "get:CreateData")
    beego.Router("/", &controllers.MainController{})
}
