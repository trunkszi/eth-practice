package main

import (
	"context"
	"ethtool"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"io/ioutil"
	"math/big"
	"practices/contract/warpper/eztoken"
	"time"
)

func main(){
	fmt.Println("log monitor demo")
	go trigger()
	push_monitor()
}

func trigger(){
	client,err := ethtool.Dial("http://localhost:7545")
	assert(err)

	data,err := ioutil.ReadFile("../contract/build/EzToken.addr")
	assert(err)
	contractAddress := common.HexToAddress(string(data))
	inst,err := eztoken.NewEztoken(contractAddress,client)
	assert(err)

	credential,err := ethtool.HexToCredential("0x57dc2c69b9de918b4c42f2c64abe3e1228602f03816d8e728571a4d341ecaad0")
	assert(err)

	toAddress := common.HexToAddress("0xf74dFb6664815dF0D69B3784Aa2369eb097EA1b7")
	amount := big.NewInt(100)

	chainId,err := client.NetworkID(context.Background())
	assert(err)

	timer := time.Tick(5 * time.Second)
	for range timer {
		value := big.NewInt(0)
		gasPrice, err := client.SuggestGasPrice(context.Background())
		assert(err)

		nonce, err := client.PendingNonceAt(context.Background(), credential.Address)
		assert(err)

		gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			From:     credential.Address,
			To:       nil,
			Data:     nil,
			Value:    value,
			GasPrice: gasPrice,
		})
		txOpts, err := bind.NewKeyedTransactorWithChainID(credential.PrivateKey, chainId)
		txOpts.Context = context.Background()
		txOpts.GasPrice = gasPrice
		txOpts.Value = value
		txOpts.Nonce = big.NewInt(int64(nonce))
		txOpts.GasLimit = gasLimit * 10

		tx, err := inst.Transfer(txOpts,toAddress,amount)
		assert(err)
		fmt.Println("trigger txid: ",tx.Hash().Hex())
	}
}

func push_monitor(){
	client,err := ethtool.Dial("ws://localhost:7545")
	assert(err)

	data,err := ioutil.ReadFile("/Users/quincy/Desktop/go-project/src/practices/contract/build/EzToken.addr")
	assert(err)
	contractAddress := common.HexToAddress(string(data))
	_ = contractAddress

	query := ethereum.FilterQuery{
		//    Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub,err := client.SubscribeFilterLogs(context.Background(),query,logs)
	assert(err)

	for {
		select {
		case err := <- sub.Err():
			panic(err)
		case log := <- logs:
			fmt.Println("captured log:")
			fmt.Println("-> address: ",log.Address.Hex())
			fmt.Println("-> data: ",log.Data)
			fmt.Println("-> topics: ",log.Topics)
		}
	}
}

func assert(err error){
	if err != nil {
		panic(err)
	}
}

