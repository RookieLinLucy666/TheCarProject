package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"oraclebackend/xuperchain"
)

type DataShareController struct {
	beego.Controller
}

func (c *DataShareController) UrlMapping()  {
	c.Mapping("CreateData", c.CreateData)
	//c.Mapping("Post", c.Post)
	//c.Mapping("GetOne", c.GetOne)
	//c.Mapping("GetAll", c.GetAll)
	//c.Mapping("Put", c.Put)
	//c.Mapping("Delete", c.Delete)
}

func (c *DataShareController) CreateData()  {
	var v xuperchain.Metadata
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil{
		c.Data["json"] = err.Error()
	}
	id := xuperchain.InvokeCreateCfa(v.Uploader, v.Name, v.Type, v.Ip, v.Route, v.Abstract)

}