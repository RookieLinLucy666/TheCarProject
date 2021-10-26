package xuperchain

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/xuperchain/xuper-sdk-go/v2/account"
	"github.com/xuperchain/xuper-sdk-go/v2/xuper"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

const (
	Address = "ezDq8L2yHpqTSmKFs3HCiFmVELN67hFF3"
	Contract_Addr = "XC1234567890111113@xuper"
	Contract_Name = "SDKNativeCount16"
	Mnemonic = "抱 丙 德 斯 伐 珠 凸 踏 杆 寻 宣 取"
)


//var Mnemonic string
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

/**
  CreateAccount
  @Description: 创建账户
	命令行转钱：
	./xchain-cli transfer --to ehxodjwhVtESXcqpdUZ2LExs53DS4aHkE --amount 10000000000 --keys data/keys/ -H 127.0.0.1:37101
**/
func CreateAccount() {
	var acc *account.Account
	var err error
	// 测试创建账户
	acc, err = account.CreateAccount(1, 1)
	if err != nil {
		fmt.Printf("create account error: %v\n", err)
	}
	fmt.Println(acc)
	fmt.Println(acc.Mnemonic)

	//Mnemonic = acc.Mnemonic

	//// 测试恢复账户
	//acc, err = account.RetrieveAccount("钢 车 稀 叔 送 扫 永 确 描 当 矮 北}", 1)
	//if err != nil {
	//	fmt.Printf("retrieveAccount err: %v\n", err)
	//	return
	//}
	//fmt.Printf("RetrieveAccount: to %v\n", acc)
	//
	//// 创建账户并存储到文件中
	//acc, err = account.CreateAndSaveAccountToFile("./keys", "123", 1, 1)
	//if err != nil {
	//	fmt.Printf("createAndSaveAccountToFile err: %v\n", err)
	//}
	//fmt.Printf("CreateAndSaveAccountToFile: %v\n", acc)
	//
	//// 从文件中恢复账户
	//acc, err = account.GetAccountFromFile("keys/", "123")
	//if err != nil {
	//	fmt.Printf("getAccountFromFile err: %v\n", err)
	//}
	//fmt.Printf("getAccountFromFile: %v\n", acc)
	return
}

/**
  CreateContractAccount
  @Description: 创建合约账户
	命令行给合约账户转钱：
	./xchain-cli transfer --to XC1234567890111113@xuper --amount 1000000000000
**/
func CreateContractAccount() {
	// 从文件中恢复账户
	acc, err := account.RetrieveAccount(Mnemonic, 1)
	if err != nil {
		fmt.Printf("retrieveAccount err: %v\n", err)
		return
	}
	fmt.Printf("RetrieveAccount: to %v\n", acc)

	// 创建一个合约账户
	// 合约账户是由 (XC + 16个数字 + @xuper) 组成, 比如 "XC1234567890123456@xuper"
	contractAccount := Contract_Addr

	xchainClient, err := xuper.New("127.0.0.1:37101")
	tx, err := xchainClient.CreateContractAccount(acc, contractAccount)
	if err != nil {
		fmt.Printf("createContractAccount err:%s\n", err.Error())
	}
	fmt.Println(tx.Tx.Txid)
	return
}

func getAccount() *account.Account{
	account, err := account.RetrieveAccount(Mnemonic, 1)
	if err != nil {
		fmt.Printf("retrieveAccount err: %v\n", err)
		return nil
	}
	//fmt.Printf("retrieveAccount address: %v\n", account.Address)
	contractAccount := Contract_Addr
	err = account.SetContractAccount(contractAccount)
	if err != nil {
		panic(err)
	}
	return account
}

/**
  NativeContract
  @Description:
	部署合约
	命令行部署合约：在output目录下运行`./xchain-cli wasm deploy --account XC1234567890111111@xuper --cname gocounter -a '{"creator":"xchain"}' ./counter --runtime go`
**/
func DeployContract() {
	codePath := "contract/contract" // 编译好的二进制文件 go build -o counter
	code, err := ioutil.ReadFile(codePath)
	if err != nil {
		panic(err)
	}
	account := getAccount()
	contractName := Contract_Name
	xchainClient, err := xuper.New("127.0.0.1:37101")
	if err != nil {
		panic(err)
	}
	args := map[string]string{
		"creator": "xuperchain",
		"key":     "contract",
	}
	var tx *xuper.Transaction
	tx, err = xchainClient.DeployNativeGoContract(account, contractName, code, args)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deploy Native Go Contract Success! %x\n", tx.Tx.Txid)
}

