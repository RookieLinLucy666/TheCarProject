[TOC]



### 端口查看

```bash
netstat -lnp|grep 47101
kill -9 <pid>
```

### 生成链配置

```bash
./xchain-cli createChain
```

### 启动链

```bash
./xchain
```

### 链状态

```bash
./xchain-cli status -H 127.0.0.1:37101
```

### 创建账号

```bash
./xchain-cli account newkeys -f
./xchain-cli netURL gen
```

### 打印keys

```bash
head data/keys/* && echo
head data/netkeys/* && echo
```

### 获取netUrl

```bash
./xchain-cli netURL get -H 127.0.0.1:37101
```

### 创建合约账号

- **直接创建**

```bash
./xchain-cli account new --account 1111111111111111 --fee 1000
```

- **ACL创建**

```bash
{
    "module_name": "xkernel",
    "method_name": "NewAccount",
    "args" : {
        "account_name": "1111111111111111",
        "acl": "{\"pm\": {\"rule\": 1,\"acceptValue\": 1},\"aksWeight\": {\"dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN\": 1}}"
    }
}

./xchain-cli account new --desc acl.json
```

### 转账

```bash
./xchain-cli transfer --to XC1111111111111111@xuper --amount 1000000
```

### 余额

```bash
./xchain-cli account balance XC1111111111111111@xuper
./xchain-cli account balance
```

### 查看合约账户的ACL

```bash
./xchain-cli acl query --account XC1111111111111111@xuper
```

### 查看合约方法的ACL

```bash
./xchain-cli acl query --contract group_chain --method addNode
```

### 查看账号管理的合约

```bash
./xchain-cli account contracts --account XC1111111111111111@xuper
```

## 合约部署

```bash
# c语言合约
# 运行core/contractsdk/cpp/build.sh脚本 这里依赖了一个xdev
# xdev：go build -o xdev github.com/xuperchain/xuperchain/core/cmd/xdev
./xchain-cli wasm deploy --account XC1111111111111111@xuper --cname counter ../cc/counter.wasm -a '{"creator":"xchain"}' --fee 150000

#调用
./xchain-cli wasm invoke -a '{"key":"counter"}' --method increase counter --fee 150000
./xchain-cli wasm query -a '{"key":"counter"}' --method get counter

#go语言合约
#很卡，而且手续费极高，阿里云1核2G的编译不了，建议4G的机器
GOOS=js GOARCH=wasm go build
# 便捷方式
./xchain-cli wasm deploy --account XC1111111111111111@xuper --cname gocounter -a '{"creator":"xchain"}' ./counter --runtime go
```



## 合约升级

```bash
#修改配置：conf/xuper.yaml
wasm:
  enableUpgrade: true #允许合约升级

#编写新合约并编译

#更新
./xchain-cli wasm upgrade --account XC1111111111111111@xuper --cname counter ../cc/counter.wasm --fee 150000
```

## 群组

### Node

```bash
#list 该hello链允许的节点
./xchain-cli wasm query group_chain --method listNode -a '{"bcname":"hello"}'

#add 添加节点到hello链
./xchain-cli wasm invoke group_chain --method addNode -a '{"bcname":"hello", "ip":"/ip4/127.0.0.1/tcp/47101/p2p/QmVxeNubpg1ZQjQT8W5yZC9fD7ZB1ViArwvyGUB53sqf8e", "address":"dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN"}'

#ip为./xchain-cli netURL get 获取到的url，记得修改真实ip
#address为data/keys/addres的地址

#del 从hello链删除节点
./xchain-cli wasm invoke group_chain --method delNode -a '{"bcname":"hello","ip":"/ip4/127.0.0.1/tcp/47101/p2p/QmVxeNubpg1ZQjQT8W5yZC9fD7ZB1ViArwvyGUB53sqf8e"}'
```

### Chain

```bash
#list 在群组管理的链
./xchain-cli wasm query group_chain --method listChain

#add 添加链到群组
./xchain-cli wasm invoke group_chain --method addChain -a '{"bcname":"hello"}'

#del 从群组删除链
./xchain-cli wasm invoke group_chain --method delChain -a '{"bcname":"hello"}'
```

## 合约ACL

### acl格式说明

```bash
#a.创建合约账号,配置文件案例如下：
{
    "module_name": "xkernel",
    "method_name": "NewAccount",
    "args" : {
        "account_name": "1111111111111111",
        "acl": "{\"pm\": {\"rule\": 1,\"acceptValue\": 0.6},\"aksWeight\": {\"AK1\": 0.5,\"AK2\": 0.5}}"
    }
}

#b.重置合约账户的ACL,配置文件案例如下：
{
    "module_name": "xkernel",
    "method_name": "SetAccountAcl",
    "args" : {
        "account_name": "XC1111111111111111@xuper",
        "acl": "{\"pm\": {\"rule\": 1,\"acceptValue\": 0.6},\"aksWeight\": {\"AK1\": 0.5,\"AK2\": 0.5}}"
    }
}

#c.重置合约方法的ACL,配置文件案例如下：
{
    "module_name": "xkernel",
    "method_name": "SetMethodAcl",
    "args" : {
        "contract_name": "math",
        "method_name": "transfer",
        "acl": "{\"pm\": {\"rule\": 1,\"acceptValue\": 0.6},\"aksWeight\": {\"AK1\": 0.5,\"AK2\": 0.5}}"
    }
}
```

### 设置方法ACL步骤

