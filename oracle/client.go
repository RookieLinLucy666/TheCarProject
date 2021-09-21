package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"sync"
	"time"
)

//
//  Client
//  @Description: 客户端，其角色用于连接区块链
//
type Client struct {
	nodeId     int
	url        string
	keypair    Keypair
	knownNodes []*KnownNode
	request    *RequestMsg
	replyLog   map[int]*ReplyMsg
	mutex      sync.Mutex
	StartTime  time.Time
	EndTime    time.Time
}

func NewClient(i int32) *Client {
	client := &Client{
		nodeId:     ClientNode[i].nodeID,
		url:        ClientNode[i].url,
		keypair:    ClientKeypairMap[ClientNode[i].nodeID],
		knownNodes: KnownNodes,
		request:    nil,
		replyLog:   make(map[int]*ReplyMsg),
		mutex:      sync.Mutex{},
	}
	return client
}

/**
  Start
  @Description: 启动客户端，包括发送请求，建立TCP连接等
  @receiver c
**/
func (c *Client) Start() {
	c.sendRequest()
	ln, err := net.Listen("tcp", c.url)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		reply := c.handleConnection(conn)
		if reply {
			break
		}
	}
}

/**
  handleConnection
  @Description: 监听消息
  @receiver c
  @param conn
  @return reply
**/
func (c *Client) handleConnection(conn net.Conn) (reply bool) {
	req, err := ioutil.ReadAll(conn)
	header, payload, _ := SplitMsg(req)
	if err != nil {
		panic(err)
	}
	switch header {
	case hReply:
		reply = c.handleReply(payload)
	}
	return reply
}

/**
  sendRequest
  @Description: 发送请求
  @receiver c
**/
func (c *Client) sendRequest() {
	groups := c.generateMali()

	reqmsg := &RequestMsg{
		Dataset: Dataset,
		Model: Model,
		Global_Epoch: Global_Epoch,
		Local_Epoch: Local_Epoch,
		NodeCount: NodeCount,
		NIID: NIID,
		Groups: groups,
	}
	sig, err := c.signMessage(reqmsg)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	//logBroadcastMsg(hRequest, reqmsg)
	data := &NetMsg{
		Header:           hRequest,
		RequestMsg:       reqmsg,
		Signature:        sig,
		ClientNodePubkey: c.keypair.pubkey,
		ClientUrl:        c.url,
	}
	marshalMsg, _ := json.Marshal(data)
	c.StartTime = time.Now()
	Send(marshalMsg, c.findPrimaryNode().url)
	//Send(ComposeMsg(hRequest, reqmsg, sig), c.findPrimaryNode().url)
	c.request = reqmsg
}

/**
  handleReply
  @Description: 处理返回的请求
  @receiver c
  @param payload
  @return bool
**/
func (c *Client) handleReply(payload []byte) bool {
	c.EndTime = time.Now()
	fmt.Println("Finish calculation.")
	return true
}

/**
  generateMali
  @Description: 生成恶意节点，在Demo中恶意节点数量为0
  @receiver c
  @return []int
**/
func (c *Client) generateMali() []int{
	//测试5次
	nums := generateRandomNumber(1, NodeCount, int(math.Floor(NodeCount * Malicious)))
	return nums
}

/**
  signMessage
  @Description: 节点给消息签名
  @receiver c
  @param msg
  @return []byte
  @return error
**/
func (c *Client) signMessage(msg interface{}) ([]byte, error) {
	sig, err := signMessage(msg, c.keypair.privkey)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

/**
  findPrimaryNode
  @Description: 寻找预言机的广播节点
  @receiver c
  @return *KnownNode
**/
func (c *Client) findPrimaryNode() *KnownNode {
	nodeId := ViewID % len(c.knownNodes)
	for _, knownNode := range c.knownNodes {
		if knownNode.nodeID == nodeId {
			return knownNode
		}
	}
	return nil
}

func (c *Client) countTolerateFaultNode() int {
	return (len(c.knownNodes) - 1) / 3
}

func (c *Client) countNeedReceiveMsgAmount() int {
	f := c.countTolerateFaultNode()
	return f + 1
}
