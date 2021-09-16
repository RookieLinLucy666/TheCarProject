package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuperchain/contract-sdk-go/driver"
	"strconv"

	"github.com/xuperchain/contract-sdk-go/code"
)

type carFileAsset struct {}

const (
	CAR_FILE_ASSET_ID	=	"CarFileAssetId_"
	CAR_FILE_ASSET_COUNT = "CarFileAssetCount"
	TASK_QUEUE = "QueryQueue"
	META = "Meta"
	RESULT = "Result"
)

//元数据结构体
type metaData struct {
	Uploader	string `json:"uploader"`
	Name 		string `json:"name"`
	Type		string `json:"type"`
	Ip			string `json:"ip"`
	Route		string	`json:"route"`
	Abstract	string	`json:"abstract"`
}

//查询任务结构体
type queryTask struct {
	Id			string `json:"id"`
	MetaData	[]byte `json:"meta_data"`
}

//联邦学习需求结构体
type faderatedAIDemand struct {
	Model 		string `json:"model"`
	Dataset 	string `json:"dataset"`
	Round 		string `json:"round"`
	Epoch 		string `json:"epoch"`
}

//联邦学习数据结构体
type faderatedAIData struct {
	Id						string `json:"id"`
	MetaDataByte			[]byte `json:"meta_data_byte"`
	FaderatedAIDemandByte	[]byte `json:"faderated_ai_demand_byte"`
}

//计算共享任务返回结构体
type faderatedAIResult struct {
	ParamAddr		string
	ParamAbstract	string
	CorrectRate		string
}


//Initialize
/**
 * @Author Mengeshall
 * @Date 2021/9/14
 * @Description 初始化函数，对Id计数器、任务队列、元数据字典以及返回结果字典进行初始化
 */
func (cfa *carFileAsset) Initialize(ctx code.Context) code.Response{
	//初始化Id计数器
	if _, err := ctx.GetObject([]byte(CAR_FILE_ASSET_COUNT)); err != nil {
		if err := ctx.PutObject([]byte(CAR_FILE_ASSET_COUNT), []byte("0")); err != nil {
			return code.Error(err)
		}
	}
	//初始化任务队列
	if _, err := ctx.GetObject([]byte(TASK_QUEUE)); err != nil {
		q := initQueue()
		qStr, err := json.Marshal(q)
		if err != nil {
			return code.Error(err)
		}
		if err := ctx.PutObject([]byte(TASK_QUEUE), qStr); err != nil {
			return code.Error(err)
		}
	}
	//初始化元数据字典
	if _, err := ctx.GetObject([]byte(META)); err != nil {
		metaMap := map[string][]byte{}
		metaMapByte, err := json.Marshal(metaMap)
		if err != nil {
			return code.Error(err)
		}
		if err := ctx.PutObject([]byte(META), metaMapByte); err != nil {
			return code.Error(err)
		}
	}
	//初始化结果字典
	if _, err := ctx.GetObject([]byte(RESULT)); err != nil {
		resultMap := map[string][]byte{}
		resultMapByte, err := json.Marshal(resultMap)
		if err != nil {
			return code.Error(err)
		}
		if err := ctx.PutObject([]byte(META), resultMapByte); err != nil {
			return code.Error(err)
		}
	}
	return code.OK([]byte("initializing successfully"))
}


// Query
/**
 * @Author Mengeshall
 * @Date 2021/9/14
 * @Description 域内共享与跨域共享应用合约，负责根据Id查询元数据，并调用代理合约
 * 				——————验证部分未写——————
 */
func (cfa *carFileAsset) Query(ctx code.Context) code.Response{
	//接收参数
	args := struct {
		Id 			string `json:"id"`
		Inquirer	string `json:"inquirer"`
		Expiration	string `json:"expiration"`
	}{}
	//验证传入参数格式
	if err := code.Unmarshal(ctx.Args(), &args); err != nil {
		return code.Error(err)
	}
	//验证时间是否过期

	//根据Id查询元数据
	metaMapByte, err := ctx.GetObject([]byte(META))
	metaMap := map[string][]byte{}
	if err := json.Unmarshal(metaMapByte, &metaMap); err != nil {
		return code.Error(err)
	}
	data := metaMap[args.Id]
	if err != nil {
		return code.Error(err)
	}
	//调用代理合约进行查询
	task := queryTask{
		Id: args.Id,
		MetaData: data,
	}
	if err := cfa.QueryAgentAccept(ctx, task); err != nil {
		return code.Error(err)
	}
	return code.OK([]byte("call agent query contract successfully"))
}


//QueryAgentAccept
/**
 * @Author Mengeshall
 * @Date 2021/9/14
 * @Description 域内共享与跨域共享代理合约，负责将查询任务加入任务队列，并出发查询事件
 */