```bash
#创建需要签名的合约方法文件
vi addnode.json
{
    "module_name": "xkernel",
    "method_name": "SetMethodAcl",
    "args" : {
        "contract_name": "group_chain",
        "method_name": "addNode",
        "acl": "{\"pm\": {\"rule\": 1,\"acceptValue\": 1},\"aksWeight\": {\"dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN\": 1}}"
    }
}

#创建签名地址目录
mkdir -p data/acl

#写入签名地址到文件
echo "XC1111111111111111@xuper/dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN" > data/acl/addrs

./xchain-cli multisig gen --desc addnode.json --from XC1111111111111111@xuper --fee 1
./xchain-cli multisig sign --tx tx.out --output sign.out --keys data/keys
./xchain-cli multisig send --tx tx.out sign.out sign.out

#--tx参数：交易文件，发起者的签名，在签名地址中所有地址的签名；如果有多个签名，用逗号隔开。

#查看是否设置成功
./xchain-cli acl query --contract group_chain --method addNode

#此时用其他账户调用addNode会报：Failed to post tx:RWACL_INVALID_ERROR
```

## 更改手续费

```bash
#转账手续费
vi fee.json
{
    "Module": "kernel",
    "Method": "UpdateTransferFeeAmount",
    "Args": {
        "old_transfer_fee_amount": 1000, #当前手续费
        "new_transfer_fee_amount": 10000 #新的手续费
    }
}
#执行更改操作
./xchain-cli transfer --desc fee.json


#创建合约账户手续费
vi acc.json
{
    "Module": "kernel",
    "Method": "UpdateNewAccountResourceAmount",
    "Args": {
        "old_new_account_resource_amount": 1000,
        "new_new_account_resource_amount": 10000
    }
}
#执行更改操作
./xchain-cli transfer --desc acc.json
```

## 配置说明

### xuper.json

```bash
{
    "version" : "1",
    
    #预存账户（可以有多个）
    "predistribution":[
        {
            "address" : "dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN", # 预存地址
            "quota" : "100000000000000000000"                # 预存金额
        }
    ],
    
    # 最大块大小MB
    "maxblocksize" : "128",
    
    # 挖矿奖励金额
    "award" : "1000000",
    
    # 精度（没有用的字段）
    "decimals" : "8",
    
    # 挖块奖励衰减
    "award_decay": {
        "height_gap": 31536000, # 衰减区块高度
        "ratio": 1              # 衰减比例，1为不衰减，0.5为衰减一般
    },
    
    # 共识机制
    "genesis_consensus": {
    	# 共识名称
        "name": "tdpos",
        # 共识配置
        "config": {
            # tdpos共识初始时间，声明tdpos共识的起始时间戳，建议设置为一个刚过去不旧的时间戳
            "timestamp": "1548123921000000000",
            # 每轮的旷工数，如果某一轮的投票不足以选出足够的矿工数则复用前一轮的矿工
            "proposer_num":"2",
            # 出块间隔
            "period":"3000",
            # 切换矿工时的间隔，需要为period的整数倍
            "alternate_interval":"6000",
            # 切换下一轮时的间隔，需要为period的整数配
            "term_interval":"9000",
            # 每个矿工的出块个数
            "block_num":"200",
            # 投票时候选人的每一票单价
            "vote_unit_price":"1",
            # 指定初始矿工，个数需要符合proposer_num的个数，指定的矿工需要在网络中存在，不然系统轮到该节点出块时会不出块
            #注意，这个map只能有“1”这个key，如果要增加旷工，在数据中增加即可
            "init_proposer": {
                "1":["RU7Qv3CrecW5waKc1ZWYnEuTdJNjHc43u","XpQXiBNo1eHRQpD9UbzBisTPXojpyzkxn"]
            }
        }
    }
}
```

### single共识

```bash
#字段说明参考tdpos
{
    "version" : "1"
    , "predistribution":[
        {
            "address" : "dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN"
            , "quota" : "100000000000000000000"
        }
    ]
    , "maxblocksize" : "128"
    , "award" : "428100000000"
    , "decimals" : "8"
    , "award_decay": {
        "height_gap": 31536000,
        "ratio": 1
    },"genesis_consensus":{
        "name": "single",
        "config": {
                        "miner":"dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN",#出块地址
                        "period": "3000"                            #出块间隔
                }
        }
}
```



## 错误记录

```bash
msg="ProcessBlock SendBlock error" err="get block error when sync block"
msg="HandleSendBlock ProcessBlock error" error="get block error when sync block"
#节点的创世配置跟主链不一样，修改vi data/config/xuper.json，改成跟主链配置一样

Failed to post tx:TX_SIGN_ERROR
# 检查data/acl/addrs是否跟ACL地址匹配

Failed to post tx:RWACL_INVALID_ERROR
# 检查方法的ACL是否需要签名
# 调用合约时，data/acl/addrs不能包含合约账户
# 修改方法权限时，data/acl/addrs需要包含合约账户

Failed to post tx:TX_DUPLICATE_ERROR
# 该笔签名交易已经生效了，需要重新生成tx.out

Failed to post tx:CONNECT_REFUSE
# 用返回的logid去logs/xchain.logs看具体错误，如果是创建平行链合约账户，指定--fee 1000即可

#修改合约方法的acl，签名地址必须要包含合约账号，以下2种会报错
    #只添加合约账户的acl管理账户 Failed to post tx:RWACL_INVALID_ERROR
    echo "dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN" > data/acl/addrs

    #添加合约方法的调用账户 Failed to post tx:TX_SIGN_ERROR
    echo "Ybqc14FZnvNxpJCBLLdBYrKF2Y2mHmpcs" > data/acl/addrs
    
    #这种才正确，包含合约地址
    echo "XC1111111111111111@xuper/dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN" > data/acl/addrs
    #如果方法alc包含其他账户权重，地址中也要加上
    echo "XC1111111111111111@xuper/Ybqc14FZnvNxpJCBLLdBYrKF2Y2mHmpcs" >> data/acl/addrs

#调用合约时的错误：
#只有noed1签名 RWACL_INVALID_ERROR
./xchain-cli multisig send --tx tx.out node1.out node1.out
#签名地址没有node2的地址 TX_SIGN_ERROR
./xchain-cli multisig send --tx tx.out node1.out node1.out,node2.out


#创建平行链合约账户是，需要指定--fee 1，否则会连接被拒
Failed to post tx:CONNECT_REFUSE


#本地账户没有合约账户的操作权限
verify contract owner permission failed

#检查下合约名有没有错
contract for account not confirmed
```





