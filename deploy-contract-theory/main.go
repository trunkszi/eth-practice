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
)

func main(){
	fmt.Println("deploy contract theory demo")
	abiBytes,err := ioutil.ReadFile("../contract/build/EzToken.abi")
	assert(err)
	//fmt.Println(abiBytes)

	binHexBytes,err := ioutil.ReadFile("../contract/build/EzToken.bin")
	assert(err)
	binBytes := common.Hex2Bytes(string(binHexBytes))
	//fmt.Println(binBytes)

	tokenAbi,err := abi.JSON(strings.NewReader(string(abiBytes)))
	assert(err)
	encodedParams,err := tokenAbi.Pack(
		"",
		big.NewInt(1000000),
		"HAPPY COIN",
		uint8(0),
		"HAPY",
	)
	assert(err)
	//fmt.Println(encodedParams)
	data := append(binBytes,encodedParams...)

	client,err := ethtool.Dial("http://localhost:7545")
	assert(err)

	ctx := context.Background()

	accounts, err := client.EthAccounts(ctx)
	assert(err)

	msg := map[string]interface{}{
		"from": accounts[0],
		"data": common.Bytes2Hex(data),
		"gas": big.NewInt(2000000),
	}
	txid,err := client.EthSendTransaction(ctx,msg)
	assert(err)
	fmt.Println("txid: ",txid.Hex())

	//wait

	receipt,err := client.EthGetTransactionReceipt(ctx,txid)
	assert(err)
	fmt.Println("contract address: ", receipt.ContractAddress.Hex())
	err =ioutil.WriteFile("../contract/build/EzToken.addr",[]byte(receipt.ContractAddress.Hex()),0644)
	assert(err)
	fmt.Println("done.")

}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}