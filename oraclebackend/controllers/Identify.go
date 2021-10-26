package controllers

import (
	"github.com/astaxie/beego"
	"oraclebackend/xuperchain"
	"strconv"
)

type IdentifyController struct {
	beego.Controller
}

type User struct {
	Name		string `json:"name"`
	Abstract	string `json:"abstract"`
}

func (c *IdentifyController) AddUser()  {
	var u User
	if err := c.ParseForm(&u); err != nil {
		c.Data["json"] = err.Error()
	}
	rst := xuperchain.InvokeAddUser(u.Name, u.Abstract)
	c.Data["json"] = rst
}

func (c *IdentifyController) CheckUser()  {
	var u User
	if err := c.ParseForm(&u); err != nil {
		c.Data["json"] = err.Error()
	}
	rst, _ := strconv.Atoi(xuperchain.InvokeCheckUser(u.Name, u.Abstract))
	if rst == 0 {
		c.Data["json"] = struct {
			Code	int `json:"code"`
			Desc	string `json:"describe"`
		}{
			rst,
			"user does not have permission",
		}
	} else {
		c.Data["json"] = struct {
			Code	int `json:"code"`
			Desc	string `json:"describe"`
		}{
			rst,
			"permission verification passed",
		}
	}
}