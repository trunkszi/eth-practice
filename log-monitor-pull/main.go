package main

import (
	"context"
	"ethtool"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"math/big"
	"strings"
	"time"
)

func main(){
	fmt.Println("log monitor demo")
	go trigger()
	filter_monitor()
}

func trigger(){
	abiBytes,err := ioutil.ReadFile("../contract/build/EzToken.abi")
	assert(err)
	tokenAbi,err := abi.JSON(strings.NewReader(string(abiBytes)))
	assert(err)

	addrBytes,err := ioutil.ReadFile("../contract/build/EzToken.addr")
	assert(err)
	contractAddress := common.HexToAddress(string(addrBytes))

	client,err := ethtool.Dial("http://localhost:7545")
	assert(err)

	ctx := context.Background()

	accounts,err := client.EthAccounts(ctx)
	assert(err)

	data,err := tokenAbi.Pack("transfer",accounts[1],big.NewInt(100))
	assert(err)
	msg := map[string]interface{}{
		"from": accounts[0],
		"to": contractAddress,
		"data": common.Bytes2Hex(data),
		"gas": big.NewInt(4000000),
	}

	ticker := time.Tick(5 * time.Second)
	for range ticker {
		txid,err := client.EthSendTransaction(ctx,msg)
		assert(err)
		fmt.Println("trigger txid: ",txid.Hex())
	}
}

func filter_monitor(){
	client,err := ethtool.Dial("http://localhost:7545")
	assert(err)

	ctx := context.Background()

	opts := map[string]string{}
	fid,err := client.EthNewFilter(ctx,opts)
	assert(err)
	fmt.Println("filter id: ",fid)

	ticker := time.Tick(2 * time.Second)
	for range ticker {
		logs,err := client.EthGetLogFilterChanges(ctx,fid)
		assert(err)
		for _,log := range logs {
			fmt.Printf("captured log:%+v\n", log)
		}
	}
}

func assert(err error){
	if err != nil {
		panic(err)
	}
}

