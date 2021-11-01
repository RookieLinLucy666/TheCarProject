package xuperchain

import (
	"encoding/json"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/client/xuperchain"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	nodereq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/xuperchain/node"
	userreq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/xuperchain/user"
	"log"
)

const (
	UserId = "tangpingduibudui"
	UserAddr = "21z5qmVyXU6hJFwuDjRM96fyyd6CQ4fv4Q" //正式网
	//UserAddr = "2AUp6bPChg88sWEVEAD8qoEj1egTxzYRwj" //测试网
	ContractName = "cc_appxc_01" //正式网
	//ContractName = "xc00000989012866" //测试网
	APIAddr = "https://weifangnode.bsngate.com:17602/" //正式网
	//APIAddr = "http://52.83.209.158:17502" //测试网
	AppCode = "app0001202111011059523531748" //正式网
	//AppCode = "app006102fa0d96fe047b8b715ee11c7de464f" //测试网
	UserCode = "rookielinlucy"

)

var metadata Metadata
var learning FederatedAIDemand

type LearningEvent struct {
	Id string	`json:"id"`
	Metadata string `json:"meta_data_byte"`
	FaderatedAIDemandByte string `json:"faderated_ai_demand_byte"`
}

type DataEvent struct {
	Id string	`json:"id"`
	Metadata string `json:"meta_data"`
}


type Metadata struct {
	Uploader	string `json:"uploader"`
	Name 		string `json:"name"`
	Type		string `json:"type"`
	Ip			string `json:"ip"`
	Route		string	`json:"route"`
	Abstract	string	`json:"abstract"`
}

type FederatedAIDemand struct {
	Model 		string `json:"model"`
	Dataset 	string `json:"dataset"`
	Round 		string `json:"round"`
	Epoch 		string `json:"epoch"`
}

func ConfigClient() *xuperchain.XuperChainClient{
	api:= APIAddr //Node gateway address
	userCode:= UserCode //User No.
	appCode := AppCode //Application No.
	puk := "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEIlh1C0iWAdcKnM/yAaZZT/42NVzT\nVyr31H9MDhHbPkp+/B3gsp5iZOr6OTAGO9jpN10/YMIrxt2IMg5auIEvMA==\n-----END PUBLIC KEY-----\n" //Application public key
	prk :="-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgemK0VVY7YUdStw1U\nka54fsVkWUhhRYC1vZRsK0rFKwegCgYIKoEcz1UBgi2hRANCAATc6OGru9Hzntpx\nOvNJ5srMFnOYrZGcIaC8Ed6KbbUs3na7ghgLAAVG22ZaVD1kuquy0p0s0vFjXvEx\nq9y0ps57\n-----END PRIVATE KEY-----" //Application private key
	mspDir:="" //Certificate storage directory
	cert :="" //Certificate

	//api:= APIAddr //Node gateway address
	//userCode:= UserCode //User No.
	//appCode := AppCode //Application No.
	//puk := "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAECwJ5ftuqndO9H3ks1hD8cB6IA9lx\n/b0Z2hnFZ77rgRm9Q4lY1aqIhkM63Lh6X7uwPsoRC1xkS0PMp5x/jnRWcw==\n-----END PUBLIC KEY-----" //Application public key
	//prk :="-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgHaEvjmM9ZMt0xCHT\nY65RBRkWxY9bBfl/Fag0bvP1r9OgCgYIKoEcz1UBgi2hRANCAATQeZSDbPzUA57d\nUZTQBjdiY36CNk6ecsEuMvG3XpNxoJzome32RDEUkDc/qihPAmHaK48SCuVaoG5B\nHk+QBDaJ\n-----END PRIVATE KEY-----" //Application private key
	//mspDir:="" //Certificate storage directory
	//cert :="" //Certificate

	config,err :=config.NewConfig(api, userCode, appCode, puk, prk, mspDir, cert )
	if err !=nil{
		log.Fatal(err)
	}
	client, err := xuperchain.NewXuperChainClient(config)
	if err !=nil{
		log.Fatal(err)
	}
	return client
}

