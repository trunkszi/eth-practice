package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	fmt.Println("rpc http demo")
	web3ClientVersion()
	ethAccounts()
}

func web3ClientVersion() {
	msg := `{
    "jsonrpc": "2.0",
    "method": "web3_clientVersion",
    "params": [],
    "id": 1
  }`
	rsp,err := http.Post("http://localhost:7545","application/json",bytes.NewBuffer([]byte(msg)))
	assert(err)
	ret,err := ioutil.ReadAll(rsp.Body)
	assert(err)
	fmt.Println("web3_clientVersion: ", string(ret))
}

func ethAccounts() {
	msg := `{
    "jsonrpc": "2.0",
    "method": "eth_accounts",
    "params": [],
    "id": 1
  }`
	rsp,err := http.Post("http://localhost:7545","application/json",bytes.NewBuffer([]byte(msg)))
	assert(err)
	ret,err := ioutil.ReadAll(rsp.Body)
	assert(err)
	fmt.Println("eth_accounts: ", string(ret))
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}