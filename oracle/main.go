package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
	"strconv"
	"sync"
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

var Clients []*Client

var json jsoniter.API

func init() {
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}

/**
  main
  @Description: 开启客户端节点
**/
func main() {
	for i := 0; i < NodeCount; i++ {
		server := NewServer(i)
		go server.Start()
	}

	time.Sleep(2 * time.Second)

	fileName := "output/debug.log." + strconv.FormatInt(time.Now().Unix(), 10) //在工程路径下和src同级，也可以写绝对路径，不过要注意转义符
	logFile, err := os.Create(fileName)                                        //创建该文件，返回句柄
	if err != nil {
		log.Fatalln("open file error !")
	}
	debugLog := log.New(logFile, "", log.Llongfile)
	defer logFile.Close() //确保文件在该函数执行完以后关闭

	Clients := make([]*Client, ClientCount)
	wg := sync.WaitGroup{}
	for i := 0; i < ClientCount; i++ {
		wg.Add(1)
		go func(i int32) {
			defer wg.Done()
			client := NewClient(i)
			client.Start()
			Clients[i] = client
		}(int32(i))
	}
	wg.Wait()
	for i := 0; i < ClientCount; i++ {
		time := Clients[i].EndTime.UnixNano() - Clients[i].StartTime.UnixNano()
		fmt.Println("time consume : ", float64(time)/1000000000)
		debugLog.Printf("time consume : %v", float64(time)/1000000000)
	}

}
