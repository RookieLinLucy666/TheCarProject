import copy
import os
import sys

import torch
from torchvision import transforms
from torchvision import datasets
from torch.utils.data import DataLoader, Dataset
import torch.nn.functional as F
import torch.optim as optim


#python mnist_agg_iid.py  1 5 5

resume = False
# resume = True
batch_size = 64
# Convert the PIL image to tensor
# 0.13017,0.3081 是像素pixel的平均值和标准差

# python cnn_mnist_break.py 3 5 0
# python cnn_mnist_break.py 3 5 1
# python cnn_mnist_break.py 3 5 2
# python cnn_mnist_break.py 3 5 3
# python cnn_mnist_break.py 3 5 4

class DatasetSplit(Dataset):
    """An abstract Dataset class wrapped around Pytorch Dataset class.
    """

    def __init__(self, dataset, idxs):
        self.dataset = dataset
        self.idxs = [int(i) for i in idxs]

    def __len__(self):
        return len(self.idxs)

    def __getitem__(self, item):
        image, label = self.dataset[self.idxs[item]]
        return torch.tensor(image), torch.tensor(label)

class Net(torch.nn.Module):
    def __init__(self):
        super(Net, self).__init__()
        self.conv1 = torch.nn.Conv2d(1, 10, kernel_size=5)  # channel_in, channel_out
        self.conv2 = torch.nn.Conv2d(10, 20, kernel_size=5)
        self.pooling = torch.nn.MaxPool2d(2)  # 取最大，即大小缩小一半
        self.fc = torch.nn.Linear(320, 10)

    # 计算图参考PPT
    def forward(self, x):
        batch_size = x.size(0)
        x = F.relu(self.pooling(self.conv1(x)))
        x = F.relu(self.pooling(self.conv2(x)))
        x = x.view(batch_size, -1)  # flatten
        x = self.fc(x)
        return x

def average_weight(w):
    """
    Returns the average of the weights.
    """
    w_avg = copy.deepcopy(w[0])
    for key in w_avg.keys():
        for i in range(1, len(w)):
            w_avg[key] += w[i][key]
        w_avg[key] = torch.div(w_avg[key], len(w))
    return w_avg

def test():
    running_loss = 0.0
    correct = 0
    total = 0
    with torch.no_grad():
        for data in test_loader:
            inputs, target = data
            inputs, target = inputs.to(device), target.to(device)
            outputs = model(inputs)
            loss = criterion(outputs, target)
            _, predicted = torch.max(outputs.data, dim=1)#选择概率最大的输出
            total += target.size(0)
            correct += (predicted == target).sum().item()
            running_loss += loss.item()

    print('Accuracy on agg set in global_epoch ' + str(current_rounds) + ': %.4f %%' % (100 * float(correct) / float(total)))
    print('Loss on agg set in global_epoch ' + str(current_rounds) + ': %.4f' % (float(running_loss) / float(total)))

def valid():
    correct = 0
    total = 0
    with torch.no_grad():
        for data in valid_loader:
            inputs, target = data
            inputs, target = inputs.to(device), target.to(device)
            outputs = model(inputs)
            _, predicted = torch.max(outputs.data, dim=1)#选择概率最大的输出
            total += target.size(0)
            correct += (predicted == target).sum().item()
        return 100 * float(correct)/float(total)

current_rounds = int(sys.argv[1])
local_epoch = int(sys.argv[2])
node_count = int(sys.argv[3])
# node_id = int(sys.argv[4])

transform = transforms.Compose([
    transforms.ToTensor(),
    transforms.Normalize((0.1307,), (0.3081,))
])

model = Net()
device = torch.device("cpu")
# device = torch.device(
#     'cuda:{}'.format(node_id % gpu_count) if torch.cuda.is_available() else 'cpu')
# device = torch.device("cpu")
model.to(device)  # convert parameters and buffers of all modules to cuda tensor
criterion = torch.nn.CrossEntropyLoss()
optimizer = optim.SGD(model.parameters(), lr=0.01, momentum=0.5)

valid_dataset = datasets.MNIST(root='mnist/',
                              train=False,
                              download=True,
                              transform=transform)

all_idx = [i for i in range(len(valid_dataset))]
idxs_valid = all_idx[:int(len(all_idx)/3)]


valid_loader = DataLoader(DatasetSplit(valid_dataset, idxs_valid),
                         shuffle=True,
                         batch_size=batch_size)

local_correct = {}
local_weights = []
sort_weights = []
sort_node = []

for i in range(1, node_count):
    path_checkpoint = "train/non-iid/checkpoint/mnist/ckpt_" + str(i) + "_" + str(local_epoch-1) + ".pth"  # 断点路径
    checkpoint = torch.load(path_checkpoint)  # 加载断点
    model.load_state_dict(checkpoint['net']) # 加载模型可学习参数
    weight = model.state_dict()

    local_weights.append(copy.deepcopy(weight))
    correct = valid()
    local_correct[str(i)] = correct

local_correct = sorted(local_correct.items(), key = lambda kv:(kv[1], kv[0]), reverse=True)

# print("local", local_correct)

#用1/3数据集测试，选择表现较好的前1/3个节点的参数，可以用反证法证明，具体的证明可参考vitalik的证明
for element in local_correct:
    # print(int(element[0]))
    sort_weights.append(copy.deepcopy(local_weights[int(element[0])-1]))

global_weight = average_weight(sort_weights[:int(len(sort_weights)/3)])
# global_weight = average_weight(local_weights)
model.load_state_dict(global_weight)

checkpoint = {
    "net": model.state_dict(),
    'optimizer':optimizer.state_dict(),
    'epoch': 1,
}
if not os.path.isdir("train/non-iid/checkpoint/agg/"):
    os.mkdir("train/non-iid/checkpoint/agg/")
torch.save(checkpoint, 'train/non-iid/checkpoint/agg/mnist_ckpt_'+ str(current_rounds) + ".pth")

test_dataset = datasets.MNIST(root='mnist/',
                              train=False,
                              download=True,
                              transform=transform)

test_loader = DataLoader(test_dataset,
                         shuffle=True,
                         batch_size=batch_size)

test()
