pragma solidity ^0.4.25;
import "./ComputingShare.sol";

contract ComputingShare{
    address oracleAddress;

    struct metaData{
        string source;
        string dataAbstract;
        uint dataType;
        bool isValid;
    }

    struct computingResult {
        string paramAddr;
        string paramAbstract;
        string correctRate;
        string data;
    }

    event computingShareEvent(
        uint id,
        string dataAbstract,
        string model,
        string dataset,
        string round,
        string epoch
    );

    event dataShareEvent(
        uint id,
        string source,
        string dataAbstract
    );

    mapping(uint => metaData) meta;
    mapping(uint => computingResult) result;
    uint count;

    constructor() public {
        count = 0;
    }

    function initledger() public returns(string){
        count = 0;
        return "initledger success";
    }

    /**
     * Author Mengeshall
     * Date 2021/10/13
     * Description 工具函数，添加metadata
     */
    function addmetadata(uint id, string _abstract, string _source, uint _type) public returns (string) {
        if(meta[id].isValid){
            return meta[id].dataAbstract;
        }
        meta[id].dataAbstract = _abstract;
        meta[id].dataType = _type;
        meta[id].source = _source;
        meta[id].isValid = true;
        return meta[id].dataAbstract;
    }

    /**
     * Author Mengeshall
     * Date 2021/10/13
     * Description 工具函数，删除metadata
     */
    function deletemetadata(uint id) public returns (bool) {
        if(meta[id].isValid){
            meta[id].dataAbstract = '';
            meta[id].source = '';
            meta[id].dataType = 0;
            meta[id].isValid = false;
            return true;
        }
        return false;
    }

    /**
     * Author Mengeshall
     * Date 2021/10/13
     * Description 工具函数，查询metadata
     */
    function getmetadata(uint id) public returns (string, uint, string, bool) {
        return (meta[id].dataAbstract, meta[id].dataType, meta[id].source, meta[id].isValid);
    }

    /**
     * Author Mengeshall
     * Date 2021/10/14
     * Description 工具函数，查询result
     */
    function getresult(uint id) public returns (string, string, string, string) {
        return (result[id].paramAddr, result[id].paramAbstract, result[id].correctRate, result[id].data);
    }

    /**
     * Author Mengeshall
     * Date 2021/10/13
     * Description 应用合约，负责上传metadata
     */
    function createmetadata(string dataAbstract, string source, uint dataType) public returns(uint){
        addmetadata(count, dataAbstract, source, dataType);
        count++;
        return count-1;
    }

    /**
     * Author Mengeshall
     * Date 2021/10/13
     * Description 计算共享合约，负责出发计算共享事件
     */
    function computingshare(uint id, string model, string dataset, string round, string epoch) public returns(string){
        //验证是否有该id对应的metadata
        require(meta[id].isValid != false && meta[id].dataType == 1, "given id incorrect");
        //发出计算共享事件
        emit computingShareEvent(id, meta[id].dataAbstract, model, dataset, round, epoch);
        return "emit computingshare event successfully";
    }

    /**
     * Author Mengeshall
     * Date 2021/10/13
     * Description 计算共享回调合约，负责将结果上链
     */
     function computingsharecallback(uint id, string paramAddr, string paramAbstract, string correctRate) public returns(string){
        //验证是否有该id对应的metadata
        require(meta[id].isValid != false && meta[id].dataType == 1, "given id incorrect");
        //计算共享结果上链
        result[id] = computingResult(paramAddr, paramAbstract, correctRate, "");
        return result[id].correctRate;
    }

    /**
     * Author Mengeshall
     * Date 2021/10/14
     * Description 数据共享合约，负责发出数据共享事件
     */
    function datashare(uint id) public returns(string){
        //验证是否有该id对应的metadata
        require(meta[id].isValid != false && meta[id].dataType == 2, "given id incorrect");
        //发出数据共享事件
        emit dataShareEvent(id,meta[id].source, meta[id].dataAbstract);
        return "emit datashare event successfully";
    }

    /**
     * Author Mengeshall
     * Date 2021/10/14
     * Description 数据共享回调合约，负责将结果上链
     */
    function datasharecallback(uint id, string _result) public returns(string){
        //验证是否有该id对应的metadata
        require(meta[id].isValid != false && meta[id].dataType == 2, "given id incorrect");
        //数据共享结果上链
        result[id] = computingResult("", "", "", _result);
        return result[id].data;
    }
}