# 测试步骤

### 环境配置

```bash
sudo apt-get install docker.io
sudo apt-get install git
sudo apt-get install gcc
sudo apt-get install g++
sudo apt-get install make
sudo apt-get install zip

# 安装golang
#下载安装包
wget https://studygolang.com/dl/golang/go1.13.5.linux-amd64.tar.gz

#解压
sudo tar -C /usr/local -xzf go1.13.5.linux-amd64.tar.gz

#加入环境变量
vi /etc/profile
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$GOPATH:$GOBIN:$GOROOT/bin:$PATH
export GOPROXY="https://goproxy.cn,direct"

#更新
source /etc/profile

#验证安装是否成功
go version
返回：go version go1.13.5 linux/amd64
```

### 获取版本

```bash
git clone https://github.com/StarAllianceFoundation/SAFChain.git
```

### 编译

```bash
cd SAFChain && make

#复制一份输出文件作为节点2来用
cp -r output output2

#进入目录
cd output
```



## 节点一

### 启动节点

```bash
./xchain-cli createChain
./xchain
```

### 创建合约账户

```bash
./xchain-cli account new --account 1111111111111111 --fee 500000
./xchain-cli transfer --to XC1111111111111111@xuper --amount 1000000000 --fee 500000
```

### 安装群组合约

```bash
#编译合约
cd ../core/contractsdk/cpp
./build.sh
cp -r build/ ../../../cc

#部署
./xchain-cli wasm deploy --account XC1111111111111111@xuper --cname group_chain ../cc/group_chain.wasm --fee 500000

#查看是否安装
./xchain-cli account contracts --account XC1111111111111111@xuper

#跳过这里继续往下看，下面的是多方签名的部署合约步骤
#多方签名操作步骤（如果合约账户是多方的）
#-m 为该操作需要多方签名
./xchain-cli wasm deploy --account XC1111111111111111@xuper --cname group_chain ../cc/group_chain.wasm --fee 500000 -m

mkdir -p data/acl
echo "XC1111111111111111@xuper/SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635" > data/acl/addrs
./xchain-cli multisig sign --tx tx.out --output sign.out --keys data/keys
./xchain-cli multisig send --tx tx.out sign.out sign.out
```



## 节点二

### 添加启动时连接的节点

```bash
#获取创世节点的ip
./xchain-cli netURL get

#进入节点2
cd output2
vi conf/xchain.yaml
	#修改端口，避免冲突
	tcpServer:
  		port: :37102
  		metricPort: :37202
    p2p:
      port: 47102
      
	#添加启动时连接的创世节点 
    - "/ip4/127.0.0.1/tcp/47101/p2p/QmVxeNubpg1ZQjQT8W5yZC9fD7ZB1ViArwvyGUB53sqf8e"
```

### 创建账号和netURL

```bash
./xchain-cli account newkeys -f
./xchain-cli netURL gen
```

### 启动节点

```bash
./xchain-cli createChain
./xchain

#查看节点的状态，此时同步了主链的区块
./xchain-cli status -H :37101 | grep Height
./xchain-cli status -H :37102 | grep Height
```



## 创建平行链

```bash
#节点一执行，它的账户才有钱
cd ../output

vi wtf.json

{
    "Module": "kernel",
    "Method": "CreateBlockChain",
    "Args": {
        "name": "wtf",
        "data": "{\"version\":\"1\",\"consensus\":{\"miner\":\"SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635\",\"type\":\"single\"},\"predistribution\":[{\"address\":\"\",\"quota\":\"\"}],\"maxblocksize\":\"128\",\"period\":\"10000\",\"award\":\"10000\",\"award_decay\":{\"height_gap\":210000,\"ratio\":0.5},\"gas_price\":{\"cpu_rate\":0,\"mem_rate\":0,\"disk_rate\":0,\"xfee_rate\":1},\"new_account_resource_amount\":1000,\"transfer_fee_amount\":1000}"
    }
}

# 注意data的配置出了共识的字段名，其他都可以跟主链配置一样
# 可以通过该网站转义json http://json.cn/json/jsonzip.html

#创建平行链
#查看当前全网总额
./xchain-cli status | grep utxo
#查看当前区块高度
./xchain-cli status | grep Height

#--amount:需要大于全网总额50%，--frozen：需要高于当前高度
./xchain-cli transfer --desc wtf.json --amount 1345000000000 --fee 500000 --frozen 350

#在原版的conf/xuper.yaml中规定了创建需要转100块钱过去，不过该值是动态的，也就是关掉链修改配置改成0，就不用钱了。这个可能会引发某些问题，在金银链中限制了只有转账金额大于全网总额50%才能够创建平行链。

#查看平行链的区块高度
./xchain-cli status  | grep wtf -A 5 | grep Height
./xchain-cli status -H :37102  | grep wtf -A 5 | grep Height
```



