package main

import (
	"bytes"
	"crypto/rsa"
	"fmt"
	"golang.org/x/crypto/bn256"
	"math/big"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

type Node struct {
	NodeID      int
	knownNodes  []*KnownNode
	sequenceID  int
	View        int
	msgQueue    chan []byte
	keypair     Keypair
	msgLog      *MsgLog
	requestPool map[string]*RequestMsg
	mutex       sync.Mutex
	global		int
	blsSK      *big.Int
	blsPK      *bn256.G2
	blslog      map[int]*BlsLog
	group 		int
}

type Keypair struct {
	privkey *rsa.PrivateKey
	pubkey  *rsa.PublicKey
}

type MsgLog struct {
	aggLog map[int]map[int]bool
}


type BlsLog struct {
	sigs []*bn256.G1
	pks  []*bn256.G2
	msgs []string
}

type VerifyBls struct {
	Asig []byte
	Msgs []string
	Pks [][]byte
}

func NewNode(nodeID int) *Node {
	blsSK,blsPK,_,_:= KeyGenerate()
	return &Node{
		nodeID,
		KnownNodes,
		0,
		ViewID,
		make(chan []byte),
		KnownKeypairMap[nodeID],
		&MsgLog{
			make(map[int]map[int]bool),
		},
		make(map[string]*RequestMsg),
		sync.Mutex{},
		0,
		blsSK,
		blsPK,
		make(map[int]*BlsLog),
		0,
	}
}

func (node *Node) getSequenceID() int {
	seq := node.sequenceID
	node.sequenceID++
	return seq
}

func (node *Node) Start() {
	go node.handleMsg()
}

func (node *Node) handleMsg() {
	for {
		msg := <-node.msgQueue
		netMsg := NetMsg{}
		json.Unmarshal(msg, &netMsg)
		//header, payload, sign := SplitMsg(msg)
		switch netMsg.Header {
		case hRequest:
			node.handleRequest(netMsg.RequestMsg, netMsg.Signature, netMsg.ClientNodePubkey, netMsg.ClientUrl)
		case hTrain:
			node.handleTrain(netMsg.TrainMsg, netMsg.Signature, netMsg.ClientUrl)
		case hAgg:
			node.handleAgg(netMsg.AggMsg, netMsg.Signature, netMsg.ClientUrl)
		}
	}
}

func (node *Node) handleRequest(request *RequestMsg, sig []byte, clientNodePubkey *rsa.PublicKey, clientNodeUrl string) {
	var trainMsg TrainMsg

	trainMsg = TrainMsg{
		Dataset: request.Dataset,
		Model: request.Model,
		Gloabl_Epoch: request.Global_Epoch,
		Local_Epoch: request.Local_Epoch,
		NodeCount: request.NodeCount,
		NIID: request.NIID,
		Groups: request.Groups,
	}

	msgSig, err := node.signMessage(trainMsg)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	data := &NetMsg{
		Header:        hTrain,
		TrainMsg: &trainMsg,
		Signature:     msgSig,
		ClientUrl:     clientNodeUrl,
	}
	marshalMsg, _ := json.Marshal(data)
	node.broadcast(marshalMsg)
}

func (node *Node) handleTrain(trainMsg *TrainMsg, sig []byte, clientNodeUrl string) {
	var aggMsg AggMsg

	if isContain(trainMsg.Groups, node.NodeID) {
		node.group = 1
	}

	primaryID := node.findPrimaryNode()
	msgPubkey, primary_url := node.findNodePubkey(primaryID)
	if msgPubkey == nil {
		fmt.Println("Can not find primary node's public key")
	}

	_, err := verifySignatrue(trainMsg, sig, msgPubkey)
	if err != nil {
		fmt.Println("Verify signature failed: %v\n", err)
		return
	}

	node.mutex.Lock()
	global := node.global
	node.mutex.Unlock()

	if global < trainMsg.Gloabl_Epoch {
		if global > 0 {
			verifiedJson := trainMsg.VerifyBls
			var verified VerifyBls
			err = json.Unmarshal(verifiedJson,&verified)
			if err != nil {
				fmt.Printf("2/error happened:%v", err)
				return
			}
			asig,_ := new(bn256.G1).Unmarshal(verified.Asig)
			length := len(verified.Msgs)
			pks := make([]*bn256.G2,length,length)
			for i:=0;i<length;i++{
				pks[i],_= new(bn256.G2).Unmarshal(verified.Pks[i])
			}
			ok := AVerify(asig,verified.Msgs,pks)
			if !ok {
				fmt.Println("Aggregation error.")
				return
			} else {
				//fmt.Println("Aggregation success.")
			}
		}

		node.global += 1
		//fmt.Println("global:", node.global)
		//进行一轮本地训练
		var out bytes.Buffer
		var stderr bytes.Buffer
		var cmd *exec.Cmd

		if node.group == 0{
			if node.NodeID / 5 == 1{
				time.Sleep(time.Duration(Duration)*time.Second)
			}
			if node.NodeID / 5 == 2 {
				time.Sleep(time.Duration(2 * Duration)*time.Second)
			}
			if node.NodeID / 5 == 3 {
				time.Sleep(time.Duration(3 * Duration)*time.Second)
			}
			if node.NodeID / 5 == 4 {
				time.Sleep(time.Duration(4 * Duration)*time.Second)
			}
			if node.NodeID / 5 == 5 {
				time.Sleep(time.Duration(5 * Duration)*time.Second)
			}
			if trainMsg.NIID == 0 {
				cmd = exec.Command("python", "train/iid/"+ trainMsg.Model + "_" + trainMsg.Dataset + "_train.py",
					strconv.Itoa(node.global), strconv.Itoa(trainMsg.Local_Epoch), strconv.Itoa(trainMsg.NodeCount), strconv.Itoa(node.NodeID))
			} else {
				cmd = exec.Command("python", "train/non-iid/"+ trainMsg.Model + "_" + trainMsg.Dataset + "_train.py",
					strconv.Itoa(node.global), strconv.Itoa(trainMsg.Local_Epoch), strconv.Itoa(trainMsg.NodeCount), strconv.Itoa(node.NodeID))
			}

			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			}
			fmt.Println(out.String())

			if err != nil{
				fmt.Println("Training error ",err)
			}
		} else {
			//fmt.Println("Malicious Node execution: "+ strconv.Itoa(node.NodeID))
			if trainMsg.NIID == 0 {
				cmd = exec.Command("python", "train/iid/"+ trainMsg.Model + "_" + trainMsg.Dataset + "_mali.py",
				strconv.Itoa(node.global), strconv.Itoa(trainMsg.Local_Epoch), strconv.Itoa(trainMsg.NodeCount), strconv.Itoa(node.NodeID))
			} else {
				cmd = exec.Command("python", "train/non-iid/"+ trainMsg.Model + "_" + trainMsg.Dataset + "_mali.py",
				strconv.Itoa(node.global), strconv.Itoa(trainMsg.Local_Epoch), strconv.Itoa(trainMsg.NodeCount), strconv.Itoa(node.NodeID))
			}


			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			}
			fmt.Println(out.String())

			if err != nil{
				fmt.Println("Training error ",err)
			}
		}

		//filepath := "train/iid/train_log_" + strconv.Itoa(node.NodeID) + ".txt"
		//message := ReadFile(filepath)

		blssig := Sign(node.blsSK, out.String())

		aggMsg = AggMsg{
			Global_Epoch: trainMsg.Gloabl_Epoch,
			Current_Epoch: node.global,
			NodeID:       node.NodeID,
			NIID: trainMsg.NIID,
			Dataset: trainMsg.Dataset,
			Model: trainMsg.Model,
			NodeCount: trainMsg.NodeCount,
			Local_Epoch: trainMsg.Local_Epoch,
			BlsSig: blssig.Marshal(),
			BlsPK: node.blsPK.Marshal(),
			Message: out.String(),
			Groups: trainMsg.Groups,
		}

		msgSig, err := node.signMessage(aggMsg)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		data := &NetMsg{
			Header:        hAgg,
			AggMsg: &aggMsg,
			Signature:     msgSig,
			ClientUrl:     clientNodeUrl,
		}
		marshalMsg, _ := json.Marshal(data)
		Send(marshalMsg, primary_url)
	}
}

