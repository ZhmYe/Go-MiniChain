# Go-MiniChain
### 实验一： 区块链系统简单实现

#### 一、简介

​	**本实验将参考比特币中的区块结构，使用Golang程序设计语言实现一个简单的区块链系统，以帮助读者更好地理解区块链的概念。本实验所实现的简易区块链系统名为minichain，该系统模拟比特币的挖矿过程，使用一个工作线程进行交易的打包、Merkle树根哈希值的计算以及相应的挖矿过程（随机替换nonce值，计算出满足挖矿难度条件的区块哈希值）。正确补全代码包中相应的功能函数后，运行主程序，minichain将产生新的区块。**

#### 二、TODO

​	在`consensus/MinerNode.go`中有三个`todo`

- `getBlockBody(transactions []Transaction)`
  - 根据传入的参数中的交易，构造并返回一个相应的区块体对象，还需要根据这些交易计算Merkle树的根哈希值

- `mine(blockBody BlockBody)`
  - 在循环中完成"挖矿"操作，其实就是通过不断的变换区块中的nonce字段，直至区块的哈希值满足难度条件，即可将该区块加入区块链中
- `getBlock(blockBody BlockBody)`
  - 该方法供mine方法调用，构造一个区块头对象，然后用一个区块对象组合区块头和区块体

#### 三、运行

- 在Goland下直接在`main.go`中运行main即可

- 在命令行下

  ```
  go run main.go
  ```