func (cfa *carFileAsset) QueryAgentAccept(ctx code.Context, task queryTask) error {
	//获取链上队列
	qByte, err := ctx.GetObject([]byte(TASK_QUEUE))
	if err != nil {
		return err
	}
	q := Queue{}
	if err := json.Unmarshal(qByte, &q); err != nil {
		return err
	}
	//新查询任务进队
	q.Enqueue(task.Id)
	//新队列上链
	qByte, err = json.Marshal(q)
	if err != nil {
		return err
	}
	if err := ctx.PutObject([]byte(TASK_QUEUE), qByte); err != nil {
		return err
	}
	//触发查询事件
	taskByte, err := json.Marshal(task)
	if err != nil {
		return err
	}
	if err := ctx.EmitEvent("queryEvent", taskByte); err != nil {
		return err
	}
	return nil
}


//QueryCallBack
/**
 * @Author Mengeshall
 * @Date 2021/9/15
 * @Description 域内共享与跨域共享回调合约，负责验证回调数据、回调结果上链、任务队列出队
				——————验证部分未写——————
 */
func (cfa *carFileAsset) QueryCallBack(ctx code.Context) code.Response {
	//接收参数
	args := struct {
		Id		string `json:"id"`
		Data	string `json:"data"`
		Asig	string `json:"asig"`
		Pks		string `json:"pks"`
	}{}
	if err := code.Unmarshal(ctx.Args(), &args); err !=nil {
		return code.Error(err)
	}
	//获取resultMap
	resultMapByte, err := ctx.GetObject([]byte(RESULT))
	if err != nil {
		return code.Error(err)
	}
	resultMap := map[string][]byte{}
	if err := json.Unmarshal(resultMapByte, &resultMap); err != nil {
		return code.Error(err)
	}
	//添加预言机返还结果
	resultMap[args.Id] = []byte(args.Data)
	//传回链上
	resultMapByte, err = json.Marshal(resultMap)
	if err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(RESULT), resultMapByte); err != nil {
		return code.Error(err)
	}
	//获取任务队列并出队
	qByte, err := ctx.GetObject([]byte(TASK_QUEUE))
	if err != nil {
		return code.Error(err)
	}
	q := Queue{}
	if err := json.Unmarshal(qByte, &q); err != nil {
		return code.Error(err)
	}
	//出队
	finishedId := fmt.Sprintf("%v", q.Dequeue())
	//新队列上链
	qByte, err = json.Marshal(q)
	if err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(TASK_QUEUE), qByte); err != nil {
		return code.Error(err)
	}
	return code.OK([]byte(finishedId))
}


//ComputingShare
/**
 * @Author Mengeshall
 * @Date 2021/9/15
 * @Description 计算共享应用合约，负责查询元数据，调用代理合约
 */
func (cfa *carFileAsset) ComputingShare(ctx code.Context) code.Response{
	//接收参数
	args := struct{
		Id			string `json:"id"`
		Model		string `json:"model"`
		Dataset		string `json:"dataset"`
		Round		string `json:"round"`
		Epoch		string `json:"epoch"`
	}{}
	if err := code.Unmarshal(ctx.Args(), &args); err != nil {
		return code.Error(err)
	}
	//根据Id查询元数据
	metaDataByte, err := ctx.GetObject([]byte(args.Id))
	if err != nil {
		return code.Error(err)
	}
	//构造传入代理合约数据格式
	faderated := faderatedAIDemand{
		Model: args.Model,
		Dataset: args.Dataset,
		Round: args.Round,
		Epoch: args.Epoch,
	}
	faderatedByte, err := json.Marshal(faderated)
	if err != nil {
		return code.Error(err)
	}
	data := faderatedAIData{
		Id: args.Id,
		MetaDataByte: metaDataByte,
		FaderatedAIDemandByte: faderatedByte,
	}
	//调用代理合约
	if err := cfa.ComputingShareAgent(ctx, data); err != nil {
		return code.Error(err)
	}
	return code.OK([]byte(args.Id))
}


//ComputingShareAgent
/**
 * @Author Mengeshall
 * @Date 2021/9/15
 * @Description 计算共享代理合约，负责任务队列入队、出发计算共享事件
 */
func (cfa *carFileAsset) ComputingShareAgent(ctx code.Context, data faderatedAIData) error {
	//获取链上队列
	qByte, err := ctx.GetObject([]byte(TASK_QUEUE))
	if err != nil {
		return err
	}
	q := Queue{}
	if err := json.Unmarshal(qByte, &q); err != nil {
		return err
	}
	//新任务入队
	q.Enqueue(data.Id)
	//任务队列上链
	qByte, err = json.Marshal(q)
	if err != nil {
		return err
	}
	if err := ctx.PutObject([]byte(TASK_QUEUE), qByte); err != nil {
		return err
	}
	//触发计算共享任务
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := ctx.EmitEvent("computingShareEvent", dataByte); err != nil {
		return err
	}
	return nil
}


