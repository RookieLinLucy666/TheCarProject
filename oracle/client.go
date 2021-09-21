package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"sync"
	"time"
)

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

func (c *Client) handleReply(payload []byte) bool {
	c.EndTime = time.Now()
	fmt.Println("Finish calculation.")
	return true
}

func (c *Client) generateMali() []int{
	//测试5次
	nums := generateRandomNumber(1, NodeCount, int(math.Floor(NodeCount * Malicious)))
	return nums
}

func (c *Client) signMessage(msg interface{}) ([]byte, error) {
	sig, err := signMessage(msg, c.keypair.privkey)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

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