func (node *Node) handleAgg(aggMsg *AggMsg, sig []byte, clientNodeUrl string) {

	global_epoch := aggMsg.Global_Epoch
	current_epoch := aggMsg.Current_Epoch

	node.mutex.Lock()
	if node.msgLog.aggLog[current_epoch] == nil {
		node.msgLog.aggLog[current_epoch] = make(map[int]bool)
	}
	node.msgLog.aggLog[current_epoch][aggMsg.NodeID] = true
	blssig, _ := new(bn256.G1).Unmarshal(aggMsg.BlsSig)
	blspk, _ := new(bn256.G2).Unmarshal(aggMsg.BlsPK)
	if node.blslog[current_epoch] == nil {
		node.blslog[current_epoch] = &BlsLog{
			sigs: nil,
			pks:  nil,
			msgs: nil,
		}
	}
	node.blslog[current_epoch].msgs = append(node.blslog[current_epoch].msgs, string(aggMsg.Message))
	node.blslog[current_epoch].sigs = append(node.blslog[current_epoch].sigs, blssig)
	node.blslog[current_epoch].pks = append(node.blslog[current_epoch].pks, blspk)
	node.mutex.Unlock()

	sum := node.findAggMsgCount(current_epoch)
	N := sum
	pks := make([][]byte, N, N)
	msgs := make([]string, N, N)
	sigs := make([]*bn256.G1, N, N)

	for i := 0; i < N; i++ {
		pks[i] = node.blslog[current_epoch].pks[i].Marshal()
		msgs[i] = node.blslog[current_epoch].msgs[i]
		sigs[i] = node.blslog[current_epoch].sigs[i]
	}
	asig := Aggregate(sigs)

	if sum == NodeCount -1 {
		//进行一轮参数聚合过程
		//Done:（在python文件中修改聚合的过程，即聚合的时候需要先拿出1/3的数据对节点提供的参数有效性进行判断）: 选择2/3的参数聚合，并下放。在这个过程中，可能存在leader节点故意不聚合参数，可采用bls签名，让leader不得不验证签名并提供参数。
		//Done: 假设leader节点都是彻底的恶意节点，即恶意节点并不会真正聚合签名，进而其他节点无法相信该节点得出的聚合结果。
		// cmd := exec.Command("python", "train/iid/"+ aggMsg.Model + "_" + aggMsg.Dataset + "_agg.py",
		// 	strconv.Itoa(current_epoch), strconv.Itoa(aggMsg.Local_Epoch), strconv.Itoa(aggMsg.NodeCount))
		var cmd *exec.Cmd
		if aggMsg.NIID == 0{
			cmd = exec.Command("python", "train/iid/"+ aggMsg.Model + "_" + aggMsg.Dataset + "_agg.py",
				strconv.Itoa(current_epoch), strconv.Itoa(aggMsg.Local_Epoch), strconv.Itoa(aggMsg.NodeCount))
		} else {
			cmd = exec.Command("python", "train/non-iid/"+ aggMsg.Model + "_" + aggMsg.Dataset + "_agg.py",
				strconv.Itoa(current_epoch), strconv.Itoa(aggMsg.Local_Epoch), strconv.Itoa(aggMsg.NodeCount))
		}

		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		}
		fmt.Println(out.String())

		if current_epoch ==  global_epoch {
			var replyMsg ReplyMsg
			replyMsg = ReplyMsg{
				Digest: out.String(),
			}

			Send(ComposeMsg(hReply, replyMsg, []byte{}), clientNodeUrl)
		} else {
			var trainMsg TrainMsg

			verifyMsg := VerifyBls{
				Asig: asig.Marshal(),
				Msgs: msgs,
				Pks: pks,
			}
			verifyMsgJson, _ := json.Marshal(verifyMsg)

			trainMsg = TrainMsg{
				Dataset: aggMsg.Dataset,
				Model: aggMsg.Model,
				Gloabl_Epoch: aggMsg.Global_Epoch,
				Local_Epoch: aggMsg.Local_Epoch,
				NodeCount: aggMsg.NodeCount,
				NIID: aggMsg.NIID,
				VerifyBls: verifyMsgJson,
				Groups: aggMsg.Groups,
			}

			msgSig, err := node.signMessage(trainMsg)
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}

			data := &NetMsg{
				Header:        hTrain,
				TrainMsg: &trainMsg,
				Signature:     msgSig,
				ClientUrl:     clientNodeUrl,
			}
			marshalMsg, _ := json.Marshal(data)
			node.broadcast(marshalMsg)		}

	} else {
		//fmt.Println("Wait for training results.")
	}

}