/**
  InvokeCreateCfa
  @Description:  链下调用CreateCfa(创建任务id)
  @param uploader "xuperchain"
  @param name "counter"
  @param argtype "data"
  @param ip "127.0.0.1"
  @param route "xuperchain"
  @param abstract "162accb12e079d4b805f65f7a773c5e10cf537fef5ff99fde901ef0b1c963af8"
  @return string
**/
func InvokeCreateCfa(uploader string, name string, argtype string, ip string, route string, abstract string) string{
	account := getAccount()

	contractName := Contract_Name
	xchainClient, err := xuper.New("127.0.0.1:37101")
	if err != nil {
		panic(err)
	}
	args := map[string]string{
		"uploader": uploader,
		"name":     name,
		"type": argtype, //data, cross, compute
		"ip": ip,
		"route": route,
		"abstract": abstract,
	}

	var tx *xuper.Transaction
	tx, err = xchainClient.InvokeNativeContract(account, contractName, "CreateCfa", args)
	if err != nil {
		panic(err)
	}
	//fmt.Printf(string(tx.ContractResponse.Body))
	return string(tx.ContractResponse.Body)
}

/**
  Query
  @Description: 链下调用Query(数据协同的调用)
**/
func InvokeQuery(id string) []byte{
	account := getAccount()

	contractName := Contract_Name
	xchainClient, err := xuper.New("127.0.0.1:37101")
	if err != nil {
		panic(err)
	}
	args := map[string]string{
		"id": id,
		"inquirer": "xuperchain",
		"expiration": "1000",
	}

	var tx *xuper.Transaction

	tx, err = xchainClient.InvokeNativeContract(account, contractName, "Query", args)
	if err != nil {
		panic(err)
	}
	if tx != nil {
		fmt.Printf("InvokeQuery：%s\n", tx.ContractResponse.Body)
	}
	return tx.ContractResponse.Body
}

/**
  ListenQueryEvent
  @Description: 链下监听queryEvent事件
  @return error
**/
func ListenQueryEvent() error{
	// 创建节点客户端。
	client, err := xuper.New("127.0.0.1:37101")
	if err != nil {
		return err
	}

	// 监听时间，返回 Watcher，通过 Watcher 中的 channel 获取block。
	//过滤掉空数据
	//watcher, err := client.WatchBlockEvent(xuper.WithSkipEmplyTx())
	//通过事件名称去过滤
	watcher, err := client.WatchBlockEvent(xuper.WithEventName("queryEvent"))
	if err != nil {
		return err
	}

	go func() {
		for {
			b, ok := <-watcher.FilteredBlockChan
			if !ok {
				fmt.Println("watch block event channel closed.")
				return
			}
			//fmt.Printf("%+v\n", b)
			//QueryBlockByHeight(b.BlockHeight)
			if len(b.Txs) != 0 {
				QueryTxByID("queryEvent", b.Txs[0].Txid)
			}
		}
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("close watch")
	// 关闭监听。
	watcher.Close()
	client.Close()
	return nil
}

func InvokeQueryCallback(id string, data string, asig string, pks string) {
	account := getAccount()

	contractName := Contract_Name
	xchainClient, err := xuper.New("127.0.0.1:37101")
	if err != nil {
		panic(err)
	}
	args := map[string]string{
		"id": id,
		"data": data,
		"asig": asig,
		"pks": pks,
	}

	var tx *xuper.Transaction

	tx, err = xchainClient.InvokeNativeContract(account, contractName, "Query", args)
	if err != nil {
		panic(err)
	}
	if tx != nil {
		fmt.Printf("InvokeQueryCallback：%s\n", tx.ContractResponse.Body)
	}
}

/**
  InvokeComputingShare
  @Description: 调用ComputingShare合约
  @param id
  @param model "cnn"
  @param dataset "mnist"
  @param round "2"
  @param epoch "2"
**/
func InvokeComputingShare(id string, model string, dataset string, round string, epoch string) {
	account := getAccount()

	contractName := Contract_Name
	xchainClient, err := xuper.New("127.0.0.1:37101")
	if err != nil {
		panic(err)
	}
	args := map[string]string{
		"id": id,
		"model": model,
		"dataset": dataset,
		"round": round,
		"epoch": epoch,
	}

	var tx *xuper.Transaction

	tx, err = xchainClient.InvokeNativeContract(account, contractName, "ComputingShare", args)
	if err != nil {
		panic(err)
	}
	if tx != nil {
		fmt.Printf("InvokeComputingShare：%s\n", tx.ContractResponse.Body)
	}
}

/**
  ListenComputingShareEvent
  @Description: 链下监听computingShareEvent事件
  @return error
**/
func ListenComputingShareEvent() error{
	// 创建节点客户端。
	client, err := xuper.New("127.0.0.1:37101")
	if err != nil {
		return err
	}

	// 监听时间，返回 Watcher，通过 Watcher 中的 channel 获取block。
	//过滤掉空数据
	//watcher, err := client.WatchBlockEvent(xuper.WithSkipEmplyTx())
	//通过事件名称去过滤
	watcher, err := client.WatchBlockEvent(xuper.WithEventName("computingShareEvent"))
	if err != nil {
		return err
	}

	go func() {
		for {
			b, ok := <-watcher.FilteredBlockChan
			if !ok {
				fmt.Println("watch block event channel closed.")
				return
			}
			//fmt.Printf("%+v\n", b)
			//QueryBlockByHeight(b.BlockHeight)
			if len(b.Txs) != 0 {
				QueryTxByID("computingShareEvent", b.Txs[0].Txid)
			}
		}
	}()

	time.Sleep(time.Second * 3)

	// 关闭监听。
	watcher.Close()
	client.Close()
	return nil
}

func InvokeComputingCallBack(id string, result string, asig string, pks string) {
	account := getAccount()

	contractName := Contract_Name
	xchainClient, err := xuper.New("127.0.0.1:37101")
	if err != nil {
		panic(err)
	}
	args := map[string]string{
		"id": id,
		"faderated_ai_result": result,
		"asig": asig,
		"pks": pks,
	}

	var tx *xuper.Transaction

	tx, err = xchainClient.InvokeNativeContract(account, contractName, "ComputingCallBack", args)
	if err != nil {
		panic(err)
	}
	if tx != nil {
		fmt.Printf("InvokeComputingCallBack：%s\n", tx.ContractResponse.Body)
	}
}

/**
  QueryTxByID
  @Description: 根据交易ID查询交易
  @param txID 交易ID
**/
func QueryTxByID(eventType string, txID string) {
	node := "127.0.0.1:37101"
	xclient, _ := xuper.New(node)

	output, _ := xclient.QueryTxByID(txID)
	outputExt := output.GetTxOutputsExt()

	event := strings.Split(string(outputExt[0].Value),"\n")
	if eventType == "queryEvent" {
		queryEvent := event[2][13:]
		metadata = abstractQueryEvent(queryEvent)
		//fmt.Println(metadata)
	}

	if eventType == "computingShareEvent" {
		computingShareEvent := event[1][41:]
		metadata, learning = abstractComputingEvent(computingShareEvent)
	}


}

/**
  abstractQueryEvent
  @Description: 从http response中抽取数据协同事件内容
  @param event
**/
func abstractQueryEvent(event string) Metadata{
	var e DataEvent
	dec := json.NewDecoder(strings.NewReader(event))
	for {
		if err := dec.Decode(&e); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(e.Metadata)
	}

	sDec, err := base64.StdEncoding.DecodeString(e.Metadata)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return Metadata{}
	}

	//fmt.Println(string(sDec))

	var meta Metadata
	meta_decode := json.NewDecoder(strings.NewReader(string(sDec)))
	for {
		if err := meta_decode.Decode(&meta); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(meta)
	}
	return meta
}