## 添加群组配置

```bash
#获取第二个节点的地址
cat ../output2/data/keys/address
#address：SAFT4yvUpe2fmWBKF6FKKheur9hpmukCMr8h

#转点钱过去
./xchain-cli transfer --to 节点2地址 --amount 1000000000 --fee 500000
#./xchain-cli transfer --to SAFT4yvUpe2fmWBKF6FKKheur9hpmukCMr8h --amount 1000000000 --fee 500000

#获取第二个节点的url
./xchain-cli netURL get -H :37102
#neturl："/ip4/127.0.0.1/tcp/47102/p2p/QmRJAoWnA9WZBBhjDZH12onDBtnAsj2nrSReQfGzSNZVZo"

#从节点一执行，查看第二个节点的ip
./xchain-cli status | grep peers -A 2
#peer："peers": [ "127.0.0.1:37102" ]

#添加节点，替换地址和ip
#节点一
./xchain-cli wasm invoke group_chain --method addNode -a '{"bcname":"wtf", "ip":"/ip4/127.0.0.1/tcp/47101/p2p/QmVxeNubpg1ZQjQT8W5yZC9fD7ZB1ViArwvyGUB53sqf8e", "address":"SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635"}' --fee 500000

#节点二
./xchain-cli wasm invoke group_chain --method addNode -a '{"bcname":"wtf","ip":"节点2ip", "address":"节点2地址"}' --fee 500000
#./xchain-cli wasm invoke group_chain --method addNode -a '{"bcname":"wtf","ip":"/ip4/127.0.0.1/tcp/47102/p2p/QmRJAoWnA9WZBBhjDZH12onDBtnAsj2nrSReQfGzSNZVZo", "address":"SAFT4yvUpe2fmWBKF6FKKheur9hpmukCMr8h"}' --fee 500000

#查看群组中的节点
./xchain-cli wasm query group_chain --method listNode -a '{"bcname":"wtf"}'

#添加平行链到群组管理
./xchain-cli wasm invoke group_chain --method addChain -a '{"bcname":"wtf"}' --fee 500000

#群组管理中的链列表
./xchain-cli wasm query group_chain --method listChain

#查看平行链的区块高度
./xchain-cli status  | grep wtf -A 5 | grep Height
./xchain-cli status -H :37102  | grep wtf -A 5 | grep Height

#从wtf删除节点2
./xchain-cli wasm invoke group_chain --method delNode -a '{"bcname":"wtf", "ip":"节点2ip"}' --fee 500000
#./xchain-cli wasm invoke group_chain --method delNode -a '{"bcname":"wtf", "ip":"/ip4/127.0.0.1/tcp/47102/p2p/QmRJAoWnA9WZBBhjDZH12onDBtnAsj2nrSReQfGzSNZVZo"}' --fee 500000

#查看节点列表
./xchain-cli wasm query group_chain --method listNode -a '{"bcname":"wtf"}'
#此时已经删除掉了节点2

#查看平行链的区块高度
./xchain-cli status  | grep wtf -A 5 | grep Height
./xchain-cli status -H :37102  | grep wtf -A 5 | grep Height
#此时可以看到节点2的区块已经不同步了

#重新添加节点2
./xchain-cli wasm invoke group_chain --method addNode -a '{"bcname":"wtf","ip":"节点2ip", "address":"节点2地址"}' --fee 500000
#./xchain-cli wasm invoke group_chain --method Node -a '{"bcname":"wtf","ip":"/ip4/127.0.0.1/tcp/47102/p2p/QmRJAoWnA9WZBBhjDZH12onDBtnAsj2nrSReQfGzSNZVZo", "address":"SAFT4yvUpe2fmWBKF6FKKheur9hpmukCMr8h"}' --fee 500000

#此时再查看节点2，区块又在同步了，可能有延迟
./xchain-cli status  | grep wtf -A 5 | grep Height
./xchain-cli status -H :37102  | grep wtf -A 5 | grep Height

#剔除管理中的某一条链
./xchain-cli wasm invoke group_chain --method delChain -a '{"bcname":"wtf"}' --fee 500000
./xchain-cli wasm query group_chain --method listChain
```



## 群组合约方法ACL设置