func (node *Node) findAggMsgCount(current_epoch int) (int) {
	sum := 0
	node.mutex.Lock()
	for _, exist := range node.msgLog.aggLog[current_epoch] {
		if exist == true {
			sum ++
		}
	}
	node.mutex.Unlock()
	return sum
}

//func (node *Node) findVerifiedCommitMsgCount(digest string) (int, error) {
//	sum := 0
//	node.mutex.Lock()
//	for _, exist := range node.msgLog.commitLog[digest] {
//
//		if exist == true {
//			sum++
//		}
//	}
//	node.mutex.Unlock()
//	return sum, nil
//}

//func (node *Node) handlePrePrepare(prePrepareMsg *PrePrepareMsg, sig []byte, clientNodeUrl string) {
//	pnodeId := node.findPrimaryNode()
//	//logHandleMsg(hPrePrepare, prePrepareMsg, pnodeId)
//	msgPubkey := node.findNodePubkey(pnodeId)
//	if msgPubkey == nil {
//		fmt.Println("can't find primary node's public key")
//		return
//	}
//	// verify msg's signature
//	_, err := verifySignatrue(prePrepareMsg, sig, msgPubkey)
//	if err != nil {
//		fmt.Printf("verify signature failed:%v\n", err)
//		return
//	}
//
//	// verify prePrepare's digest is equal to request's digest
//	if prePrepareMsg.Digest != prePrepareMsg.Request.CRequest.Digest {
//		fmt.Printf("verify digest failed\n")
//		return
//	}
//	node.mutex.Lock()
//	node.requestPool[prePrepareMsg.Request.CRequest.Digest] = &prePrepareMsg.Request
//	node.mutex.Unlock()
//	err = node.verifyRequestDigest(prePrepareMsg.Digest)
//	if err != nil {
//		fmt.Printf("%v\n", err)
//		return
//	}
//	// put preprepare's msg into log
//	node.mutex.Lock()
//	if node.msgLog.preprepareLog[prePrepareMsg.Digest] == nil {
//		node.msgLog.preprepareLog[prePrepareMsg.Digest] = make(map[int]bool)
//	}
//	node.msgLog.preprepareLog[prePrepareMsg.Digest][pnodeId] = true
//	node.mutex.Unlock()
//	prepareMsg := PrepareMsg{
//		prePrepareMsg.Digest,
//		ViewID,
//		prePrepareMsg.SequenceID,
//		node.NodeID,
//	}
//	// sign prepare msg
//	msgSig, err := signMessage(prepareMsg, node.keypair.privkey)
//	if err != nil {
//		fmt.Printf("%v\n", err)
//		return
//	}
//	data := &NetMsg{
//		Header:     hPrepare,
//		PrepareMsg: &prepareMsg,
//		Signature:  msgSig,
//		ClientUrl:  clientNodeUrl,
//	}
//	marshalMsg, _ := json.Marshal(data)
//	node.mutex.Lock()
//	// put prepare msg into log
//	if node.msgLog.prepareLog[prepareMsg.Digest] == nil {
//		node.msgLog.prepareLog[prepareMsg.Digest] = make(map[int]bool)
//	}
//	node.msgLog.prepareLog[prepareMsg.Digest][node.NodeID] = true
//	node.mutex.Unlock()
//	//logBroadcastMsg(hPrepare, prepareMsg)
//	node.broadcast(marshalMsg)
//}

