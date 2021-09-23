package fisco

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"time"
)

/**
  CreateAccount
  @Description: 创建外部账户
**/
func CreateAccount() {
	//生成随机私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	//将私钥转换为字节类型
	privateKeyBytes := crypto.FromECDSA(privateKey)
	//将字节类型的私钥转换为十六进制编码，在十六进制编码之后删除“0x”
	fmt.Println("private key: ", hexutil.Encode(privateKeyBytes)[2:]) // privateKey in hex without "0x"
	//根据私钥派生出公钥
	publicKey := privateKey.Public()
	//将公钥转换为十六进制编码，并剥离了0x和前2个字符04，它始终是EC前缀，不是必需的
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//将公钥转换为字节类型
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publick key: ", hexutil.Encode(publicKeyBytes)[4:])  // publicKey in hex without "0x"
	//根据公钥生成地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address)  // account address
}

/**
  NativeContract
  @Description:
	部署、调用、查询合约
	利用abigen工具生成`Store.go`，具体参考markdown文档的合约编译部分
	将`go-sdk/store/Store.go`拷贝到项目目录下
	在当前项目目录下，使用`get_account.sh`生成密钥，具体参考markdown文档的命令行生成账户部分
	将config.toml的KeyFile属性替换成`get_account.sh`生成的`pem`文件地址，如`KeyFile="fisco/accounts/0x22a163da252e56d8c1befc5efea123991f761c36.pem.pem"`
	需要注意的是，项目的执行目录在orale/main.go下，所以调用fisco目录下的文件路径前缀都要加`fisco/`
**/
func NativeContract() {
	configs, err := conf.ParseConfigFile("fisco/config.toml")
	if err != nil {
		log.Fatalf("ParseConfigFile failed, err: %v", err)
	}
	client, err := client.Dial(&configs[0])
	if err != nil {
		log.Fatal(err)
	}
	//部署合约
	//input := "Store deployment 1.0"
	//address, tx, instance, err := DeployStore(client.GetTransactOpts(), client, input)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("contract address: ", address.Hex()) // the address should be saved, will use in next example
	//fmt.Println("transaction hash: ", tx.Hash().Hex())

	// 加载合约
	contractAddress := common.HexToAddress("0x0EFAB177fdF1270a03C3EA3604169baFCb4Bd661")
	instance, err := NewStore(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("================================")
	storeSession := StoreSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	version, err := storeSession.Version()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("version :", version) // "Store deployment 1.0"

	// 调用合约
	fmt.Println("================================")
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	tx, receipt, err := storeSession.SetItem(key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	// 查询合约
	result, err := storeSession.Items(key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("get item: " + string(result[:])) // "bar"
}

func ListenEvent() {
	configs, err := conf.ParseConfigFile("fisco/config.toml")
	if err != nil {
		log.Fatalf("ParseConfigFile failed, err: %v", err)
	}
	client, err := client.Dial(&configs[0])
	if err != nil {
		log.Fatal(err)
	}

	client.SubscribeBlockNumberNotify(func(i int64) {
		fmt.Println("block notify")

		txByte, err := client.GetTransactionByBlockNumberAndIndex(context.Background(), i, 0 )
		//blockByte, err := client.GetBlockByNumber(context.Background(), i, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("aaaa", string(txByte))

	})

	NativeContract()

	time.Sleep(1 * time.Second)
}


