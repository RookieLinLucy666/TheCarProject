## fisco智能合约命令行测试测试

注：下文中所有 $address 均为实际的合约部署地址

### 部署

启用fisco控制台，并将合约置于/$your_path/fisco/console/contracts/solidity中，然后部署合约

```
deploy ComputingShare
```
若成功则得到类似如下结果，其中第二行为 $address

```
transaction hash: 0xedc10d37171f4e46a7b6f6f2b92cb7543452cabb13f3f359c8c28c5b2b69dad9
contract address: 0x52f75dcffba5ed4a0f9bfab18fbc00d2d6dae48f
currentAccount: 0xe99d90f744aabb523fc13b3b15f147f75310a6c7
```

### 数据上传测试

metadata数据结构如下，若为计算共享元数据则dataType=1，数据共享元数据dataType=2

```
struct metaData{
    string source;
    string dataAbstract;
    uint dataType;
    bool isValid;
}
```
上传计算共享元数据

```
call ComputingShare $address createmetadata "abstract" "source" 1
```
得到结果，其中Return values中的内容为该数据id

```
transaction hash: 0x82ec97ee367acb4bfeba466b2bde595a8792bfe8c5ae9c4e20c54932252b4a88
---------------------------------------------------------------------------------------------
transaction status: 0x0
description: transaction executed successfully
---------------------------------------------------------------------------------------------
Receipt message: Success
Return message: Success
Return value size:1
Return types: (UINT)
Return values:(0)
---------------------------------------------------------------------------------------------
Event logs
Event: {}
```
查询metadata，并得到结果

```
call ComputingShare &address getmetadata 0
```

```
transaction hash: 0x164ac2f50dab2cbaef026aad57f7e313e0aef27ae7c44746d5aa7f6534eeb215
---------------------------------------------------------------------------------------------
transaction status: 0x0
description: transaction executed successfully
---------------------------------------------------------------------------------------------
Receipt message: Success
Return message: Success
Return value size:4
Return types: (STRING, UINT, STRING, BOOL)
Return values:(abstract, 1, source, true)
---------------------------------------------------------------------------------------------
Event logs
Event: {}
```

### 计算共享测试

触发计算共享事件

```
call ComputingShare $address computingshare 0 "model" "dataset" "round" "epoch"
```

```
transaction hash: 0x1536354f4fde16b40b91badb9548ae7841bbf3eca59f42cd9d938a74fdca24e6
---------------------------------------------------------------------------------------------
transaction status: 0x0
description: transaction executed successfully
---------------------------------------------------------------------------------------------
Receipt message: Success
Return message: Success
Return value size:1
Return types: (STRING)
Return values:(emit computingshare event successfully)
---------------------------------------------------------------------------------------------
Event logs
Event: {"computingShareEvent":[[0,"abstract","model","dataset","round","epoch"]]}

```

链下预言机处理计算共享事件，并调用回调函数上传结果

```
 call ComputingShare $address computingsharecallback 0 "paramAddr" "paramAbstract" "88.7%"
```

```
transaction hash: 0x9dbe1d33d5748ea46c847465bc4be82e503ef4d8f2dc5bf0c83b613049e18364
---------------------------------------------------------------------------------------------
transaction status: 0x0
description: transaction executed successfully
---------------------------------------------------------------------------------------------
Receipt message: Success
Return message: Success
Return value size:1
Return types: (STRING)
Return values:(88.7%)
---------------------------------------------------------------------------------------------
Event logs
Event: {}
```
查询结果

```
call ComputingShare $address getresult 0
```

```
transaction hash: 0x0fc9227d10d2500a5651419c97d5c83b677601002dc3a911a3d72084f32e1aa1
---------------------------------------------------------------------------------------------
transaction status: 0x0
description: transaction executed successfully
---------------------------------------------------------------------------------------------
Receipt message: Success
Return message: Success
Return value size:4
Return types: (STRING, STRING, STRING, STRING)
Return values:(paramAddr, paramAbstract, 88.7%, )
---------------------------------------------------------------------------------------------
Event logs
Event: {}
```

### 数据共享测试

上传数据过程略，得到返回的元数据id=1，触发数据共享事件

```
call ComputingShare $address datashare 1
```

```
transaction hash: 0x4875d0769397831c0ca24dca4b87a41eb34a0ca489c3fc7558af0f516247e471
---------------------------------------------------------------------------------------------
transaction status: 0x0
description: transaction executed successfully
---------------------------------------------------------------------------------------------
Receipt message: Success
Return message: Success
Return value size:1
Return types: (STRING)
Return values:(emit datashare event successfully)
---------------------------------------------------------------------------------------------
Event logs
Event: {"dataShareEvent":[[1,"source2","dataAbstract2"]]}

```
预言机处理后进行回调

```
call ComputingShare $address datasharecallback 1 "this is datashare result"
```

```
transaction hash: 0xa7b7cbd64bf576742d3f24a103cbfde3801ca81acef264e77ad8fbc141b5c136
---------------------------------------------------------------------------------------------
transaction status: 0x0
description: transaction executed successfully
---------------------------------------------------------------------------------------------
Receipt message: Success
Return message: Success
Return value size:1
Return types: (STRING)
Return values:(this is datashare result)
---------------------------------------------------------------------------------------------
Event logs
Event: {}
```
查询结果


```
call ComputingShare $address getresult 1
```

```
transaction hash: 0x16aa17f86aa05a479651a22edcadc5768e40df35382fcb4b7f92e1765e5feab7
---------------------------------------------------------------------------------------------
transaction status: 0x0
description: transaction executed successfully
---------------------------------------------------------------------------------------------
Receipt message: Success
Return message: Success
Return value size:4
Return types: (STRING, STRING, STRING, STRING)
Return values:(, , , this is datashare result)
---------------------------------------------------------------------------------------------
Event logs
Event: {}

```