```bash
#官方并没有提供批量设置权限的接口，目前只能一个一个调用，后续有需求的话就自己加吧

vi addnode.json
{
    "module_name": "xkernel",
    "method_name": "SetMethodAcl",
    "args" : {
        "contract_name": "group_chain",
        "method_name": "addNode",
        "acl": "{\"pm\": {\"rule\": 1,\"acceptValue\": 1},\"aksWeight\": {\"SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635\": 1,\"其他地址\": 1}}"
    }
}
#实例：
{
    "module_name": "xkernel",
    "method_name": "SetMethodAcl",
    "args" : {
        "contract_name": "group_chain",
        "method_name": "addNode",
        "acl": "{\"pm\": {\"rule\": 1,\"acceptValue\": 1},\"aksWeight\": {\"SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635\": 1,\"SAFT4yvUpe2fmWBKF6FKKheur9hpmukCMr8h\": 1}}"
    }
}

#创建签名地址，只需要合约账户所有者的签名即可
mkdir -p data/acl
echo "XC1111111111111111@xuper/SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635" > data/acl/addrs
./xchain-cli multisig gen --desc addnode.json --from XC1111111111111111@xuper --fee 500000
./xchain-cli multisig sign --tx tx.out --output sign.out --keys data/keys
./xchain-cli multisig send --tx tx.out sign.out sign.out

#查看一下方法的acl
./xchain-cli acl query --contract group_chain --method addNode
#可以看到合约方法acl已经设置成功了
#如果显示unconfirmed，等待1分钟后再查即可

#重复以上步骤，设置一下删除节点的合约方法acl，并且只保留一个地址
vi delnode.json
{
    "module_name": "xkernel",
    "method_name": "SetMethodAcl",
    "args" : {
        "contract_name": "group_chain",
        "method_name": "delNode",
        "acl": "{\"pm\": {\"rule\": 1,\"acceptValue\": 1},\"aksWeight\": {\"SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635\": 1}}"
    }
}
./xchain-cli multisig gen --desc delnode.json --from XC1111111111111111@xuper --fee 500000
./xchain-cli multisig sign --tx tx.out --output sign.out
./xchain-cli multisig send --tx tx.out sign.out sign.out
./xchain-cli acl query --contract group_chain --method delNode
#可以看到合约方法acl已经设置成功了

#测试是否可用
#目前2个节点都调用addNode
#node1
./xchain-cli wasm invoke group_chain --method addNode -a '{"bcname":"wtf", "ip":"node1"
, "address":"test"}' --fee 500000
#node2
./xchain-cli wasm invoke group_chain --method addNode -a '{"bcname":"wtf", "ip":"node2"
, "address":"test"}' --fee 500000 --keys ../output2/data/keys 

#查看下node列表
./xchain-cli wasm query group_chain --method listNode -a '{"bcname":"wtf"}'
#可以看到ip已经添加

#测试删除
./xchain-cli wasm invoke group_chain --method delNode -a '{"bcname":"wtf", "ip":"node1"}' --fee 500000
#返回成功
./xchain-cli wasm invoke group_chain --method delNode -a '{"bcname":"wtf", "ip":"node2"}' --fee 500000 --keys ../output2/data/keys
#返回失败：Failed to post tx:RWACL_INVALID_ERROR acl错误，可以拿返回的logid去xuper.log查，可以看到错误是:msg="post tx verify tx error" valid_err="ACL not enough"

#测试多签操作
#修改删除节点的方法acl，这个需要2个节点签名
vi delnode.json
{
    "module_name": "xkernel",
    "method_name": "SetMethodAcl",
    "args" : {
        "contract_name": "group_chain",
        "method_name": "delNode",
        "acl": "{\"pm\": {\"rule\": 1,\"acceptValue\": 1},\"aksWeight\": {\"SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635\": 0.5,\"其他地址\": 0.5}}"
    }
}

#实例：
{
    "module_name": "xkernel",
    "method_name": "SetMethodAcl",
    "args" : {
        "contract_name": "group_chain",
        "method_name": "delNode",
        "acl": "{\"pm\": {\"rule\": 1,\"acceptValue\": 1},\"aksWeight\": {\"SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635\": 0.5,\"SAFT4yvUpe2fmWBKF6FKKheur9hpmukCMr8h\": 0.5}}"
    }
}

#修改acl
./xchain-cli multisig gen --desc delnode.json --from XC1111111111111111@xuper --fee 500000
./xchain-cli multisig sign --tx tx.out --output sign.out --keys data/keys
./xchain-cli multisig send --tx tx.out sign.out sign.out
./xchain-cli acl query --contract group_chain --method delNode

#先添加node2
./xchain-cli wasm invoke group_chain --method addNode -a '{"bcname":"wtf", "ip":"node2"
, "address":"test"}' --fee 500000
./xchain-cli wasm query group_chain --method listNode -a '{"bcname":"wtf"}'

#测试删除
./xchain-cli wasm invoke group_chain --method delNode -a '{"bcname":"wtf", "ip":"node2"
, "address":"test"}' --fee 500000
#返回失败：Failed to post tx:RWACL_INVALID_ERROR acl错误，可以拿返回的logid去xuper.log查，可以看到错误是:msg="post tx verify tx error" valid_err="ACL not enough"

#采用多方调用
#添加签名地址 不能包含合约账户，会报RWACL_INVALID_ERROR错误
echo "SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635" > data/acl/addrs
echo "其他地址" >> data/acl/addrs
#echo "SAFT4yvUpe2fmWBKF6FKKheur9hpmukCMr8h" >> data/acl/addrs

#查看是否写入地址列表
cat data/acl/addrs 
#SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635
#SAFT4yvUpe2fmWBKF6FKKheur9hpmukCMr8h

#生成交易文件
./xchain-cli wasm invoke group_chain --method delNode -a '{"bcname":"wtf", "ip":"node2"
, "address":"test"}' --fee 500000 -m

#node1签名
./xchain-cli multisig sign --tx tx.out --output node1.out --keys data/keys
#node2签名
./xchain-cli multisig sign --tx tx.out --output node2.out --keys ../output2/data/keys

#发送交易
./xchain-cli multisig send --tx tx.out node1.out,node2.out node1.out,node2.out
./xchain-cli wasm query group_chain --method listNode -a '{"bcname":"wtf"}'
#至此 test2 已删除
```



## 平行链安装合约

