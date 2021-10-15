package fisco

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"strconv"
	"strings"
)

const (
	contractAddr = "0x89a266754e2ecf22f8aa8db6418c46ca5a80669f"
)

type ComputingShareEvent struct {
	id *big.Int
	dataAbstract string
	model string
	dataset string
	round string
	epoch string
}

func getClient() *client.Client{
	configs, err := conf.ParseConfigFile("fisco/config.toml")
	if err != nil {
		log.Fatalf("ParseConfigFile failed, err: %v", err)
	}
	client, err := client.Dial(&configs[0])
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return client
}

func getSession() ComputingSession{
	client := getClient()
	//部署合约
	//input := "Store deployment 1.0"
	//address, tx, instance, err := DeployStore(client.GetTransactOpts(), client, input)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("contract address: ", address.Hex()) // the address should be saved, will use in next example
	//fmt.Println("transaction hash: ", tx.Hash().Hex())

	// 加载合约
	contractAddress := common.HexToAddress(contractAddr)
	instance, err := NewComputing(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("================================")
	computingSession := ComputingSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	return computingSession
}

func CreateMetadata(abstract string, source string, datatype *big.Int) *big.Int{
	session := getSession()
	_, receipt, err := session.Createmetadata(abstract, source, datatype)
	if err != nil {
		log.Fatal(err)
	}
	output := receipt.GetOutput()
	count, err := strconv.ParseInt(output[2:], 16, 32)
	if err != nil {
		log.Fatal(err)
	}
	return big.NewInt(count)
}

func GetMetadata(id *big.Int) {
	session := getSession()
	_, receipt, err := session.Getmetadata(id)
	if err != nil {
		log.Fatal("Session error: ", err)
	}
	outputHex := receipt.GetOutput()
	output, err := hex.DecodeString(outputHex[2:])
	if err != nil {
		log.Fatal("Decode error: ", err)
	}
	fmt.Println("GetMetadata: ",  string(output))
}

func ComputingShare(id *big.Int, model string, dataset string, round string, epoch string) {
	session := getSession()
	_, receipt, err := session.Computingshare(id, model, dataset, round, epoch)
	if err != nil {
		log.Fatal("Session error: ", err)
	}
	outputHex := receipt.GetOutput()
	output, err := hex.DecodeString(outputHex[2:])
	if err != nil {
		log.Fatal("Decode error: ", err)
	}
	fmt.Println("ComputingShare: ",  string(output))
}

func ComputingShareCallback(id *big.Int, paraAddr string, paraAbstract string, corectRate string) {
	session := getSession()
	_, receipt, err := session.Computingsharecallback(id, paraAddr, paraAbstract, corectRate)
	if err != nil {
		log.Fatal("Session error: ", err)
	}
	outputHex := receipt.GetOutput()
	output, err := hex.DecodeString(outputHex[2:])
	if err != nil {
		log.Fatal("Decode error: ", err)
	}
	fmt.Println("ComputingShareCallback: ",  string(output))
}

func GetResult(id *big.Int) {
	session := getSession()
	_, receipt, err := session.Getresult(id)
	if err != nil {
		log.Fatal("Session error: ", err)
	}
	outputHex := receipt.GetOutput()
	output, err := hex.DecodeString(outputHex[2:])
	if err != nil {
		log.Fatal("Decode error: ", err)
	}
	fmt.Println("GetResult: ",  string(output))
}

func DataShare(id *big.Int) {
	session := getSession()
	_, receipt, err := session.Datashare(id)
	if err != nil {
		log.Fatal("Session error: ", err)
	}
	outputHex := receipt.GetOutput()
	output, err := hex.DecodeString(outputHex[2:])
	if err != nil {
		log.Fatal("Decode error: ", err)
	}
	fmt.Println("DataShare: ",  string(output))
}

func DataShareCallback(id *big.Int, result string) {
	session := getSession()
	_, receipt, err := session.Datasharecallback(id, result)
	if err != nil {
		log.Fatal("Session error: ", err)
	}
	outputHex := receipt.GetOutput()
	output, err := hex.DecodeString(outputHex[2:])
	if err != nil {
		log.Fatal("Decode error: ", err)
	}
	fmt.Println("DataShareCallback: ",  string(output))
}

func ListenComputingEvent() {
	client := getClient()
	contractAddress := common.HexToAddress(contractAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}

func GetComputingEvent() {
	client := getClient()

	contractAddress := common.HexToAddress(contractAddr)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(1),
		ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal("Filterlogs error: ", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(ComputingABI))
	if err != nil {
		log.Fatal("Abi error: ", err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		event := ComputingShareEvent{}
		err := contractAbi.Unpack(&event, "computingShareEvent", vLog.Data)
		if err != nil {
			log.Fatal("Abiunpack error:", err)
		}

		fmt.Println(event) // bar

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}
}

func run() {
	var id *big.Int
	id = CreateMetadata("abstract", "source", big.NewInt(1))
	GetMetadata(id)
	ComputingShare(id, "model", "dataset", "round", "epoch")
	ComputingShareCallback(id, "paramAddr", "paramAbstract", "88.7%")
	GetResult(id)
	// 需要将datatype改为2，然后才能进行datashare
	id = CreateMetadata("abstract", "source", big.NewInt(2))
	DataShare(id)
	DataShareCallback(id, "this is datashare result")
	GetResult(id)
}