func RegisterUser() (string, string){
	client := ConfigClient()
	body := userreq.RegisterUserReqDataBody{
		UserId: UserId,
	}
	res, err := client.RegisterUser(body)
	if err != nil {
		log.Fatal(err)
	}
	if res.Header.Code != 0 {
		log.Fatal(res.Header.Msg)
	}
	fmt.Println(res.Body.UserId, res.Body.UserAddr)
	return res.Body.UserId, res.Body.UserAddr
}

func InvokeCreateAdvBSN(uploader string, name string, argtype string, ip string, route string, abstract string) string{
	xchainClient := ConfigClient()
	param := Metadata{
		Uploader: uploader,
		Name: name,
		Type: argtype,
		Ip: ip,
		Route: route,
		Abstract: abstract,
	}
	paramStr, _ := json.Marshal(param)
	tmp := struct {
		Value	string `json:"value"`
	}{
		Value: string(paramStr),
	}
	tmpStr, _ := json.Marshal(tmp)
	body := nodereq.CallContractReqDataReqDataBody{
		UserId:       UserId,
		UserAddr:     UserAddr,
		ContractName: ContractName,
		FuncName:     "createAdv",
		FuncParam:    string(tmpStr),
	}
	res, err := xchainClient.ReqChainCode(body)
	if err != nil {
		log.Fatal(err)
	}
	if res.Header.Code != 0 {
		log.Fatal(res.Header.Msg)
	}

	return res.Body.TxId
}

func GetVariable() (Metadata, FederatedAIDemand) {
	return metadata, learning
}

func InvokeAddUserBSN(name string, abstract string) string{
	xchainClient := ConfigClient()
	param := struct {
		Name		string `json:"name"`
		Abstract	string `json:"abstract"`
	}{
		Name: name,
		Abstract: abstract,
	}
	paramStr, _ := json.Marshal(param)
	body := nodereq.CallContractReqDataReqDataBody{
		UserId:       UserId,
		UserAddr:     UserAddr,
		ContractName: ContractName,
		FuncName:     "addUser",
		FuncParam:    string(paramStr),

	}
	res, err := xchainClient.ReqChainCode(body)
	if err != nil {
		log.Fatal(err)
	}
	if res.Header.Code != 0 {
		log.Fatal(res.Header.Msg)
	}
	return res.Body.TxId
}

func QueryVerifyUserBSN(name string, abstract string) string {
	xchainClient := ConfigClient()
	param := struct {
		Name		string `json:"name"`
		Abstract	string `json:"abstract"`
	}{
		Name: name,
		Abstract: abstract,
	}
	paramStr, _ := json.Marshal(param)
	body := nodereq.CallContractReqDataReqDataBody{
		UserId:       UserId,
		UserAddr:     UserAddr,
		ContractName: ContractName,
		FuncName:     "verifyUser",
		FuncParam:    string(paramStr),

	}
	res, err := xchainClient.ReqChainCode(body)
	if err != nil {
		log.Fatal(err)
	}
	if res.Header.Code != 0 {
		log.Fatal(res.Header.Msg)
	}
	fmt.Println("QueryVerifyUser: ", res.Body.QueryInfo)
	return res.Body.QueryInfo
}

func QueryMetaId() string{
	xchainClient := ConfigClient()

	body := nodereq.CallContractReqDataReqDataBody{
		UserId:       UserId,
		UserAddr:     UserAddr,
		ContractName: ContractName,
		FuncName:     "queryMetaId",
		FuncParam:    "{\"key\":\"zxlcounter\"}",

	}
	res, err := xchainClient.ReqChainCode(body)
	if err != nil {
		log.Fatal(err)
	}
	if res.Header.Code != 0 {
		log.Fatal(res.Header.Msg)
	}
	return res.Body.QueryInfo
}

