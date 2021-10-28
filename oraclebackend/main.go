/*
 * @Author: your name
 * @Date: 2021-10-23 02:33:02
 * @LastEditTime: 2021-10-23 03:45:11
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: /TheCarProject/oraclebackend/main.go
 */
package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"log"
	_ "github.com/go-sql-driver/mysql"
	jsoniter "github.com/json-iterator/go"
	"oraclebackend/oracle"
	_ "oraclebackend/routers"
	"os"
	"strconv"
	"time"
)

const (
	Duration = 10 //当网络中的节点数量超过10个时，分批训练，否则显卡内存不足。
	NIID = 1 //1非独立同分布，0独立同分布
	NodeCount   = 4 //3个为训练节点，1个为聚合节点,一块显卡最多容纳十六个
	ClientCount = 1
	Dataset = "fmnist"//mnist, fmnist
	Model = "cnn"//mlp
	Global_Epoch = 2//全局聚合次数，在demo中采用cpu训练，通过减少训练轮次降低cpu的负载
	Local_Epoch = 2//本地聚合次数
	ViewID = 0
	Malicious = 0 //恶意节点，在Demo中不考虑恶意节点
)

var Clients []*oracle.Client

var json jsoniter.API

func init() {
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}

func main() {
	for i := 0; i < NodeCount; i++ {
		server := oracle.NewServer(i)
		go server.Start()
	}

	time.Sleep(2 * time.Second)

	fileName := "../oracle/output/debug.log." + strconv.FormatInt(time.Now().Unix(), 10) //在工程路径下和src同级，也可以写绝对路径，不过要注意转义符
	logFile, err := os.Create(fileName)                                        //创建该文件，返回句柄
	if err != nil {
		log.Fatalln(err)
	}
	//debugLog := log.New(logFile, "", log.Llongfile)
	defer logFile.Close() //确保文件在该函数执行完以后关闭

	//Clients = make([]*oracle.Client, ClientCount)
	//oracle.NewClient(0)
	//client := oracle.NewClient(0)
	//client.Start()
	//Clients[0] = client
	//for i := 0; i < ClientCount; i++ {
	//	fmt.Println("11111")
	//	time := Clients[i].EndTime.UnixNano() - Clients[i].StartTime.UnixNano()
	//	fmt.Println("time consume : ", float64(time)/1000000000)
	//	debugLog.Printf("time consume : %v", float64(time)/1000000000)
	//}
	orm.RegisterDataBase("default", "mysql", "root:lyj030325@tcp(127.0.0.1:3306)/the_car_project?charset=utf8")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
     //允许访问所有源
	AllowAllOrigins: true,
			//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
			//其中Options跨域复杂请求预检
	AllowMethods:   []string{"GET", "POST"},
			//指的是允许的Header的种类
	AllowHeaders:  []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			//公开的HTTP标头列表
	ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			//如果设置，则允许共享身份验证凭据，例如cookie
	AllowCredentials: true,
	}))
	beego.Run()
}