/**
  abstractComputingEvent
  @Description: 从http response中抽取计算协同事件内容
  @param event
  @return Metadata
  @return FederatedAIDemand
**/
func abstractComputingEvent(event string) (Metadata, FederatedAIDemand){
	var e LearningEvent
	dec := json.NewDecoder(strings.NewReader(event))
	for {
		if err := dec.Decode(&e); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(e.Metadata)
	}

	metadata, err := base64.StdEncoding.DecodeString(e.Metadata)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return Metadata{}, FederatedAIDemand{}
	}
	//fmt.Println(string(metadata))

	federate, err := base64.StdEncoding.DecodeString(e.FaderatedAIDemandByte)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return Metadata{}, FederatedAIDemand{}
	}
	//fmt.Println(string(federate))

	var meta Metadata
	meta_decode := json.NewDecoder(strings.NewReader(string(metadata)))
	for {
		if err := meta_decode.Decode(&meta); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(meta)
	}

	var demand FederatedAIDemand
	demand_decode := json.NewDecoder(strings.NewReader(string(federate)))
	for {
		if err := demand_decode.Decode(&demand); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(demand)
	}

	return meta, demand
}

/**
  GetVariable
  @Description: 获取全局变量
  @return Metadata
  @return FederatedAIDemand
**/
func GetVariable() (Metadata, FederatedAIDemand) {
	return metadata, learning
}

/**
  run
  @Description: 合约运行函数
**/
func run() {
	//CreateAccount()
	//CreateContractAccount()
	//xuperchain.DeployContract()
	//id := xuperchain.InvokeCreateCfa("xuperchain", "counter", "data", "127.0.0.1", "xuperchain","162accb12e079d4b805f65f7a773c5e10cf537fef5ff99fde901ef0b1c963af8")
	//xuperchain.InvokeQuery(id)
	//xuperchain.ListenQueryEvent()
	//xuperchain.InvokeQueryCallback(id, "91.2", "aaa", "bbb")
	//time.Sleep(time.Second * 3)
	//xuperchain.InvokeComputingShare(id)
	//xuperchain.ListenComputingShareEvent()
	//xuperchain.InvokeComputingCallBack(id)
}



