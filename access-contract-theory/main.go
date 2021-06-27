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

func main() {
	fmt.Println("access contract theory demo")

	abiBytes, err := ioutil.ReadFile("../contract/build/EzToken.abi")
	assert(err)
	tokenAbi, err := abi.JSON(strings.NewReader(string(abiBytes)))
	assert(err)

	addrBytes, err := ioutil.ReadFile("../contract/build/EzToken.addr")
	assert(err)
	contractAddress := common.HexToAddress(string(addrBytes))
	fmt.Println("contract address: ", contractAddress.Hex())

	client, err := ethtool.Dial("http://localhost:7545")
	assert(err)

	ctx := context.Background()

	accounts, err := client.EthAccounts(ctx)
	assert(err)

	//transfer from account 0 => 1

	data, err := tokenAbi.Pack(
		"transfer",
		accounts[1],
		big.NewInt(100),
	)
	assert(err)
	msg := map[string]interface{}{
		"from": accounts[0],
		"to":   contractAddress,
		"gas":  big.NewInt(2000000),
		"data": common.Bytes2Hex(data),
	}
	txid, err := client.EthSendTransaction(ctx, msg)
	assert(err)
	fmt.Println("txid: ", txid.Hex())

	//balanceOf
	data, err = tokenAbi.Pack(
		"balanceOf",
		accounts[0],
	)
	msg = map[string]interface{}{
		"from": accounts[0],
		"to":   contractAddress,
		"data": common.Bytes2Hex(data),
	}
	ret, err := client.EthCall(ctx, msg)
	assert(err)
	fmt.Println("balance: ", ret)

	//abi decode balance
	//balancePtr := new(*big.Int)
	balance, err := tokenAbi.Unpack("balanceOf", ret)
	assert(err)

	fmt.Println("balance decoded: ", balance)
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}