//ComputingCallBack
/**
 * @Author Mengeshall
 * @Date 2021/9/16
 * @Description 计算共享回调合约，负责验证身份、结果上链、任务队列出队
				——————验证内容未编写——————
 */
func (cfa *carFileAsset) ComputingCallBack(ctx code.Context) code.Response {
	//接收参数
	args := struct {
		Id					string `json:"id"`
		FaderatedAIResult	string `json:"faderated_ai_result"`
		Asig				string `json:"asig"`
		Pks					string `json:"pks"`
	}{}
	if err := code.Unmarshal(ctx.Args(), &args); err != nil {
		return code.Error(err)
	}
	//结果上链
	resultMapByte, err := ctx.GetObject([]byte(RESULT))
	if err != nil {
		return code.Error(err)
	}
	resultMap := map[string][]byte{}
	if err := json.Unmarshal(resultMapByte, &resultMap); err != nil {
		return code.Error(err)
	}
	resultMap[args.Id] = []byte(args.FaderatedAIResult)
	resultMapByte, err = json.Marshal(resultMap)
	if err := ctx.PutObject([]byte(RESULT), resultMapByte); err != nil {
		return code.Error(err)
	}
	//任务队列出队
	//获取任务队列并出队
	qByte, err := ctx.GetObject([]byte(TASK_QUEUE))
	if err != nil {
		return code.Error(err)
	}
	q := Queue{}
	if err := json.Unmarshal(qByte, &q); err != nil {
		return code.Error(err)
	}
	//出队
	finishedId := fmt.Sprintf("%v", q.Dequeue())
	//新队列上链
	qByte, err = json.Marshal(q)
	if err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(TASK_QUEUE), qByte); err != nil {
		return code.Error(err)
	}
	return code.OK([]byte(finishedId))
}


//CreateCfa
/**
 * @Author Mengeshall
 * @Date 2021/9/13
 * @Description 负责进行元数据结构体上链
 */
func (cfa *carFileAsset) CreateCfa(ctx code.Context) code.Response{
	args := metaData{}
	//验证传入参数格式
	if err := code.Unmarshal(ctx.Args(), &args); err != nil {
		return code.Error(err)
	}
	//获取id数量并加1
	countByte, err := ctx.GetObject([]byte(CAR_FILE_ASSET_COUNT))
	if err != nil {
		return code.Error(err)
	}
	count, _ := strconv.Atoi(string(countByte))
	newCount := count +1
	newId := CAR_FILE_ASSET_ID + strconv.Itoa(newCount)
	//json-->string-->[]byte
	data := metaData{
		Uploader: 	args.Uploader,
		Name: 		args.Name,
		Type:		args.Type,
		Ip:			args.Ip,
		Route: 		args.Route,
		Abstract: 	args.Abstract,
	}
	dataByte, _ := json.Marshal(data)
	//上传数据
	metaMapByte, err := ctx.GetObject([]byte(META))
	if err != nil {
		return code.Error(err)
	}
	metaMap := map[string][]byte{}
	metaMap[newId] = dataByte
	metaMapByte, err = json.Marshal(metaMap)
	if err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(newId), metaMapByte); err != nil {
		return code.Error(err)
	}
	//更新id数量
	if err := ctx.PutObject([]byte(CAR_FILE_ASSET_COUNT), []byte(strconv.Itoa(newCount))); err != nil {
		return code.Error(err)
	}
	return code.OK([]byte(strconv.Itoa(newCount)))
}


//UpdateCfa
/**
 * @Author Mengeshall
 * @Date 2021/9/13
 * @Description 负责元数据修改，只有上传者有权限修改元数据
 */
func (cfa *carFileAsset) UpdateCfa(ctx code.Context) code.Response {
	args := struct {
		Id			string `json:"id"`
		Operator	string `json:"operator"`
		MetaData 	string `json:"data"`
	}{}
	newData := metaData{}
	if err := code.Unmarshal(ctx.Args(), &args); err != nil {
		return code.Error(err)
	}
	if err := json.Unmarshal([]byte(args.MetaData), &newData); err != nil {
		return code.Error(err)
	}
	//若操作者不为上传者，无权更新cfa
	if args.Operator != newData.Uploader {
		return code.Error(code.ErrPermissionDenied)
	}
	newDataByte, _ := json.Marshal(newData)
	metaMapByte, err := ctx.GetObject([]byte(META))
	if err != nil {
		return code.Error(err)
	}
	metaMap := map[string][]byte{}
	if err := json.Unmarshal(metaMapByte, &metaMap); err != nil {
		return code.Error(err)
	}
	metaMap[args.Id] = newDataByte
	metaMapByte, err = json.Marshal(metaMap)
	if err != nil {
		return code.Error(err)
	}
	if err := ctx.PutObject([]byte(META), metaMapByte); err != nil {
		return code.Error(err)
	}
	return code.OK([]byte(args.Id))
}

func main() {
	driver.Serve(new(carFileAsset))
}