//func (node *Node) handlePrepare(prepareMsg *PrepareMsg, sig []byte, clientNodeUrl string) {
//	//logHandleMsg(hPrepare, prepareMsg, prepareMsg.NodeID)
//	// verify prepareMsg
//	pubkey := node.findNodePubkey(prepareMsg.NodeID)
//	_, err := verifySignatrue(prepareMsg, sig, pubkey)
//	if err != nil {
//		fmt.Printf("verify signature failed:%v\n", err)
//		return
//	}
//	// verify request's digest
//	err = node.verifyRequestDigest(prepareMsg.Digest)
//	if err != nil {
//		fmt.Printf("%v\n", err)
//		return
//	}
//	// verify prepareMsg's digest is equal to preprepareMsg's digest
//	pnodeId := node.findPrimaryNode()
//	exist := node.msgLog.preprepareLog[prepareMsg.Digest][pnodeId]
//	if !exist {
//		fmt.Printf("this digest's preprepare msg by %d not existed\n", pnodeId)
//		return
//	}
//	// put prepareMsg into log
//	node.mutex.Lock()
//	if node.msgLog.prepareLog[prepareMsg.Digest] == nil {
//		node.msgLog.prepareLog[prepareMsg.Digest] = make(map[int]bool)
//	}
//	node.msgLog.prepareLog[prepareMsg.Digest][prepareMsg.NodeID] = true
//	node.mutex.Unlock()
//	// if receive prepare msg >= 2f +1, then broadcast commit msg
//	limit := node.countNeedReceiveMsgAmount()
//	sum, err := node.findVerifiedPrepareMsgCount(prepareMsg.Digest)
//	if err != nil {
//		fmt.Printf("error happened:%v", err)
//		return
//	}
//	if sum >= limit {
//		// if already Send commit msg, then do nothing
//		node.mutex.Lock()
//		exist, _ := node.msgLog.commitLog[prepareMsg.Digest][node.NodeID]
//		node.mutex.Unlock()
//		if exist != false {
//			return
//		}
//		//Send commit msg
//		commitMsg := CommitMsg{
//			prepareMsg.Digest,
//			prepareMsg.ViewID,
//			prepareMsg.SequenceID,
//			node.NodeID,
//		}
//		sig, err := node.signMessage(commitMsg)
//		if err != nil {
//			fmt.Printf("sign message happened error:%v\n", err)
//		}
//		data := &NetMsg{
//			Header:    hCommit,
//			CommitMsg: &commitMsg,
//			Signature: sig,
//			ClientUrl: clientNodeUrl,
//		}
//		marshalMsg, _ := json.Marshal(data)
//		// put commit msg to log
//		node.mutex.Lock()
//		if node.msgLog.commitLog[commitMsg.Digest] == nil {
//			node.msgLog.commitLog[commitMsg.Digest] = make(map[int]bool)
//		}
//		node.msgLog.commitLog[commitMsg.Digest][node.NodeID] = true
//		node.mutex.Unlock()
//		//logBroadcastMsg(hCommit, commitMsg)
//		node.broadcast(marshalMsg)
//	}
//}

