package oracle

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"oraclebackend/xuperchain"
	"reflect"
)

const headerLength = 12

type HeaderMsg string

const (
	hRequest    HeaderMsg = "Request"
	hTrain HeaderMsg = "Train"
	hAgg    HeaderMsg = "Aggregate"
	hData   HeaderMsg = "Data"
	hCross	HeaderMsg = "Cross"
	hAggData HeaderMsg = "AggData"
	hAggCross HeaderMsg = "AggCross"
	hReply     HeaderMsg = "Commit"
)

type Msg interface {
	String() string
}

type NetMsg struct {
	Header           HeaderMsg
	Signature        []byte
	ClientNodePubkey *rsa.PublicKey
	ClientUrl        string
	RequestMsg       *RequestMsg
	TrainMsg    *TrainMsg
	AggMsg       *AggMsg
	DataMsg		*DataMsg
	AggDataMsg	*AggDataMsg
	CrossMsg 	*CrossMsg
	AggCrossMsg *AggCrossMsg
	ReplyMsg        *ReplyMsg
	ID 	string `json:"id"`
}

//<REQUEST, o, t, c>
type RequestMsg struct {
	Dataset string	`json:"dataset"`
	Model 	string 	`json:"model"`
	NIID 	int 	`json:"niid"`
	Global_Epoch int `json:"gloabl_epoch"`
	Local_Epoch int  `json:"local_epoch"`
	NodeCount int 	 `json:"node_count"`
	Groups []int 	`json:"groups"`
	Metadata xuperchain.Metadata 	`json:"metadata"`
	Type string `json:"type"`
}

func (msg RequestMsg) String() string {
	bmsg, _ := json.MarshalIndent(msg, "", "	")
	return string(bmsg) + "\n"
}

/**
  TrainMsg
  @Description: 执行训练任务的消息
**/
type TrainMsg struct {
	Dataset string	`json:"dataset"`
	Model 	string 	`json:"model"`
	NIID 	int 	`json:"niid"`
	Gloabl_Epoch int `json:"gloabl_epoch"`
	Local_Epoch int  `json:"local_epoch"`
	NodeCount int 	 `json:"node_count"`
	VerifyBls []byte `json:"verify_bls"`
	Groups []int 	`json:"groups"`
}

func (msg TrainMsg) String() string {
	bmsg, _ := json.MarshalIndent(msg, "", "	")
	return string(bmsg) + "\n"
}

/**
  AggMsg
  @Description: 执行参数聚合的消息
**/
type AggMsg struct {
	Global_Epoch int `json:"global_epoch"`
	Current_Epoch int `json:"current_epoch"`
	Local_Epoch   int `json:"local_epoch"`
	NIID 	int 	`json:"niid"`
	NodeID int 		`json:"node_id"`
	Dataset string	`json:"dataset"`
	Model 	string 	`json:"model"`
	NodeCount int 	 `json:"node_count"`
	BlsSig      []byte  `json:"bls_sk"`
	BlsPK      []byte  `json:"bls_pk"`
	Message    string	`json:"message"`
	Groups []int 	`json:"groups"`
}

func (msg AggMsg) String() string {
	bmsg, _ := json.MarshalIndent(msg, "", "	")
	return string(bmsg) + "\n"
}

/**
  DataMsg
  @Description: 执行数据获取的消息
**/
type DataMsg struct {
	Ip string `json:"ip"`
	Route string `json:"route"`
}

func (msg DataMsg) String() string {
	bmsg, _ := json.MarshalIndent(msg, "", "	")
	return string(bmsg) + "\n"
}

/**
  AggDataMsg
  @Description: 聚合数据的消息
**/
type AggDataMsg struct {
	NodeID     int 		`json:"node_id"`
	BlsSig      []byte  `json:"bls_sk"`
	BlsPK      []byte  `json:"bls_pk"`
	Message    string	`json:"message"`
}

func (msg AggDataMsg) String() string {
	bmsg, _ := json.MarshalIndent(msg, "", "	")
	return string(bmsg) + "\n"
}

/**
  CrossMsg
  @Description: 执行跨链请求的消息
**/
type CrossMsg struct {
	Ip string `json:"ip"`
	Route string `json:"route"`
}

func (msg CrossMsg) String() string {
	bmsg, _ := json.MarshalIndent(msg, "", "	")
	return string(bmsg) + "\n"
}

/**
  AggCrossMsg
  @Description: 执行跨链请求的消息
**/
type AggCrossMsg struct {
	NodeID     int 		`json:"node_id"`
	BlsSig      []byte  `json:"bls_sk"`
	BlsPK      []byte  `json:"bls_pk"`
	Message    string	`json:"message"`
}

func (msg AggCrossMsg) String() string {
	bmsg, _ := json.MarshalIndent(msg, "", "	")
	return string(bmsg) + "\n"
}

type ReplyMsg struct {
	ASig []byte `json:"a_sig"`
	Digest string `json:"digest"`
	PKs		[][]byte `json:"pks"`
	Msgs   []string `json:"msgs"`
	Type string `json:"type"`
	ID string `json:"id"`
}

func (msg ReplyMsg) String() string {
	bmsg, _ := json.MarshalIndent(msg, "", "	")
	return string(bmsg) + "\n"
}

func ComposeMsg(header HeaderMsg, payload interface{}, sig []byte) []byte {
	var bpayload []byte
	var err error
	t := reflect.TypeOf(payload)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.Struct:
		bpayload, err = json.Marshal(payload)
		if err != nil {
			panic(err)
		}
	case reflect.Slice:
		bpayload = payload.([]byte)
	default:
		panic(fmt.Errorf("not support type"))
	}

	b := make([]byte, headerLength)
	for i, h := range []byte(header) {
		b[i] = h
	}
	res := make([]byte, headerLength+len(bpayload)+len(sig))
	copy(res[:headerLength], b)
	copy(res[headerLength:], bpayload)
	if len(sig) > 0 {
		copy(res[headerLength+len(bpayload):], sig)
	}
	return res
}

func Send(data []byte, url string) error {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		return fmt.Errorf("%s is not online \n", url)
	}
	defer conn.Close()
	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("%v\n", err)
	}
	return nil
}

func SplitMsg(bmsg []byte) (HeaderMsg, []byte, []byte) {
	var header HeaderMsg
	var payload []byte
	var signature []byte
	hbyte := bmsg[:headerLength]
	hhbyte := make([]byte, 0)
	for _, h := range hbyte {
		if h != byte(0) {
			hhbyte = append(hhbyte, h)
		}
	}
	header = HeaderMsg(hhbyte)
	switch header {
	case hRequest, hAggData, hAggCross:
		payload = bmsg[headerLength : len(bmsg)-256]
		signature = bmsg[len(bmsg)-256:]
	case hReply, hTrain, hAgg, hCross, hData:
		payload = bmsg[headerLength:]
		signature = []byte{}
	}
	return header, payload, signature
}

//func printMsgLog(msg Msg) {
//	fmt.Println(msg.String())
//}

//func logHandleMsg(header HeaderMsg, msg Msg, from int) {
//	fmt.Printf("Receive %s msg from localhost:%d\n", header, nodeIdToPort(from))
//	printMsgLog(msg)
//}
//
//func logBroadcastMsg(header HeaderMsg, msg Msg) {
//	fmt.Printf("Send/broadcast %s msg \n", header)
//	printMsgLog(msg)
//}

