/*
 * @Author: your name
 * @Date: 2021-10-23 02:33:02
 * @LastEditTime: 2021-10-23 04:37:57
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /TheCarProject/oraclebackend/controllers/datashare.go
 */
package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"net"
	"oraclebackend/models"
	"oraclebackend/oracle"
	"oraclebackend/xuperchain"
	"strconv"
	"time"
)

type DataShareController struct {
	beego.Controller
}

func (c *DataShareController) UrlMapping()  {
	c.Mapping("CreateData", c.CreateData)
	c.Mapping("Query", c.Query)
}

func (c *DataShareController) CreateData()  {
	var v xuperchain.Metadata
	if err := c.ParseForm(&v); err != nil{
		c.Data["json"] = err.Error()
	}
	v1 := models.Metadata{
		Uploader: v.Uploader,
		Name: v.Name,
		Type: v.Type,
		Ip: v.Ip,
		Route: v.Route,
		Abstract: v.Abstract,
	}
	v1.BcId = xuperchain.InvokeCreateCfa(v.Uploader, v.Name, v.Type, v.Ip, v.Route, v.Abstract)
	if _, err := models.AddMetadata(&v1); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v1
	} else {
		c.Data["json"] = err.Error()
	}
}

func (c *DataShareController) Query()  {
	var v *models.Metadata
	//idStr := c.Ctx.Input.Param(":id")
	idStr := "7"
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, _ = models.GetMetadataById(id)
	o := oracle.NewClient(0)
	o.SendRequest(v.Type, v.BcId, xuperchain.FederatedAIDemand{}, id)
	ln, err := net.Listen("tcp", o.Url)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		reply := o.HandleConnection(conn)
		if reply {
			break
		}
	}
	time.Sleep(5000)
	msg, _ := models.GetOneResultByBcId(v.BcId)
	c.Data["json"] = struct{
		Msg		string
	}{
		Msg: msg.Result,
	}
	c.ServeJSON()
}

func (c *DataShareController) ComputingShare()  {
	args := struct {
		Bcid		string `json:"bcid"`
		Type		string `json:"type"`
		Model 		string `json:"model"`
		Dataset 	string `json:"dataset"`
		Round 		string `json:"round"`
		Epoch 		string `json:"epoch"`
	}{}
	if err := c.ParseForm(&args); err != nil{
		c.Data["json"] = err.Error()
	}
	demand := xuperchain.FederatedAIDemand{
		Model: args.Model,
		Dataset: args.Dataset,
		Round: args.Round,
		Epoch: args.Epoch,
	}
	o := oracle.NewClient(0)
	fmt.Println(args)
	o.SendRequest(args.Type, args.Bcid, demand, -1)
	ln, err := net.Listen("tcp", o.Url)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		reply := o.HandleConnection(conn)
		if reply {
			break
		}
	}
	//time.Sleep(5000)
	//msg, _ := models.GetOneResultByBcId(args.Bcid)
	//c.Data["json"] = struct{
	//	Msg		string
	//}{
	//	Msg: msg.Result,
	//}
	//c.ServeJSON()
}