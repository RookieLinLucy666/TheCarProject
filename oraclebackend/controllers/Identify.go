package controllers

import (
	"encoding/json"
	"fmt"
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
	//if err := c.ParseForm(&u); err != nil {
	//	c.Data["json"] = err.Error()
	//}
	json.Unmarshal(c.Ctx.Input.RequestBody, &u)
	fmt.Println(c.Ctx.Input)
	rst := xuperchain.InvokeAddUser(u.Name, u.Abstract)
	fmt.Println(rst)
	c.Data["json"] = struct {
		Desc	string `json:"desc"`
	}{
		Desc: rst,
	}
	c.ServeJSON()
}

func (c *IdentifyController) CheckUser()  {
	var u User
	if err := c.ParseForm(&u); err != nil {
		c.Data["json"] = err.Error()
	}
	fmt.Println(u)
	rst, _ := strconv.Atoi(xuperchain.InvokeCheckUser(u.Name, u.Abstract))
	fmt.Println(rst)
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
	c.ServeJSON()
}