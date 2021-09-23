[TOC]

#mac搭建FISCOBCOS
在mac上安装FISCO BCOS

#安装依赖
```
sudo apt install -y openssl curl
```

#创建操作目录, 下载安装脚本(需要翻墙)
```
# 创建操作目录
cd ~ && mkdir -p fisco && cd fisco
# 下载脚本
curl -#LO https://github.com/FISCO-BCOS/FISCO-BCOS/releases/download/v2.8.0/build_chain.sh && chmod u+x build_chain.sh
```
#搭建单群组4节点联盟链
请确保机器的30300~30303，20200~20203，8545~8548端口没有被占用
``` 
bash build_chain.sh -l 127.0.0.1:4 -p 30300,20200,8545
```
命令执行成功会输出All completed，如下所示。如果执行出错，请检查nodes/build.log文件中的错误信息。
``` 
Checking fisco-bcos binary...
Binary check passed.
==============================================================
Generating CA key...
==============================================================
Generating keys ...
Processing IP:127.0.0.1 Total:4 Agency:agency Groups:1
==============================================================
Generating configurations...
Processing IP:127.0.0.1 Total:4 Agency:agency Groups:1
==============================================================
[INFO] Execute the download_console.sh script in directory named by IP to get FISCO-BCOS console.
e.g.  bash /home/ubuntu/fisco/nodes/127.0.0.1/download_console.sh
==============================================================
[INFO] FISCO-BCOS Path   : bin/fisco-bcos
[INFO] Start Port        : 30300 20200 8545
[INFO] Server IP         : 127.0.0.1:4
[INFO] Output Dir        : /home/ubuntu/fisco/nodes
[INFO] CA Key Path       : /home/ubuntu/fisco/nodes/cert/ca.key
==============================================================
[INFO] All completed. Files in /home/ubuntu/fisco/nodes
```
#启动FISCO BCOS链
启动所有节点
``` 
bash nodes/127.0.0.1/start_all.sh
```
启动成功会输出类似下面内容的响应。否则请使用netstat -an | grep tcp检查机器的30300~30303，20200~20203，8545~8548端口是否被占用。
``` 
try to start node0
try to start node1
try to start node2
try to start node3
 node1 start successfully
 node2 start successfully
 node0 start successfully
 node3 start successfully
```
退出所有节点
``` 
bash nodes/127.0.0.1/stop_all.sh
```
# 检查进程
``` 
ps -ef | grep -v grep | grep fisco-bcos
```
正常情况会有类似下面的输出； 如果进程数不为4，则进程没有启动（一般是端口被占用导致的）
```
fisco       5453     1  1 17:11 pts/0    00:00:02 /home/ubuntu/fisco/nodes/127.0.0.1/node0/../fisco-bcos -c config.ini
fisco       5459     1  1 17:11 pts/0    00:00:02 /home/ubuntu/fisco/nodes/127.0.0.1/node1/../fisco-bcos -c config.ini
fisco       5464     1  1 17:11 pts/0    00:00:02 /home/ubuntu/fisco/nodes/127.0.0.1/node2/../fisco-bcos -c config.ini
fisco       5476     1  1 17:11 pts/0    00:00:02 /home/ubuntu/fisco/nodes/127.0.0.1/node3/../fisco-bcos -c config.ini
```
# 检查日志输出
如下，查看节点node0链接的节点数
```
tail -f nodes/127.0.0.1/node0/log/log*  | grep connected
```
正常情况会不停地输出连接信息，从输出可以看出node0与另外3个节点有连接。
```
info|2019-01-21 17:30:58.316769| [P2P][Service] heartBeat,connected count=3
info|2019-01-21 17:31:08.316922| [P2P][Service] heartBeat,connected count=3
info|2019-01-21 17:31:18.317105| [P2P][Service] heartBeat,connected count=3
```
执行下面指令，检查是否在共识
```
tail -f nodes/127.0.0.1/node0/log/log*  | grep +++
```
正常情况会不停输出++++Generating seal，表示共识正常。
```
info|2020-12-22 17:24:43.729402|[g:1][CONSENSUS][SEALER]++++++++++++++++ Generating seal on,blkNum=1,tx=0,nodeIdx=1,hash=2e133146...
info|2020-12-22 17:24:47.740603|[g:1][CONSENSUS][SEALER]++++++++++++++++ Generating seal on,blkNum=1,tx=0,nodeIdx=1,hash=eb199760...
```
# 配置及使用控制台
安装java
# 获取控制台并回到fisco目录
``` 
cd ~/fisco && curl -LO https://github.com/FISCO-BCOS/console/releases/download/v2.8.0/download_console.sh && bash download_console.sh
```
# 拷贝控制台配置文件
``` 
cp -n console/conf/config-example.toml console/conf/config.toml
```
# 配置控制台证书
``` 
cp -r nodes/127.0.0.1/sdk/* console/conf/
```
# 启动并使用控制台
``` 
cd ~/fisco/console && bash start.sh
```
输出下述信息表明启动成功，否则请检查conf/config.toml中节点端口配置是否正确
``` 
=============================================================================================
Welcome to FISCO BCOS console(2.6.0)！
Type 'help' or 'h' for help. Type 'quit' or 'q' to quit console.
 ________  ______   ______    ______    ______         _______    ______    ______    ______
|        \|      \ /      \  /      \  /      \       |       \  /      \  /      \  /      \
| $$$$$$$$ \$$$$$$|  $$$$$$\|  $$$$$$\|  $$$$$$\      | $$$$$$$\|  $$$$$$\|  $$$$$$\|  $$$$$$\
| $$__      | $$  | $$___\$$| $$   \$$| $$  | $$      | $$__/ $$| $$   \$$| $$  | $$| $$___\$$
| $$  \     | $$   \$$    \ | $$      | $$  | $$      | $$    $$| $$      | $$  | $$ \$$    \
| $$$$$     | $$   _\$$$$$$\| $$   __ | $$  | $$      | $$$$$$$\| $$   __ | $$  | $$ _\$$$$$$\
| $$       _| $$_ |  \__| $$| $$__/  \| $$__/ $$      | $$__/ $$| $$__/  \| $$__/ $$|  \__| $$
| $$      |   $$ \ \$$    $$ \$$    $$ \$$    $$      | $$    $$ \$$    $$ \$$    $$ \$$    $$
 \$$       \$$$$$$  \$$$$$$   \$$$$$$   \$$$$$$        \$$$$$$$   \$$$$$$   \$$$$$$   \$$$$$$

=============================================================================================
```
# 用控制台获取信息
``` 
# 获取客户端版本信息
[group:1]> getNodeVersion
ClientVersion{
    version='2.8.0',
    supportedVersion='2.8.0',
    chainId='1',
    buildTime='20210830 12:41:55',
    buildType='Darwin/appleclang/RelWithDebInfo',
    gitBranch='HEAD',
    gitCommitHash='30fb38ac5692468058abf6aa12869d2ae776c275'
}
#获取节点版本信息
[group:1]> getPeers
[
    PeerInfo{
        nodeID='d25afd558966bdb1ab154931eac7c807662cb0ff854f9f5b75b1ce2dcf7b2db4320d70be992b4341974576166c3ab65ecfb924a6b752e6ecdf5d485664831425',
        iPAndPort='127.0.0.1:58960',
        node='node3',
        agency='agency',
        topic='[

        ]'
    },
    PeerInfo{
        nodeID='fd1bfdb830a1b40357458340c88e6af98e14e8ab945524155d0d74b15c89c31125c47052c16769c8a728696f5ddc17d894138f095d18775b3088eeb4b57dfc15',
        iPAndPort='127.0.0.1:58980',
        node='node2',
        agency='agency',
        topic='[

        ]'
    },
    PeerInfo{
        nodeID='1659cc41593af97cb5289ef814d6c265013251017821e48bb1d2b3136fceb7a2aac7e2c31e5742ce08cdf72dd5fa6f70182bade52cc66330f45e12704e4fdfe9',
        iPAndPort='127.0.0.1:58964',
        node='node1',
        agency='agency',
        topic='[
            _block_notify_1
        ]'
    }
]

[group:1]>
```
# 编写HelloWorld合约
HelloWorld合约提供两个接口，分别是get()和set()，用于获取/设置合约变量name。合约内容如下:
``` 
pragma solidity ^0.4.24;

contract HelloWorld {
    string name;

    function HelloWorld() {
        name = "Hello, World!";
    }

    function get()constant returns(string) {
        return name;
    }

    function set(string n) {
        name = n;
    }
}
```
# 部署HelloWorld合约
为了方便用户快速体验，HelloWorld合约已经内置于控制台中，位于控制台目录下contracts/solidity/HelloWorld.sol（以后编译的合约都要放在该目录下），参考下面命令部署即可。
``` 
# 在控制台输入以下指令 部署成功则返回合约地址
[group:1]> deploy HelloWorld
transaction hash: 0x3891b8b625216486c56f6819ff99527732a8cac1fe644b3c6d539796f974e723
contract address: 0x600a6a6826eef21b88a60009edbd37480610f3ec
currentAccount: 0xb1e35d059405f1369115bf6132c46abc8f2ff048
```
# 调用HelloWorld合约
``` 
# 查看当前块高
[group:1]> getBlockNumber
1

# 调用get接口获取name变量 此处的合约地址是deploy指令返回的地址
[group:1]> call HelloWorld 0x600a6a6826eef21b88a60009edbd37480610f3ec get
---------------------------------------------------------------------------------------------
Return code: 0
description: transaction executed successfully
Return message: Success
---------------------------------------------------------------------------------------------
Return value size:1
Return types: (STRING)
Return values:(Hello, World!)
---------------------------------------------------------------------------------------------

# 查看当前块高，块高不变，因为get接口不更改账本状态
[group:1]> getBlockNumber
1


# 调用set设置name
[group:1]> call HelloWorld 0x600a6a6826eef21b88a60009edbd37480610f3ec set "Hello, FISCO BCOS"
transaction hash: 0x030d05d33a450a37f3aa0448029c9a61f543d10dda8bde87de4ccc99378d40b8
---------------------------------------------------------------------------------------------
transaction status: 0x0
description: transaction executed successfully
---------------------------------------------------------------------------------------------
Receipt message: Success
Return message: Success
Return values:[]
---------------------------------------------------------------------------------------------
Event logs
Event: {}

# 再次查看当前块高，块高增加表示已出块，账本状态已更改
[group:1]> getBlockNumber
2

# 调用get接口获取name变量，检查设置是否生效
[group:1]> call HelloWorld 0x600a6a6826eef21b88a60009edbd37480610f3ec get
---------------------------------------------------------------------------------------------
Return code: 0
description: transaction executed successfully
Return message: Success
---------------------------------------------------------------------------------------------
Return value size:1
Return types: (STRING)
Return values:(Hello, FISCO BCOS)
---------------------------------------------------------------------------------------------
# 退出控制台
[group:1]> quit
```
# sdk连接
在使用go-sdk时，需要将`fisco/nodes/127.0.0.1/sdk`中的`ca.crt`/`sdk.crt`/`sdk.key`中的三个文件拷贝到项目文件中，如`oracle/fisco`中所示。
用`go run *.go`下载fisco bcos的sdk时可能会出现以下错误：
```
# github.com/ethereum/go-ethereum/crypto/secp256k1
vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/curve.go:42:10: fatal error: 'libsecp256k1/include/secp256k1.h' file not found
#include "libsecp256k1/include/secp256k1.h"
         ^~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
```
则执行下列命令可解决
``` 
go get github.com/nomad-software/vend
vend
go run *.go
```
# 合约编译
以`Store.sol`合约为例
``` 
# 拉取sdk代码并编译
git clone https://github.com/FISCO-BCOS/go-sdk.git
cd go-sdk
go build cmd/console.go
# 搭建FISCO BCOS 2.2以上版本节点，如上所述
# config.toml默认使用channel模式，请拷贝对应的SDK证书，即拷贝`fisco/nodes/127.0.0.1/sdk`中的`ca.crt`/`sdk.crt`/`sdk.key`中的三个证书到`oracle/fisco`目录下。
```
如果在编译过程中出现如下提示
``` 
go: github.com/FISCO-BCOS/crypto@v0.0.0-20200202032121-bd8ab0b5d4f1: missing go.sum entry; to add it:
        go mod download github.com/FISCO-BCOS/crypto
```
则执行如下命令
``` 
go mod int
go mod vendor
```
在`go-sdk`目录下安装对应版本的solc编译器，编译solidity合约
```
bash tools/download_solc.sh -v 0.4.25
```
构建go-sdk的代码生成工具`abigen`
```
# 下面指令在go-sdk目录下操作，编译生成abigen工具
go build ./cmd/abigen
```
执行命令后，检查根目录下是否存在`abigen`，并将准备的智能合约`Store.sol`放置在一个`go-sdk`的一个新的目录`store`下。
编译生成go文件，先利用solc将合约文件生成abi和bin文件，以前面所提供的Store.sol为例：
```
./solc-0.4.25 --bin --abi -o ./store ./store/Store.sol
```
Store.sol目录下会生成Store.bin和Store.abi。此时利用abigen工具将Store.bin和Store.abi转换成Store.go：
``` 
./abigen --bin ./store/Store.bin --abi ./store/Store.abi --pkg store --type Store --out ./store/Store.go
```
最后store目录下面存在以下文件：
``` 
Store.abi  Store.bin  Store.go  Store.sol
```
将生成的`Store.go`文件拷贝到`oracle/fisco`文件下，进行合约调用。

# 命令行生成账户
获取脚本
``` 
curl -#LO https://raw.githubusercontent.com/FISCO-BCOS/console/master/tools/get_account.sh && chmod u+x get_account.sh && bash get_account.sh -h
```
使用脚本生成PEM格式私钥
```
bash get_account.sh
```
执行上面的命令，可以得到类似下面的输出，包括账户地址和以账户地址为文件名的私钥PEM文件。
```
[INFO] Account Address   : 0x22a163da252e56d8c1befc5efea123991f761c36
[INFO] Private Key (pem) : accounts/0x22a163da252e56d8c1befc5efea123991f761c36.pem
[INFO] Public  Key (pem) : accounts/0x22a163da252e56d8c1befc5efea123991f761c36.public.pem
```
指定PEM私钥文件计算账户地址
```
bash get_account.sh -k accounts/0x22a163da252e56d8c1befc5efea123991f761c36.pem
```
执行上面的命令，结果如下
```
[INFO] Account Address   : 0x22a163da252e56d8c1befc5efea123991f761c36
```
