# C++合约编译

xdev是xuperchain的合约测试与编译工具，在对xuperchain进行编译`make`的同时，xdev也会编译生成在output中。

cpp合约编译需要依赖docker

在output目录下，将合约拷贝到相应的目录中

```
#测试编译
cp ../core/contractsdk/cpp/example/erc20.cc ./
./xdev build -o erc20.wasm erc20.cc
#报错：
#Error: XDEV_ROOT and XCHAIN_ROOT must be set one.
#XCHAIN_ROOT is the path of $xuperchain_project_root/core.
#XDEV_ROOT is the path of $xuperchain_project_root/core/contractsdk/cpp

#需要设置路径
export XDEV_ROOT=~/go/src/github.com/xuperchain/core/contractsdk/cpp
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

或者是创建项目工程

```
./xdev init adv
#src文件下的cpp文件为源代码
#将xdev拷贝到工程文件夹中，才能执行编译操作
cp ./xdev ./adv
./xdev build -o adv.wasm
```

