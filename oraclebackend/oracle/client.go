package oracle

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bn256"
	"io/ioutil"
	"math"
	"net"
	"oraclebackend/models"
	"oraclebackend/xuperchain"
	"strconv"
	"sync"
	"time"
)

const (
	NIID = 1 //1非独立同分布，0独立同分布
	Malicious = 0 //恶意节点，在Demo中不考虑恶意节点
)

//  Client
//  @Description: 客户端，其角色用于连接区块链
//
type Client struct {
	nodeId     int
	Url        string
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
		Url:        ClientNode[i].url,
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
	//c.SendRequest()
	ln, err := net.Listen("tcp", c.Url)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		reply := c.HandleConnection(conn)
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
func (c *Client) HandleConnection(conn net.Conn) (reply bool) {
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
func (c *Client) SendRequest(dataType string, bcid string, demand xuperchain.FederatedAIDemand, id int64) {
	var reqmsg *RequestMsg

	switch dataType {
	case "compute":
		reqmsg = ListenCompute(bcid, demand)
	case "data":
		reqmsg = ListenData(bcid, id)
	case "cross":
		reqmsg = ListenCross(bcid)
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
		ClientUrl:        c.Url,
		ID: bcid,
	}
	marshalMsg, _ := json.Marshal(data)
	c.StartTime = time.Now()
	Send(marshalMsg, c.findPrimaryNode().url)
	//Send(ComposeMsg(hRequest, reqmsg, sig), c.findPrimaryNode().url)
	c.request = reqmsg
}

func ListenData(bcid string, id int64) *RequestMsg {
	xuperchain.InvokeQuery(bcid)
	//err := xuperchain.ListenQueryEvent()
	//if err != nil {
	//	fmt.Println(err)
	//	return nil
	//}
	metadata, _ := models.GetMetadataById(id)
	//fmt.Println(metadata)
	v := xuperchain.Metadata{
		Uploader: metadata.Uploader,
		Name:     metadata.Name,
		Type:     metadata.Type,
		Ip:       metadata.Ip,
		Route:    metadata.Route,
		Abstract: metadata.Abstract,
	}
	return &RequestMsg{
		NodeCount: NodeCount,
		NIID: NIID,
		Metadata: v,
		Type: v.Type,
	}
}

func ListenCross(id string) *RequestMsg{
	xuperchain.InvokeQuery(id)
	xuperchain.ListenQueryEvent()
	metadata, _ := xuperchain.GetVariable()
	//fmt.Println(metadata)

	return &RequestMsg{
		NodeCount: NodeCount,
		NIID: NIID,
		Metadata: metadata,
		Type: "cross",
	}
}

func ListenCompute(id string, demand xuperchain.FederatedAIDemand) *RequestMsg{
	xuperchain.InvokeComputingShare(id, demand.Model, demand.Dataset, demand.Round, demand.Epoch)
	//xuperchain.ListenComputingShareEvent()
	//fmt.Println("123")
	//metadata, learning := xuperchain.GetVariable()
	//
	round, _ := strconv.Atoi(demand.Round)
	epoch, _ := strconv.Atoi(demand.Epoch)

	return &RequestMsg{
		Dataset: demand.Dataset,
		Model: demand.Model,
		Global_Epoch: round,
		Local_Epoch: epoch,
		NodeCount: NodeCount,
		NIID: NIID,
		Metadata: xuperchain.Metadata{},
		Type: "compute",
	}
}

/**
  handleReply
  @Description: 处理返回的请求
  @receiver c
  @param payload
  @return bool
**/
func (c *Client) handleReply(payload []byte) bool {
	var replyMsg ReplyMsg
	err := json.Unmarshal(payload, &replyMsg)
	if err != nil {
		fmt.Printf("error happened:%v", err)
		return false
	}
	asig, _ := new(bn256.G1).Unmarshal(replyMsg.ASig)
	length := len(replyMsg.Msgs)
	pks := make([]*bn256.G2,length,length)
	for i:=0;i<length;i++{
		pks[i],_= new(bn256.G2).Unmarshal(replyMsg.PKs[i])
	}
	ok := AVerify(asig,replyMsg.Msgs,pks)
	if ok {
		if replyMsg.Type == "compute" {
			xuperchain.InvokeComputingCallBack(replyMsg.ID, replyMsg.Msgs[0], string(replyMsg.ASig), byte2string(replyMsg.PKs))
			fmt.Println("Finish Compute")
		} else if replyMsg.Type == "data" {
			xuperchain.InvokeQueryCallback(replyMsg.ID, replyMsg.Msgs[0], string(replyMsg.ASig), byte2string(replyMsg.PKs))
			fmt.Println("Finish Data")
		} else if replyMsg.Type == "cross" {
			xuperchain.InvokeComputingCallBack(replyMsg.ID, replyMsg.Msgs[0], string(replyMsg.ASig), byte2string(replyMsg.PKs))
			fmt.Println("Finish Cross")
		}
		r := models.Result{
			BcId: replyMsg.ID,
			Result: replyMsg.Result,
		}

		fmt.Println(replyMsg.Result)

		_, err := models.AddResult(&r)
		if err != nil {
			fmt.Println(err)
		}
		c.EndTime = time.Now()
		fmt.Println("Finish calculation.")
		fmt.Println(r.Result)
		return true
	}
	return false
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

