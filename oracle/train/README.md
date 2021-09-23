
```
ssh -p 666 root@10.112.187.212
88888888
scp -r -P 666 root@10.112.187.212:/home/lhq/py_repos/ ~/Desktop
ssh root@10.3.242.104

Gao506bupt

source ~/.bashrc

conda activate test

conda install pytorch cuda90 -c pytorch

pip3 install torchvision
```

#### Baseline Experiment:
The experiment involves training a single model in the conventional way.

Parameters: <br />
* ```Optimizer:```    : SGD 
* ```Learning Rate:``` 0.01

```Table 1:``` Test accuracy after training for 10 epochs:

| Model | Test Acc |
| ----- | -----    |
|  MLP  |  92.71%  |
|  CNN  |  98.42%  |

----

#### Federated Experiment:
The experiment involves training a global model in the federated setting.

Federated parameters (default values):
* ```Fraction of users (C)```: 0.1 
* ```Local Batch size  (B)```: 10 
* ```Local Epochs      (E)```: 10 本地迭代十次
* ```Optimizer            ```: SGD 
* ```Learning Rate        ```: 0.01 <br />

```Table 2:``` Test accuracy after training for 10 global epochs with:
全局聚合十次的结果
| Model |    IID   | Non-IID (equal)|
| ----- | -----    |----            |
|  MLP  |  88.38%  |     73.49%     |
|  CNN  |  97.28%  |     75.94%     |

最后在agg的时候，进行模型的测试，其他的暂时不进行测试
