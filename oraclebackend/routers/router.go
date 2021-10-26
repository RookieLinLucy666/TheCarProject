/*
 * @Author: your name
 * @Date: 2021-10-23 02:33:03
 * @LastEditTime: 2021-10-23 03:26:43
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /TheCarProject/oraclebackend/routers/router.go
 */
package routers

import (
	"github.com/astaxie/beego"
	"oraclebackend/controllers"
)

func init() {
	beego.Router("/datashare/createdata", &controllers.DataShareController{}, "post:CreateData")
	beego.Router("/datashare/query", &controllers.DataShareController{}, "get:Query")
	beego.Router("/datashare/computingshare", &controllers.DataShareController{}, "post:ComputingShare")
	beego.Router("/metadata/getall", &controllers.MetadataController{}, "get:GetAll")
	beego.Router("/identify/adduser", &controllers.IdentifyController{}, "post:AddUser")
	beego.Router("/identify/checkuser", &controllers.IdentifyController{}, "post:CheckUser")
    beego.Router("/", &controllers.MainController{})
}
