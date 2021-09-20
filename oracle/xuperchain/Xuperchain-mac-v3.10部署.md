# 资源网站

```
https://github.com/xuperchain/contract-sdk-go
https://github.com/xuperchain/xuper-sdk-go
https://github.com/xuperchain/xuperchain
https://xuperchain.readthedocs.io/zh/latest/advanced_usage/create_contracts.html
```

#Xuperchain 部署

```
#下载源码
git clone https://github.com/xuperchain/xuperchain.git
cd xuperchain
#切换到3.10.3版本
git checkout v3.10.3
#编译
make
```

在编译的时候可能碰到以下问题：

>github.com/xuperchain/xuperchain/core/cmd/xdev/internal/jstest/core/cmd/xdev/internal/jstest/runner.go:165:37: cannot use testDeps{} (type testDeps) as type testing.testDeps in argument to testing.MainStart:
>	testDeps does not implement testing.testDeps (missing SetPanicOnExit0 method)
>make: *** [build] Error 2

这是源码中没写`SetPanicOnExit0 `，解决方法如下：

```
#在源码的core/cmd/xdev/internal/jstest/deps.go
func (t testDeps) ImportPath() string                          { return "" }
func (t testDeps) StartTestLog(io.Writer)                      {}
func (t testDeps) StopTestLog() error                          { return nil }
#加入下面这行代码
func (t testDeps) SetPanicOnExit0(bool)                        {}
```

然后重新编译`make`，此时`xuperchain`目录下出现`output`文件，里面存放着编译好的可执行文件。

#创建区块链网络

此时，可以创建区块链了。

```
cd output
#创建区块链
./xchain-cli createChain
```

执行完成后，在`output`目录下的`data/blockchain/`中出现一个文件夹`xuper`，该文件夹即表示创建的区块链名称。如果网络出现错误，可以将其删除，然后在`output`目录下重新执行`./xchain-cli createChain`。

# 运行全节点

在`output`目录下运行如下命令可启动全节点。

```
nohup ./xchain &
```

在output目录下运行`cat nohup.out`可查看日志。如果sdk代码跑不通，或者是命令行无法执行xuperchain的命令，先查看日志判断网络是否宕机，随后进行下一步的判断。

如果一不小心运行两次，导致端口占用`127.0.0.1:37101/37200 has already been used`，可以执行：

```
#查询占用该端口号的进程
lsof -i:37200 
#杀死进程(pid为上个命令查出来的结果)
kill pid
```

如果想修改配置文件，但是不想关闭网络，也可以关闭全节点->修改配置文件->启动全节点。

# SDK-Go

通过指定服务器地址和端口号即可连接区块链，注意：区块链部署在本机不用配置文件，即代码运行目录（注意是项目目录，不是xuperchain所在的目录）下不需要出现`conf/sdk.yaml`，否则会出现以下报错：

```
#合约账户和账户不匹配
panic: Failed to post tx: TX_SIGN_ERROR
panic: Failed to post tx: RWACL_INVALID_ERROR
panic: EndorserCall PreExecWithFee failed: rpc error: code = Unknown desc = identity check failed
```

```
#单机部署不需要这个配置文件
# endorseService Info
# testNet addrs
endorseServiceHost: "39.156.69.83:37100"
complianceCheck:
# 是否需要进行合规性背书
isNeedComplianceCheck: false
# 是否需要支付合规性背书费用
isNeedComplianceCheckFee: false
# 合规性背书费用
complianceCheckEndorseServiceFee: 400
# 支付合规性背书费用的收款地址
complianceCheckEndorseServiceFeeAddr: aB2hpHnTBDxko3UoP2BpBZRujwhdcAFoT
# 如果通过合规性检查，签发认证签名的地址
complianceCheckEndorseServiceAddr: jknGxa6eyum1JrATWvSJKW3thJ9GKHA9n
#创建平行链所需要的最低费用
minNewChainAmount: "100"
crypto: "xchain"
txVersion: 1
```

在涉及与区块链交互的操作时，需要做一些前提准备：

在`xuperchain/conf`（即xuperchain所在目录）修改`xchain.yaml`的`xendorser`的`enable`状态设置为`true`，否则会出现以下问题：

> PreExecWithSelected UTXO failed: rpc error: code = Unimplemented sesc = unkonwn service pb.Xchain
>
> Panic: runtime error: invalid memory address or nil pointer dereference

同时，为了运行go合约和监听事件，需要将`xchain.yaml`中的`native`和`event`的`enable`都改为`true`。

在使用命令行部署合约时，需要在项目目录中执行`go build -o counter.go`命令，然后将生成的可执行文件`counter`复制到`output`目录下，随后执行以下命令：

```
./xchain-cli native deploy --account XC1234567890111111@xuper --fee 15587517 --runtime go -a '{"creator":"XC1234567890111111@xuper"}'   --cname golangcounter
```

注意：一定要保证合约账户是有余额的，否则会出现`NOT ENGOUGH`的报错。可以运行命令行或者是示例代码进行转账，具体的内容已在`oracle/main.go`中提示，在此不赘述。

部署和调用合约时，可能出现以下错误：
>desc = run wasm2c, ..., bad magic value

现在超级链已经不支持将编译后的go合约转换为wasm格式了，所以不要去修改相应的`conf/xchain.yaml`文件。

>Panic: PreExecWithSelectUTXO failed: rpc error: code = Unknown desc = contract type native not found

此时，首先检查`output`目录下的`nohup.out`日志，查看网络的状态是否完好，命令是否执行成功。如果网络宕机，则删掉网络对应的文件夹，重新启动。

当账户和合约账户不匹配时，可能出现以下两个错误：

> panic: Failed to post tx: TX_SIGN_ERROR
>
> panic: Failed to post tx: RWACL_INVALID_ERROR

需要保证合约账户是该账户生成的即可。

由于xuperchain版本变更的问题，在使用`go mod`可能出现以下问题：

> SECURITY ERROR
> This download does NOT match the one reported by the checksum server.
> The bits may have been replaced on the origin server, or an attacker may
> have intercepted the download attempt.

此时，先暂时关闭`go mod`的验证，待到下载完成之后，再开启验证。

```
go env -w GOSUMDB=off
go mod tidy
go env -w GOSUMDB=GOSUMDB="sum.golang.org"
```

当结果中出现`no config file in ./conf/sdk.yaml, use default config: &{10.144.94.18:8848 {false false 10 XBbhR82cB6PvaLJs3D4uB9f12bhmKkHeX TYyA3y8wdFZyzExtcbRNVd7ZZ2XXcfjdw} 100 xchain 0}`的日志提示时，不用理会，直接跳过即可。