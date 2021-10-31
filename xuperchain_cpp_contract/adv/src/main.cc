/*
 * @Author: your name
 * @Date: 2021-10-29 14:52:17
 * @LastEditTime: 2021-10-29 15:47:30
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \undefinedc:\Users\zhanglijia\Desktop\contract.cpp
 */
#include <iostream>
#include <map>
#include <string>
#include "xchain/xchain.h"
#include "xchain/json/json.h"
using namespace std;

struct metaData{
    std::string uploader;
    std::string name;
    std::string type;
    std::string ip;
    std::string route;
    std::string abstract;
};


struct faderatedAIDemand
{
    std::string model;
    std::string dataset;
    std::string round;
    std::string epoch;
};
 
struct AdvContract : public xchain::Contract {};
DEFINE_METHOD(AdvContract, initialize) {
    xchain::Context* ctx = self.context();
    ctx->put_object("AdvFileAssetCount", "0");
}

DEFINE_METHOD(AdvContract, createAdv) {
    xchain::Context* ctx = self.context();
//    std::string u_input = ctx->arg("uploader");
//    std::string n_input = ctx->arg("name");
//    std::string t_input = ctx->arg("type");
//    std::string i_input = ctx->arg("ip");
//    std::string r_input = ctx->arg("route");
//    std::string a_input = ctx->arg("abstract");
    std::string v = ctx->arg("value");
    auto j = xchain::json::parse(v);
//    metaData meta;
//    md.uploader = u_input;
//    md.name = n_input;
//    md.type = t_input;
//    md.ip = i_input;
//    md.route = r_input;
//    md.abstract = a_input;
    
    std::string advCount;
    ctx->get_object("AdvFileAssetCount", &advCount);
    std::string newId;
    newId = "AdvFileAssetId_"+std::to_string(std::stoi(advCount)+1);
    ctx->put_object("AdvFileAssetCount", std::to_string(std::stoi(advCount)+1));
    
    ctx->put_object(newId, j.dump());
    ctx->ok(newId);
}

DEFINE_METHOD(AdvContract, query) {
    xchain::Context* ctx = self.context();
    std::string meta_id = ctx->arg("id");
    
    std::string metaJson;
    ctx->get_object(meta_id, &metaJson);

    auto meta = xchain::json::parse(metaJson);
    
//    metaData meta;
//    ctx->get_object(meta_id, &meta);

    xchain::json j = {
                {"id", meta_id},
                {"uploader", meta["uploader"]},
                {"name", meta["name"]},
                {"type", meta["type"]},
                {"ip", meta["ip"]},
                {"route", meta["route"]},
                {"abstract", meta["abstract"]},
    };

    if(ctx->emit_event("queryEvent", j.dump())){
        ctx->ok("emit event success!");
    } else {
        ctx->error("emit event failed");
    }
}

DEFINE_METHOD(AdvContract, queryCallback){
    xchain::Context* ctx = self.context();
    std::string meta_id = ctx->arg("id");
    std::string result_id = "result_" + meta_id;
    std::string data = ctx->arg("data");
    
    std::string result_data;
    if(ctx->get_object(result_id, &result_data)){
        if(ctx->delete_object(result_id)){
            ctx->put_object(result_id, data);
        }
    } else {
        ctx->put_object(result_id, data);
    }
    ctx->ok(result_id);
}

DEFINE_METHOD(AdvContract, computingshare) {
    xchain::Context* ctx = self.context();
    std::string meta_id = ctx->arg("id");
    std::string model = ctx->arg("model");
    std::string dataset = ctx->arg("dataset");
    std::string round = ctx->arg("round");
    std::string epoch = ctx->arg("epoch");
    
//    metaData meta;
//    ctx->get_object(meta_id, &meta);
    
    std::string metaJson;
    ctx->get_object(meta_id, &metaJson);
//    auto meta = xchain::json::parse(metaJson);
    
    xchain::json fadJson = {
                {"model", model},
                {"dataset", dataset},
                {"round", round},
                {"epoch", epoch},
    };
    
//    faderatedAIDemand fad;
//    fad.model = model;
//    fad.dataset = dataset;
//    fad.round = round;
//    fad.epoch = epoch;
    
    xchain::json j = {
        {"id", meta_id},
        {"metaData", metaJson},
        {"faderatedAIDemand", fadJson},
    };
    
    if(ctx->emit_event("computingEvent", j.dump())){
        ctx->ok("emit event success!");
    } else {
        ctx->error("emit event failed!");
    }
}

DEFINE_METHOD(AdvContract, computingCallback) {
    xchain::Context* ctx = self.context();
    std::string meta_id = ctx->arg("id");
    std::string result_id = "result_" + meta_id;
    std::string data = ctx->arg("data");
    
    std::string result_data;
    if(ctx->get_object(result_id, &result_data)){
        if(ctx->delete_object(result_id)){
            ctx->put_object(result_id, data);
        }
    } else {
        ctx->put_object(result_id, data);
    }
    ctx->ok(result_id);
}

DEFINE_METHOD(AdvContract, addUser) {
    xchain::Context* ctx = self.context();
    std::string name = ctx->arg("name");
    std::string abstract = ctx->arg("abstract");
    
    std::string tmp;
    if(ctx->get_object(abstract, &tmp)){
        ctx->error("abstract has already uploaded!");
    } else {
        ctx->put_object(abstract, name);
    }
    ctx->ok("add user success!");
}

DEFINE_METHOD(AdvContract, verifyUser) {
    xchain::Context* ctx = self.context();
    std::string name = ctx->arg("name");
    std::string abstract = ctx->arg("abstract");
    
    std::string tmp;
    if(ctx->get_object(abstract, &tmp)){
        if(tmp.compare(name)==0){
            ctx->ok("1");
        } else {
            ctx->ok("0");
        }
    }
}
