package fiscosdk

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
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