```bash
#使用 --name wtf 指定每次连接的链

#先创建合约账户，平行链使用的是单节点共识，所以只有SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635是有钱的
./xchain-cli account new --account 1111111111111111 --name wtf --fee 1000

#转点钱给合约账户
./xchain-cli transfer --to XC1111111111111111@wtf --amount 1000000 --name wtf --fee 1000

#有了钱就可以创建合约了，这里使用counter合约
./xchain-cli wasm deploy --account XC1111111111111111@wtf --cname counter ../cc/counter.wasm -a '{"creator":"xchain"}' --name wtf --fee 1000

#调用
#调用次数自增1
./xchain-cli wasm invoke -a '{"key":"counter"}' --method increase counter --name wtf --fee 1000

#获取调用次数
./xchain-cli wasm query -a '{"key":"counter"}' --method get counter --name wtf
```



## 更换共识

```bash
#查看区块高度
./xchain-cli status  | grep xuper -A 5 | grep trunkHeight
```

```bash
vi proposal.json
{
    "module": "proposal",
    "method": "Propose",
    "args" : {
        "min_vote_percent": 51,
        "stop_vote_height": 950
    },
    "trigger": {
        "height": 951,
        "module": "consensus",
        "method": "update_consensus",
        "args" : {
            "name": "tdpos",
            "config": {
                "version":"2",
                "proposer_num":"2",
                "period":"3000",
                "alternate_interval":"3000",
                "term_interval":"3000",
                "block_num":"5",
                "vote_unit_price":"100000000",
                "vote_award": "5000000000",
                "init_proposer": {
                    "1":["SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635", "node2"]
                }
            }
        }
    }
}

#记住此次的交易id
./xchain-cli transfer --desc proposal.json --fee 500000
#返回的txid:c1f2f0a23c568e66f62ec51203e9956827bbb52f0589fd70b3276c790c0daa5c

#min_vote_percent设置了投票比例需要大于全网总额的51%
#查看全网总额
./xchain-cli status  | grep xuper -A 9 | grep utxoTotal
#xuper链的"utxoTotal": "4445000000000"

#查看每个账户的钱够不够达到这个值
./xchain-cli account balance
#node1余额：2967997499999
./xchain-cli account balance --keys ../output2/data/keys/
#node2余额：1516004500000

#看看是否有金额是被冻结了的
./xchain-cli account balance -Z
./xchain-cli account balance --keys ../output2/data/keys/ -Z

#对提案进行投票 --frozen：冻结高度需要大于生效高度
./xchain-cli vote c1f2f0a23c568e66f62ec51203e9956827bbb52f0589fd70b3276c790c0daa5c --frozen 952 --amount 2867997499999 --fee 500000
./xchain-cli vote c1f2f0a23c568e66f62ec51203e9956827bbb52f0589fd70b3276c790c0daa5c --frozen 952 --amount 1416004500000 --fee 500000 --keys ../output2/data/keys/

#等待区块高度达到切换高度时，查看共识，可能会有延迟
./xchain-cli tdpos status

#./xchain-cli tdpos status
{
  "term": 8,
  "block_num": 1,
  "proposer": "node2",
  "proposer_num": 2,
  "checkResult": [
    "SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635",
    "node2"
  ]
}

#./xchain-cli tdpos status -H :37102
rpc error: code = Unknown desc = leveldb: not found
#bug：无法同步数据
```



## 选举

### 操作说明

```bash
#tdpo状态
./xchain-cli tdpos status
#候选人列表
./xchain-cli tdpos query-candidates
./xchain-cli tdpos query-checkResult

#查看所有候选人的提名的记录
./xchain-cli tdpos query-nominate-records
#查看某用户的提名候选人记录
./xchain-cli tdpos query-nominate-records -a SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635
#查看候选人的被提名记录
./xchain-cli tdpos query-nominee-record -a SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635

#查看某用户的投票记录
./xchain-cli tdpos query-vote-records -a SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635

#查看某用户的得票记录
./xchain-cli tdpos query-voted-records -a SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635

./xchain-cli tdpos query-voted-records -a SAFncP8LXCMvbi1vAyY7ce1kEv21xGMchsC8
```

### 提名候选人（自己）

```bash
#提名候选人
vi node1.json
{
    "module": "tdpos",
    "method": "nominate_candidate",
    "args": {
        "candidate": "SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635",
        "neturl": "/ip4/127.0.0.1/tcp/47101/p2p/QmVxeNubpg1ZQjQT8W5yZC9fD7ZB1ViArwvyGUB53sqf8e"
    }
}

#提名者和被提名候选人的两个签名(相同一个就行)
mkdir data/acl
echo "SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635" > data/acl/addrs

#当前链的总资产，需要抵押总资产的10万分之一(取消提名后才退)
./xchain-cli status  | grep xuper -A 9 | grep utxoTotal

#生成交易文件
./xchain-cli multisig gen --output nominate.tx --desc=node1.json --amount=515000000 --frozen -1 --fee 500000

#对交易文件签名
./xchain-cli multisig sign --tx nominate.tx --output nominate.sign

#发送交易
#nominate.tx：交易文件
#nominate.sign：交易发起人的签名
#nominate.sign：交易的acl的签名（有多个就用逗号隔开）
./xchain-cli multisig send --tx nominate.tx nominate.sign nominate.sign

#查下候选人
./xchain-cli tdpos query-candidates
#["SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635"]
```



### 提名候选人（别人）

