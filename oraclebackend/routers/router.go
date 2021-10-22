package routers

import (
	"github.com/astaxie/beego"
	"oraclebackend/controllers"
)

func init() {
	beego.Router("/datashare/createdata", &controllers.DataShareController{}, "post:CreateData")
	beego.Router("/datashare/query", &controllers.DataShareController{}, "get:Query")
	beego.Router("/datashare/computingshare", &controllers.DataShareController{}, "post:ComputingShare")
    beego.Router("/", &controllers.MainController{})
}
