package main

import (
	jsonencode "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//定义api接口数据结构和序列化json字段
type Data struct {
	ID    int `json:"id"`
	IP    string `json:"IP"`
	DESC  string `json:"desc"`
	OWNER string `json:"owner"`
}

type CloumnsData struct {
	NAME  string `json:"name"`
	ALIAS string `json:"alias"`
}

type Employee struct {
	CODE    int `json:"code"`
	ISADMIN bool `json:"isadmin"`
	DATA    []Data `json:"data"`
	COLUMNS []CloumnsData `json:"columns"`
	MESSGAE string `json:"messgae"`
}

//发送http请求和json 序列化并打印数据结构;

func apitest() {
	url := "http://ops-environment.com/channel/ip/v1"
	resp, _ := http.Get(url)
	s := Employee{}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	jsonencode.Unmarshal([]byte(body), &s)
	fmt.Println(fmt.Sprintf("%+v",s))

}