```bash
vi node2.json
{
    "module": "tdpos",
    "method": "nominate_candidate",
    "args": {
        "candidate": "要提名成候选人的地址",
        "neturl": "要提名成候选人的netURL"
    }
}
#实例：
{
    "module": "tdpos",
    "method": "nominate_candidate",
    "args": {
        "candidate": "SAFT4yvUpe2fmWBKF6FKKheur9hpmukCMr8h",
        "neturl": "/ip4/127.0.0.1/tcp/47102/p2p/QmRJAoWnA9WZBBhjDZH12onDBtnAsj2nrSReQfGzSNZVZo"
    }
}

#提名者和被提名候选人的两个签名
echo "SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635" > data/acl/addrs
echo "要提名成候选人的地址" >> data/acl/addrs
#echo "SAFT4yvUpe2fmWBKF6FKKheur9hpmukCMr8h" >> data/acl/addrs
./xchain-cli multisig gen --output nominate.tx --desc=node2.json --amount=515000000 --frozen -1 --fee 500000

#node1签名
./xchain-cli multisig sign --tx nominate.tx --output nominate.sign

#node2签名
./xchain-cli multisig sign --tx nominate.tx --output candidate.sign --keys ../output2/data/keys

# 发送交易
./xchain-cli multisig send --tx nominate.tx nominate.sign nominate.sign,candidate.sign
#Txid: 7a6c5b67b71b982ad74ef08e19038b90e0be6c07256bbeb70a251d58c8a4c818

#查下候选人
./xchain-cli tdpos query-candidates
#[
#  "SAFT4yvUpe2fmWBKF6FKKheur9hpmukCMr8h",
#  "SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635"
#]
```

### node3（忽略这个，我测试新创建节点同步用的）

```bash
vi node3.json
{
    "module": "tdpos",
    "method": "nominate_candidate",
    "args": {
        "candidate": "SAFhpY8S3rfgoTpugaqRHJvVan7xueLa6qvq",
        "neturl": "/ip4/127.0.0.1/tcp/47103/p2p/QmXDtZZs9mYNT5Fk7keFCZ7ND89jCVe9N2KTpyfsLYvVrH"
    }
}
echo "SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635" > data/acl/addrs
echo "SAFhpY8S3rfgoTpugaqRHJvVan7xueLa6qvq" >> data/acl/addrs
./xchain-cli multisig gen --output nominate.tx --desc=node3.json --amount=2671500000 --frozen -1 --fee 500000
./xchain-cli multisig sign --tx nominate.tx --output nominate.sign
./xchain-cli multisig sign --tx nominate.tx --output candidate.sign --keys ../output3/data/keys
./xchain-cli multisig send --tx nominate.tx nominate.sign nominate.sign,candidate.sign
./xchain-cli tdpos query-candidates

vi vote.json3
{
    "module": "tdpos",
    "method": "vote",
    "args": {
        "candidates":["SAFhpY8S3rfgoTpugaqRHJvVan7xueLa6qvq"]
    }
}
./xchain-cli transfer --desc=vote.json --amount=100000000 --frozen -1 --fee 500000
./xchain-cli tdpos status
```

### 投票

```bash
vi vote.json
{
    "module": "tdpos",
    "method": "vote",
    "args": {
        "candidates":["候选人地址"]
    }
}
#实例：投一个
{
    "module": "tdpos",
    "method": "vote",
    "args": {
        "candidates":["SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635"]
    }
}
#实例：投多个，但不能超过共识规定的候选人数量
{
    "module": "tdpos",
    "method": "vote",
    "args": {
        "candidates":["SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635","SAFT4yvUpe2fmWBKF6FKKheur9hpmukCMr8h"]
    }
}

#投票 --amount=100000000：tdpos共识配置中规定了一票的单机就是1SAF --frozen -1：需要一直冻结该金额
./xchain-cli transfer --desc=vote.json --amount=100000000 --frozen -1 --fee 500000

#查看投票状态
./xchain-cli tdpos status

#另一个账号的挖矿金额也在上升了
./xchain-cli account balance 候选人地址
```



### 撤销提名 && 撤销投票

```bash
#查看冻结中的金额
./xchain-cli account balance -Z
#金额：5501500003

#查看投票记录
./xchain-cli tdpos query-vote-records -a SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635
#其中一笔txid：f782728513d4057db3b7d3a7de35610f59113a20fd1a5baabd5829279df9ae1b

vi thaw.json
{
    "module":"proposal",
    "method": "Thaw",
    "args" : {
        "txid":"此处为提名或者投票时的txid，且address与提名或者投票时需要相同"
    }
}
#实例：
{
    "module":"proposal",
    "method": "Thaw",
    "args" : {
        "txid":"f782728513d4057db3b7d3a7de35610f59113a20fd1a5baabd5829279df9ae1b"
    }
}

./xchain-cli transfer --desc=thaw.json --fee 500000

#再次查看，记录中已经删除了
./xchain-cli tdpos query-vote-records -a SAFn9w5b2nMxsj2e8tsTT9n9G89g4e3xm635

#此时被冻结的金额解冻了
./xchain-cli account balance -Z
#金额：5401500003
```

## 结论记录

1. 重复提名候选人，会冻结金额，但是提名不生效。（注意记录交易id，否则无法解冻金额）
2. 撤销提名候选人，立即解冻金额，但在下一轮次才从候选人列表移除。
3. 撤销投票，立即解冻金额，但在下一轮次才从投票记录移除。
4. 投票时，候选人相同，会剔除掉相同数据
5. 投票时，投票人大于提议人数，该投票不会生效。（注意记录交易id，否则无法解冻金额）
6. 投票时，投票人不是候选人，该投票不会生效
7. 投票金额小于最低票价的话，该投票会生效，但是票数为0（从源码来看，是不能设置为0的，会引发恐慌）
8. 不出块的候选人，不会有奖励
9. 当遇到票数相等时，会逆序后选出验证人列表
10. 孤儿块会被抛弃



 ## C++合约编译

