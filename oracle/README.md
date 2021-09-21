预言机的技术方案和实现等

# Goland注释
Goland插件安装Goanno，在tool-goanno setting中可修改注释模板

用法：
Mac用户，在函数上方使用组合键control+commend+/
windows用户，在函数上方使用组合键control+window+/

# 计算预言机配置

计算预言机涉及到训练，需要配置pytorch的环境。
在mac本机，需要采用如下操作:
``` 
 conda activate TF2.1
```
在服务器，需要采用如下操作：
``` 
ssh root@10.3.242.104

Gao506bupt

source ~/.bashrc

conda activate test
```

在首次运行时，由于需要下载数据集，花费的时间较长。