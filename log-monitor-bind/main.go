package main

import (
	"context"
	"ethtool"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"math/big"
	"practices/contract/warpper/eztoken"
	"time"
)

// Go封装包监听合约日志, 需要先部署合约
func main() {
	fmt.Println("monitor contract log with wrapper")
	go trigger()
	monitor()
}

func trigger() {
	addrHexBytes, err := ioutil.ReadFile("../contract/build/EzToken.addr")
	assert(err)
	contractAddress := common.HexToAddress(string(addrHexBytes))
	assert(err)

	client, err := ethtool.Dial("ws://localhost:7545")
	assert(err)

	inst, err := eztoken.NewEztoken(contractAddress, client)
	assert(err)
	//fmt.Println(inst)

	credential, err := ethtool.HexToCredential("0x57dc2c69b9de918b4c42f2c64abe3e1228602f03816d8e728571a4d341ecaad0")
	assert(err)

	value := big.NewInt(0) // in wei (0 eth)


	chainId,err := client.NetworkID(context.Background())
	assert(err)

	fmt.Println("chainId: ", chainId)


	toAddress := common.HexToAddress("0xf74dFb6664815dF0D69B3784Aa2369eb097EA1b7")
	amount := big.NewInt(100)

	timer := time.Tick(15 * time.Second)
	for range timer {
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


		tx, err := inst.Transfer(txOpts, toAddress, amount)
		assert(err)
		fmt.Println("trigger txid: ", tx.Hash().Hex())
	}
}

func monitor() {
	addrHexBytes, err := ioutil.ReadFile("../contract/build/EzToken.addr")
	assert(err)
	contractAddress := common.HexToAddress(string(addrHexBytes))
	assert(err)

	client, err := ethtool.Dial("ws://localhost:7545")
	assert(err)

	inst, err := eztoken.NewEztoken(contractAddress, client)
	assert(err)

	watchOpts := &bind.WatchOpts{}
	events := make(chan *eztoken.EztokenTransfer)
	var _from []common.Address
	var _to []common.Address
	sub, err := inst.WatchTransfer(watchOpts, events, _from, _to)

	assert(err)
	//fmt.Println(sub)

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case event := <-events:
			fmt.Println("captured:")
			fmt.Println("-> from: ", event.From.Hex())
			fmt.Println("-> to: ", event.To.Hex())
			fmt.Println("-> value:", event.Value)
		}
	}

}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}