//func (node *Node) handleCommit(commitMsg *CommitMsg, sig []byte, clientNodeUrl string) {
//	//logHandleMsg(hCommit, commitMsg, commitMsg.NodeID)
//	//verify commitMsg's signature
//	msgPubKey := node.findNodePubkey(commitMsg.NodeID)
//	verify, err := verifySignatrue(commitMsg, sig, msgPubKey)
//	if err != nil {
//		fmt.Printf("verify signature failed:%v\n", err)
//		return
//	}
//	if verify == false {
//		fmt.Printf("verify signature failed\n")
//		return
//	}
//	// verify request's digest
//	err = node.verifyRequestDigest(commitMsg.Digest)
//	if err != nil {
//		fmt.Printf("%v\n", err)
//		return
//	}
//	// put commitMsg into log
//	node.mutex.Lock()
//	if node.msgLog.commitLog[commitMsg.Digest] == nil {
//		node.msgLog.commitLog[commitMsg.Digest] = make(map[int]bool)
//	}
//	node.msgLog.commitLog[commitMsg.Digest][commitMsg.NodeID] = true
//	node.mutex.Unlock()
//	// if receive commit msg >= 2f +1, then Send reply msg to client
//	limit := node.countNeedReceiveMsgAmount()
//	sum, err := node.findVerifiedCommitMsgCount(commitMsg.Digest)
//	if err != nil {
//		fmt.Printf("error happened:%v", err)
//		return
//	}
//	if sum >= limit {
//		// if already Send reply msg, then do nothing
//		node.mutex.Lock()
//		exist := node.msgLog.replyLog[commitMsg.Digest]
//		node.mutex.Unlock()
//		if exist == true {
//			return
//		}
//		// Send reply msg
//		node.mutex.Lock()
//		requestMsg := node.requestPool[commitMsg.Digest]
//		node.mutex.Unlock()
//		//fmt.Printf("operstion:%s  message:%s executed... \n", requestMsg.Operation, requestMsg.CRequest.Message)
//		done := fmt.Sprintf("operstion:%s  message:%s done ", requestMsg.Operation, requestMsg.CRequest.Message)
//		replyMsg := ReplyMsg{
//			node.View,
//			int(time.Now().Unix()),
//			requestMsg.ClientID,
//			node.NodeID,
//			done,
//		}
//		//logBroadcastMsg(hReply, replyMsg)
//		Send(ComposeMsg(hReply, replyMsg, []byte{}), clientNodeUrl)
//		node.mutex.Lock()
//		node.msgLog.replyLog[commitMsg.Digest] = true
//		node.mutex.Unlock()
//	}
//}

