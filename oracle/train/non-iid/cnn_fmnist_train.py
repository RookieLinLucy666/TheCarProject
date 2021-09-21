import os
import sys
import time
import numpy as np
import torch
from torchvision import transforms
from torchvision import datasets
from torch.utils.data import DataLoader, Dataset
import torch.nn.functional as F
import torch.optim as optim

# resume = False
resume = True
batch_size = 64
#Convert the PIL image to tensor
#0.13017,0.3081 是像素pixel的平均值和标准差

current_round = int(sys.argv[1])
epochs = int(sys.argv[2])
node_count = int(sys.argv[3])
node_id = int(sys.argv[4])

transform = transforms.Compose([
    transforms.ToTensor(),
    transforms.Normalize((0.1307,), (0.3081,))
])

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

class Swish(torch.nn.Module):
    def forward(self, input):
        return (input * torch.sigmoid(input))

    def __repr__(self):
        return self.__class__.__name__ + ' ()'

class Net(torch.nn.Module):
    def __init__(self):
        super(Net, self).__init__()
        self.layer1 = torch.nn.Sequential(
            torch.nn.Conv2d(1, 16, kernel_size=5, padding=2),
            torch.nn.BatchNorm2d(16),
            torch.nn.ReLU(),
            torch.nn.MaxPool2d(2))
        self.layer2 = torch.nn.Sequential(
            torch.nn.Conv2d(16, 32, kernel_size=5, padding=2),
            torch.nn.BatchNorm2d(32),
            torch.nn.ReLU(),
            torch.nn.MaxPool2d(2))
        self.fc = torch.nn.Linear(7*7*32, 10)

    def forward(self, x):
        out = self.layer1(x)
        out = self.layer2(out)
        out = out.view(out.size(0), -1)
        out = self.fc(out)
        return out

train_dataset = datasets.FashionMNIST(root='fmnist/',
                                      train=True,
                                      download=True,
                                      transform=transform)

# 60,000 training imgs -->  200 imgs/shard X 300 shards
num_shards, num_imgs = 200, 300
num_user_shards = int(num_shards / node_count)
idx_shard = [i for i in range(num_shards)]
train_idx = []
idxs = np.arange(num_shards*num_imgs)
labels = train_dataset.train_labels.numpy()

# sort labels
idxs_labels = np.vstack((idxs, labels))
idxs_labels = idxs_labels[:, idxs_labels[1, :].argsort()]
idxs = idxs_labels[0, :]

# divide and assign `num_user_shards` shards/client
if current_round == 1:
    rand_set = np.random.choice(idx_shard, num_user_shards, replace=False)
    # print(rand_set)
    np.savetxt('train/non-iid/fmnist_'+str(node_id), rand_set)
else:
    rand_set = np.loadtxt('train/non-iid/fmnist_'+str(node_id))
# idx_shard = list(set(idx_shard) - rand_set)
for rand in rand_set:
    train_idx = np.concatenate(
        (train_idx, idxs[int(rand)*num_imgs:(int(rand)+1)*num_imgs]), axis=0)

# step_train = int(len(train_dataset)/node_count)
# # train_idx = [i for i in range(node_id * step_train,  (node_id + 1) * step_train - int(0.2 * step_train))]
# # valid_idx = [i for i in range((node_id + 1) * step_train - int(0.2 * step_train), (node_id + 1) * step_train)]
# all_idx = [i for i in range(len(train_dataset))]
# # train_idx = np.random.choice(all_idx, int(len(train_dataset)/node_count),
# #                              replace=False)
# train_idx = train_idx = [i for i in range(node_id * step_train,  (node_id + 1) * step_train)]
#
# # idxs_train = train_idx[:int(0.8*len(train_idx))]
# # idxs_val = train_idx[int(0.8*len(train_idx)):int(0.9*len(train_idx))]
# # idxs_test = train_idx[int(0.9*len(train_idx)):]



train_loader = DataLoader(DatasetSplit(train_dataset, train_idx),
                          shuffle=True,
                          batch_size=batch_size)