```bash
#依赖与docker来编译，xdev提供了测试的功能
#具体测试合约在test目录中

#安装好docker（这条命令是ubuntu的）
sudo apt install docker.io

#第一次编译会去下载baidu的镜像
#也可以手动下载：docker pull hub.baidubce.com/xchain/emcc

#编译好xdev，最好创建xdev的软连接到/usr/bin/这样就不用每次都进入output目录了（可选）
ln -s xdev的绝对路径 /usr/bin/

#测试编译
cp ../core/contractsdk/cpp/example/erc20.cc ./
./xdev build -o erc20.wasm erc20.cc
#报错：
#Error: XDEV_ROOT and XCHAIN_ROOT must be set one.
#XCHAIN_ROOT is the path of $xuperchain_project_root/core.
#XDEV_ROOT is the path of $xuperchain_project_root/core/contractsdk/cpp

#需要设置路径
export XCHAIN_ROOT=/root/xchain/tools/xuperchain/core #这个可以不设置
export XDEV_ROOT=~/go/src/github.com/jason-cn-dev/xuperchain/core/contractsdk/cpp #设置了xdev的输出路径

#重新编译
./xdev build -o erc20.wasm erc20.cc
#第一次运行会编译一些其他的依赖文件
CC erc20.cc
CC account.cc
CC basic_iterator.cc
CC block.cc
CC context_impl.cc
CC contract.cc
CC contract.pb.cc
CC crypto.cc
CC syscall.cc
CC transaction.cc
LD wasm
#在当前目录下生成了erc20.wasm文件

#第二次就不会了
cp ../core/contractsdk/cpp/example/counter.cc ./
./xdev build -o counter.wasm counter.cc
CC counter.cc
LD wasm
#在当前目录下生成了counter.wasm文件
```



# 开启http协议接口

```bash
#注意：v3.7的版本是没有背书和跨域功能的，SAFchain用的是v3.8的代码

## 编译
go build -o xchain-httpgw gateway/http_gateway.go

## 部署
#-http_endpoint：http服务侦听的端口，默认是：8123
#-gateway_endpoint：xchain节点的tcp端口 ，默认是：37101
#-enable_endorser：是否启动背书，就是用来给gosdk外的其他sdk跟链操作用的，默认是：true
#-allow_crosr：是否启动跨域，默认是：true
nohup ./xchain-httpgw -http_endpoint :8123 -gateway_endpoint :37101 &

# RPC接口介绍，照着填参数就行了，注意：根据txid查询交易不能直接使用，需要将txid转换成16进制的字符串才行
https://xuperchain.readthedocs.io/zh/latest/development_manuals/XuperRPC.html

#在控制台试试能不能调用接口：
#查询地址的余额
curl http://localhost:8123/v1/get_balance -d '{"bcs":[{"bcname":"xuper"}],"address":"dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN"}'
#查询链的状态
curl http://localhost:8123/v1/get_bcstatus -d '{"bcname":"xuper"}'

#使用postman试试能不能调用接口：
http://192.168.3.150:8123/v1/get_balance
#method：post 
#contentType：json/appliction
{"bcs": [{"bcname": "xuper"}],"address": "dpzuVdosQrF2kmzumhVeFQZa1aYcdgFpN"}
```



# Mongodb

## Docker-Mongodb安装

拉取镜像

```bash
docker pull mongo:latest
```

启动容器

```bash
docker run -itd --name mongo -p 27017:27017 mongo --auth
```

进入容器

```bash
docker exec -it mongo mongo admin
```

创建账户

```bash
db.createUser({ user:'admin',pwd:'this is mongodb admin password',roles:[ { role:"userAdminAnyDatabase", db:"admin" }, "readWriteAnyDatabase" ]});
```

登录账户

```bash
db.auth("admin", "this is mongodb admin password")
```

选择数据库

```bash
use jy_chain
```

## 区块订阅同步

```bash
#编译
git clone https://github.com/chaojigongshi-superconsensus/xuperdata.git
cd xuperdata
go build -mod vendor
#生成xuperdata执行文件

#启动，-restore：清空数据库，不想清空就不要加上
./xuperdata -restore -s 'mongodb://admin:this is mongodb admin password@0.0.0.0:27017'

#选项说明：
-f：订阅的json文件，默认是：json/block.json （订阅所有的出块）
-c：订阅或取消订阅，默认是：subscribe
-id：订阅事件的uuid，默认是：000
-h：节点的ip，默认是：localhost:37101
-s：mongodb数据源，默认是：mongodb://localhost:27017
-b：mongodb数据库，默认是：jy_chain
-port：供钱包发送交易id过来的端口，默认是：8081
-gosize：同步区块时的线程数，默认是：10
-restore：是否清空数据库重新同步，默认是：false
-show：是否显示接收到区块时打印该区块的高度，默认是：false

#钱包发交易来的接口：
get请求： ip:8081/getTxid?txid="123456"
#例如：http://161.117.39.102:8081/getTxid?txid=9ba381ebcc9e9066bcd4c1bbbda887c398248d4986172d60ddbaeeb117beaed3
```

