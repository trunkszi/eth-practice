package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RpcRequest struct{
	JsonRpc string `json:"jsonrpc"`
	Method string `json:"method"`
	Params []interface{} `json:"params"`
	Id int64 `json:"id"`
}

type RpcResponse struct{
	JsonRpc string `json:"jsonrpc"`
	Id int64 `json:"id"`
	Result interface{} `json:"result"`
	Error interface{}  `json:"error"`
}

func main(){
	fmt.Println("rpc http codec demo")
	web3ClientVersion()
	ethAccounts()
}

func web3ClientVersion(){
	rpcReq := RpcRequest{
		JsonRpc: "2.0",
		Method: "web3_clientVersion",
		Params: []interface{}{},
		Id: time.Now().Unix(),
	}
	payload,err := json.Marshal(rpcReq)
	assert(err)
	fmt.Println(string(payload))
	rsp,err := http.Post("http://localhost:7545","application/json",bytes.NewBuffer([]byte(payload)))
	assert(err)
	ret,err := ioutil.ReadAll(rsp.Body)
	assert(err)
	var rpcRsp RpcResponse
	err = json.Unmarshal(ret,&rpcRsp)
	assert(err)
	fmt.Println(rpcRsp)
}

func ethAccounts(){
	rpcReq := RpcRequest{
		JsonRpc: "2.0",
		Method: "eth_accounts",
		Params: []interface{}{},
		Id: time.Now().Unix(),
	}
	payload,err := json.Marshal(rpcReq)
	assert(err)
	fmt.Println(string(payload))
	rsp,err := http.Post("http://localhost:7545","application/json",bytes.NewBuffer([]byte(payload)))
	assert(err)
	ret,err := ioutil.ReadAll(rsp.Body)
	assert(err)
	var rpcRsp RpcResponse
	err = json.Unmarshal(ret,&rpcRsp)
	assert(err)
	fmt.Println(rpcRsp)
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}