func InvokeQueryBSN(metaId string) string {
	xchainClient := ConfigClient()
	param := struct {
		Id		string `json:"id"`
	}{
		Id: metaId,
	}
	paramStr, _ := json.Marshal(param)
	body := nodereq.CallContractReqDataReqDataBody{
		UserId:       UserId,
		UserAddr:     UserAddr,
		ContractName: ContractName,
		FuncName:     "query",
		FuncParam:    string(paramStr),

	}
	res, err := xchainClient.ReqChainCode(body)
	if err != nil {
		log.Fatal(err)
	}
	if res.Header.Code != 0 {
		log.Fatal(res.Header.Msg)
	}
	return res.Body.TxId
}

func InvokeQueryCallBackBSN(metaId string, data string) string {
	xchainClient := ConfigClient()
	param := struct {
		Id		string `json:"id"`
		Data	string `json:"data"`
	}{
		Id: metaId,
		Data: data,
	}
	paramStr, _ := json.Marshal(param)
	body := nodereq.CallContractReqDataReqDataBody{
		UserId:       UserId,
		UserAddr:     UserAddr,
		ContractName: ContractName,
		FuncName:     "queryCallback",
		FuncParam:    string(paramStr),

	}
	res, err := xchainClient.ReqChainCode(body)
	if err != nil {
		log.Fatal(err)
	}
	if res.Header.Code != 0 {
		log.Fatal(res.Header.Msg)
	}
	return res.Body.TxId
}

func InvokeComputingshareBSN(metaId string, model string, dataset string, round string, epoch string) string {
	xchainClient := ConfigClient()
	param := struct {
		Id		string `json:"id"`
		Model	string `json:"model"`
		Dataset	string `json:"dataset"`
		Round	string `json:"round"`
		Epoch 	string `json:"epoch"`
	}{
		Id: metaId,
		Model: model,
		Dataset: dataset,
		Round: round,
		Epoch: epoch,
	}
	paramStr, _ := json.Marshal(param)
	body := nodereq.CallContractReqDataReqDataBody{
		UserId:       UserId,
		UserAddr:     UserAddr,
		ContractName: ContractName,
		FuncName:     "computingshare",
		FuncParam:    string(paramStr),

	}
	res, err := xchainClient.ReqChainCode(body)
	if err != nil {
		log.Fatal(err)
	}
	if res.Header.Code != 0 {
		log.Fatal(res.Header.Msg)
	}
	return res.Body.TxId
}

func InvokeComputingCallBackBSN(metaId string, data string) string {
	xchainClient := ConfigClient()
	param := struct {
		Id		string `json:"id"`
		Data	string `json:"data"`
	}{
		Id: metaId,
		Data: data,
	}
	paramStr, _ := json.Marshal(param)
	body := nodereq.CallContractReqDataReqDataBody{
		UserId:       UserId,
		UserAddr:     UserAddr,
		ContractName: ContractName,
		FuncName:     "computingCallback",
		FuncParam:    string(paramStr),

	}
	res, err := xchainClient.ReqChainCode(body)
	if err != nil {
		log.Fatal(err)
	}
	if res.Header.Code != 0 {
		log.Fatal(res.Header.Msg)
	}
	return res.Body.TxId
}

func RunBSN() {
	//RegisterUser()
	//InvokeCreateAdv("xuperchain", "1111", "data", "local", "local", "dasdasds")
	//InvokeAddUser("Mengeshall", "1111111111")
	//QueryVerifyUser("Mengeshall", "1111111111")
	//InvokeQuery("AdvFileAssetId_2")
	//InvokeQueryCallBack("AdvFileAssetId_1", "97.2")
	//InvokeComputingshare("AdvFileAssetId_1", "cnn", "mnist", "1", "1")
	//InvokeComputingCallBack("AdvFileAssetId_1", "97.2")
	//QueryMetaId()
}