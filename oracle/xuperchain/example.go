package xuperchain
//
//import (
//	"fmt"
//	"github.com/xuperchain/xuper-sdk-go/v2/account"
//	"github.com/xuperchain/xuper-sdk-go/v2/xuper"
//	"time"
//)
//
//const (
//	Mnemonic = "钢 车 稀 叔 送 扫 永 确 描 当 矮 北"
//	Address = "jPSFQjS6jwMi8tU8F2EFLAiWXpMGno4FU"
//	Contract_Addr = "XC1234567890111112@xuper"
//)
//
//
///**
//  CreateAccount
//  @Description: 创建账户
//**/
//func CreateAccount() {
//	var acc *account.Account
//	var err error
//	// 测试创建账户
//	acc, err = account.CreateAccount(1, 1)
//	if err != nil {
//		fmt.Printf("create account error: %v\n", err)
//	}
//	fmt.Println(acc)
//	fmt.Println(acc.Mnemonic)
//
//	// 测试恢复账户
//	acc, err = account.RetrieveAccount("钢 车 稀 叔 送 扫 永 确 描 当 矮 北}", 1)
//	if err != nil {
//		fmt.Printf("retrieveAccount err: %v\n", err)
//		return
//	}
//	fmt.Printf("RetrieveAccount: to %v\n", acc)
//
//	// 创建账户并存储到文件中
//	acc, err = account.CreateAndSaveAccountToFile("./keys", "123", 1, 1)
//	if err != nil {
//		fmt.Printf("createAndSaveAccountToFile err: %v\n", err)
//	}
//	fmt.Printf("CreateAndSaveAccountToFile: %v\n", acc)
//
//	// 从文件中恢复账户
//	acc, err = account.GetAccountFromFile("keys/", "123")
//	if err != nil {
//		fmt.Printf("getAccountFromFile err: %v\n", err)
//	}
//	fmt.Printf("getAccountFromFile: %v\n", acc)
//	return
//}
//
///**
//  CreateContractAccount
//  @Description:
//	创建合约账户
//	命令行创建合约账号：./xchain-cli account new --account 1234567890111111 --fee 1000
//	生成的账户：XC1111111111111111@xuper
//	在创建合约账号之前，需要给账户转钱，才能支付手续费（记得修改地址）
//	命令行运行：./xchain-cli transfer --to jPSFQjS6jwMi8tU8F2EFLAiWXpMGno4FU --amount 100000000 --keys data/keys/ -H 127.0.0.1:37101
//	XC1111111111111111@xuper对应的account的助记词为巴 碱 仿 幼 浸 知 讼 朋 蒸 邵 雄 即
//**/
//func CreateContractAccount() {
//	// 从文件中恢复账户
//	acc, err := account.RetrieveAccount(Mnemonic, 1)
//	if err != nil {
//		fmt.Printf("retrieveAccount err: %v\n", err)
//		return
//	}
//	fmt.Printf("RetrieveAccount: to %v\n", acc)
//
//	// 创建一个合约账户
//	// 合约账户是由 (XC + 16个数字 + @xuper) 组成, 比如 "XC1234567890123456@xuper"
//	contractAccount := Contract_Addr
//
//	xchainClient, err := xuper.New("127.0.0.1:37101")
//	tx, err := xchainClient.CreateContractAccount(acc, contractAccount)
//	if err != nil {
//		fmt.Printf("createContractAccount err:%s\n", err.Error())
//	}
//	fmt.Println(tx.Tx.Txid)
//	return
//}
//
///**
//  akTransfer
//  @Description: 普通账户转账
//  @param to 目的账户
//  @param amount 金额
//**/
//func akTransfer(to *account.Account, amount string) {
//	// 创建或者使用已有账户，此处为使用已有账户。
//	// 恢复账户
//	//nuSMPvo6UUoTaT8mMQmHbfiRbJNbAymGh
//	//./xchain-cli transfer --to nuSMPvo6UUoTaT8mMQmHbfiRbJNbAymGh --amount 1000000000000 --keys data/keys/ -H 127.0.0.1:37101
//	from, err := account.RetrieveAccount(Mnemonic, 1)
//	if err != nil {
//		fmt.Printf("retrieveAccount err: %v\n", err)
//		return
//	}
//	fmt.Printf("RetrieveAccount: to %v\n", from)
//
//	//to, err := account.CreateAccount(1, 1)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//fmt.Println(to.Address)
//	//fmt.Println(to.Mnemonic)
//
//	// 节点地址。
//	node := "127.0.0.1:37101"
//
//	// 创建节点客户端。
//	xclient, _ := xuper.New(node)
//
//	// 转账前查看两个地址余额。
//	fmt.Println(xclient.QueryBalance(from.Address))
//	fmt.Println(xclient.QueryBalance(to.Address))
//
//	tx, err := xclient.Transfer(from, to.Address, amount)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("%x\n", tx.Tx.Txid)
//
//	// 转账后查看两个地址余额。
//	fmt.Println(xclient.QueryBalance(from.Address))
//	fmt.Println(xclient.QueryBalance(to.Address))
//}
//
///**
//  contractAccountTransfer
//  @Description:
//	合约账户转账
//	命令行给合约账户转账：./xchain-cli transfer --to XC1234567890111112@xuper --amount 1000000000000
//**/
//func contractAccountTransfer() {
//	// 创建或者使用已有账户，此处为新创建一个账户。
//	me, err := account.RetrieveAccount(Mnemonic, 1)
//	if err != nil {
//		fmt.Printf("retrieveAccount err: %v\n", err)
//		return
//	}
//	fmt.Printf("RetrieveAccount: to %v\n", me)
//
//	//me, err := account.CreateAccount(1, 1)
//	//if err != nil {
//	//	panic(err)
//	//}
//
//	akTransfer(me, "100")
//
//	// XC1234567890111111@xuper 为合约账户，如果没有合约账户需要先创建合约账户。
//	//该合约账户必须是账户生成的，否则会报错
//	me.SetContractAccount(Contract_Addr)
//	fmt.Println(me.Address)
//	fmt.Println(me.Mnemonic)
//	fmt.Println(me.GetContractAccount())
//	fmt.Println(me.GetAuthRequire())
//
//	to, err := account.CreateAccount(1, 1)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(to.Address)
//	fmt.Println(to.Mnemonic)
//
//	// 节点地址。
//	node := "127.0.0.1:37101"
//	xclient, _ := xuper.New(node)
//
//	// 转账前查看两个地址余额。
//	fmt.Println(xclient.QueryBalance(me.GetContractAccount()))
//	fmt.Println(xclient.QueryBalance(to.Address))
//
//	tx, err := xclient.Transfer(me, to.Address, "10")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("%x\n", tx.Tx.Txid)
//
//	// 转账后查看两个地址余额。
//	fmt.Println(xclient.QueryBalance(me.GetContractAccount())) // 转账时使用的是合约账户，因此查询余额时也是合约账户。
//	fmt.Println(xclient.QueryBalance(to.Address))
//}
//
///**
//  QueryBlockByHeight
//  @Description: 根据区块高度查询区块信息
//  @param height 区块高度
//**/
//func QueryBlockByHeight(height int64) {
//	// 示例代码省略了 err 的检查。
//	node := "127.0.0.1:37101"
//	xclient, _ := xuper.New(node)
//	blockResult, _ := xclient.QueryBlockByHeight(height)
//	if blockResult.GetHeader().GetError() != 0 {
//		// 处理错误。
//	} else {
//		// 处理区块数据。
//		block := blockResult.GetBlock()
//		fmt.Println(block.GetBlockid())
//		fmt.Println(block.GetHeight())
//		fmt.Println(block.GetTxCount())
//		fmt.Println(block.Transactions)
//	}
//}
//
///**
//  QueryTxByID
//  @Description: 根据交易ID查询交易
//  @param txID 交易ID
//**/
//func QueryTxByID(txID string) {
//	node := "127.0.0.1:37101"
//	xclient, _ := xuper.New(node)
//
//	output, _ := xclient.QueryTxByID(txID)
//	outputExt := output.GetTxOutputsExt()
//	// 输出合约事件内容
//	fmt.Println(string(outputExt[0].Key))
//	fmt.Println(string(outputExt[0].Value))
//}
//
///**
//  QueryBalance
//  @Description: 查询账户余额
//**/
//func QueryBalance() {
//	// 示例代码省略了 err 的检查。
//	node := "127.0.0.1:37101"
//	xclient, _ := xuper.New(node)
//
//	// 查询账户余额，默认 xuper 链。
//	bal, _ := xclient.QueryBalance(Contract_Addr)
//	fmt.Println(bal)
//
//	//// 查询账户余额，在 hello 链。
//	//bal, _ = xclient.QueryBalance("nuSMPvo6UUoTaT8mMQmHbfiRbJNbAymGh", xuper.WithQueryBcname("hello"))
//	//fmt.Println(bal)
//	//
//	//// 查询账户余额详细数据
//	//balDetails, _ := xclient.QueryBalanceDetail("nuSMPvo6UUoTaT8mMQmHbfiRbJNbAymGh")
//	//for _, bd := range balDetails {
//	//	fmt.Println(bd.Balance)
//	//	fmt.Println(bd.IsFrozen)
//	//}
//
//}
//
///**
//  NativeContract
//  @Description:
//	部署合约、调用合约、查询合约
//	命令行部署合约：在output目录下运行`./xchain-cli wasm deploy --account XC1234567890111111@xuper --cname gocounter -a '{"creator":"xchain"}' ./counter --runtime go`
//**/
//func NativeContract() {
//	//codePath := "example/contract/counter" // 编译好的二进制文件 go build -o counter
//	//code, err := ioutil.ReadFile(codePath)
//	//if err != nil {
//	//	panic(err)
//	//}
//
//	account, err := account.RetrieveAccount(Mnemonic, 1)
//	if err != nil {
//		fmt.Printf("retrieveAccount err: %v\n", err)
//		return
//	}
//	fmt.Printf("retrieveAccount address: %v\n", account.Address)
//	contractAccount := Contract_Addr
//	contractName := "SDKNativeCount2"
//	err = account.SetContractAccount(contractAccount)
//	if err != nil {
//		panic(err)
//	}
//
//	xchainClient, err := xuper.New("127.0.0.1:37101")
//	if err != nil {
//		panic(err)
//	}
//	args := map[string]string{
//		"creator": "xuperchain",
//		"key":     "counter",
//	}
//	var tx *xuper.Transaction
//	//tx, err = xchainClient.DeployNativeGoContract(account, contractName, code, args)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//fmt.Printf("Deploy Native Go Contract Success! %x\n", tx.Tx.Txid)
//
//	tx, err = xchainClient.InvokeNativeContract(account, contractName, "increase", args)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Invoke Native Go Contract Success! %x\n", tx.Tx.Txid)
//
//	tx, err = xchainClient.QueryNativeContract(account, contractName, "get", args)
//	if err != nil {
//		panic(err)
//	}
//	if tx != nil {
//		fmt.Printf("查询结果：%s\n", tx.ContractResponse.Body)
//	}
//}
//
///**
//  ListenEvent
//  @Description: 监听合约事件
//  @return error
//**/
//func ListenEvent() error{
//	// 创建节点客户端。
//	client, err := xuper.New("127.0.0.1:37101")
//	if err != nil {
//		return err
//	}
//
//	// 监听时间，返回 Watcher，通过 Watcher 中的 channel 获取block。
//	//过滤掉空数据
//	//watcher, err := client.WatchBlockEvent(xuper.WithSkipEmplyTx())
//	//通过事件名称去过滤
//	watcher, err := client.WatchBlockEvent(xuper.WithEventName("increase"))
//	if err != nil {
//		return err
//	}
//
//	go func() {
//		for {
//			b, ok := <-watcher.FilteredBlockChan
//			if !ok {
//				fmt.Println("watch block event channel closed.")
//				return
//			}
//			//fmt.Printf("%+v\n", b)
//			//QueryBlockByHeight(b.BlockHeight)
//			if len(b.Txs) != 0 {
//				QueryTxByID(b.Txs[0].Txid)
//			}
//		}
//	}()
//
//	time.Sleep(time.Second * 3)
//
//	// 关闭监听。
//	watcher.Close()
//	client.Close()
//	return nil
//}
//
///**
//  run
//  @Description: 运行代码
//**/
//func run() {
//	//CreateAccount()
//	//CreateContractAccount()
//	//akTransfer()
//	//contractAccountTransfer()
//	//QueryBlockByHeight()
//	//QueryBalance()
//	NativeContract()
//	ListenEvent()
//}
//
//