# valid_loader = DataLoader(DatasetSplit(train_dataset, idxs_val),
#                          shuffle=True,
#                          batch_size=batch_size)
#
# test_loader = DataLoader(DatasetSplit(train_dataset, idxs_test),
#                           shuffle=True,
#                           batch_size=batch_size)

# if self.model_args.dataset == 'mnist' or self.model_args.dataset == 'cifar':
#     optimizer = torch.optim.SGD(net.parameters(), lr=self.model_args.lr, momentum=self.model_args.momentum)
# else:
#     optimizer = torch.optim.Adam(net.parameters(), lr=self.model_args.lr)

model = Net()
gpu_count = 4
cuda = ""
if node_id % gpu_count == 0:
    cuda = "cuda:0"
if node_id % gpu_count == 1:
    cuda = "cuda:1"
if node_id % gpu_count == 2:
    cuda = "cuda:2"
if node_id % gpu_count == 3:
    cuda = "cuda:3"
device = torch.device(cuda if torch.cuda.is_available() else "cpu")
model.to(device)#convert parameters and buffers of all modules to cuda tensor
criterion = torch.nn.CrossEntropyLoss()
optimizer = optim.Adam(model.parameters(),lr=0.01, )

if resume and current_round != 1:
    path_checkpoint = "train/non-iid/checkpoint/agg/fmnist_ckpt_" + str(current_round-1) + ".pth"  # 断点路径
    checkpoint = torch.load(path_checkpoint)  # 加载断点

    model.load_state_dict(checkpoint['net'])  # 加载模型可学习参数

    optimizer.load_state_dict(checkpoint['optimizer'])  # 加载优化器参数
    start_epoch = checkpoint['epoch']  # 设置开始的epoch

def train(epoch):
    running_loss = 0.0
    for batch_idx, data in enumerate(train_loader, 0):
        inputs, target = data
        inputs, target = inputs.to(device), target.to(device)
        #send the inputs and targets at every step to the GPU
        optimizer.zero_grad()

        outputs = model(inputs)
        loss = criterion(outputs, target)
        loss.backward()
        optimizer.step()

        # running_loss += loss.item()
        # if batch_idx % 100 == 99:
        #     print('[%d, %5d] loss: %.3f' % (epoch + 1, batch_idx + 1, running_loss / 100))
        #     running_loss = 0.0

    checkpoint = {
        "net": model.state_dict(),
        'optimizer':optimizer.state_dict(),
        "epoch": epoch
    }
    if not os.path.isdir("train/non-iid/checkpoint/fmnist/"):
        os.mkdir("train/non-iid/checkpoint/fmnist/")
    torch.save(checkpoint, 'train/non-iid/checkpoint/fmnist/ckpt_'+ str(node_id) + '_%s.pth' %(str(epoch)))

# def valid():
#     correct = 0
#     total = 0
#     with torch.no_grad():
#         for data in valid_loader:
#             inputs, target = data
#             inputs, target = inputs.to(device), target.to(device)
#             outputs = model(inputs)
#             _, predicted = torch.max(outputs.data, dim=1)#选择概率最大的输出
#             total += target.size(0)
#             correct += (predicted == target).sum().item()

    # print('Accuracy on valid set in node ' + str(node_id) +': %.4f %%' % (100 * float(correct) / float(total)))

# def test():
#     correct = 0
#     total = 0
#     with torch.no_grad():
#         for data in test_loader:
#             inputs, target = data
#             inputs, target = inputs.to(device), target.to(device)
#             outputs = model(inputs)
#             _, predicted = torch.max(outputs.data, dim=1)#选择概率最大的输出
#             total += target.size(0)
#             correct += (predicted == target).sum().item()

    # print('Accuracy on test set in node ' + str(node_id) +': %.4f %%' % (100 * float(correct) / float(total)))

if __name__ == '__main__':
    for epoch in range(epochs):
        train(epoch)
        # valid()
        # test()