func (node *Node) verifyRequestDigest(digest string) error {
	node.mutex.Lock()
	_, ok := node.requestPool[digest]
	if !ok {
		node.mutex.Unlock()
		return fmt.Errorf("verify request digest failed\n")

	}
	node.mutex.Unlock()
	return nil
}

//func (node *Node) findVerifiedPrepareMsgCount(digest string) (int, error) {
//	sum := 0
//	node.mutex.Lock()
//	for _, exist := range node.msgLog.prepareLog[digest] {
//		if exist == true {
//			sum++
//		}
//	}
//	node.mutex.Unlock()
//	return sum, nil
//}

//func (node *Node) findVerifiedCommitMsgCount(digest string) (int, error) {
//	sum := 0
//	node.mutex.Lock()
//	for _, exist := range node.msgLog.commitLog[digest] {
//
//		if exist == true {
//			sum++
//		}
//	}
//	node.mutex.Unlock()
//	return sum, nil
//}

func (node *Node) broadcast(data []byte) {
	for _, knownNode := range node.knownNodes {
		if knownNode.nodeID != node.NodeID {
			err := Send(data, knownNode.url)
			if err != nil {
				fmt.Printf("%v", err)
			}
		}
	}

}

func (node *Node) findNodePubkey(nodeId int) (*rsa.PublicKey, string) {
	for _, knownNode := range node.knownNodes {
		if knownNode.nodeID == nodeId {
			return knownNode.pubkey, knownNode.url
		}
	}
	return nil, ""
}

func (node *Node) signMessage(msg interface{}) ([]byte, error) {
	sig, err := signMessage(msg, node.keypair.privkey)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (node *Node) findPrimaryNode() int {
	return ViewID % len(node.knownNodes)
}

func (node *Node) countTolerateFaultNode() int {
	return (len(node.knownNodes) - 1) / 3
}

func (node *Node) countNeedReceiveMsgAmount() int {
	f := node.countTolerateFaultNode()
	return 2*f + 1
}
