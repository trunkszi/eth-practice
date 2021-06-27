package main

import (
	"context"
	"ethtool"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

func main(){
	fmt.Println("pending tx monitor demo")
	go trigger()
	push_monitor()
}

func trigger(){
	client,err := ethtool.Dial("http://localhost:7545")
	assert(err)

	accounts,err := client.EthAccounts(context.Background())
	assert(err)

	timer := time.Tick(5 * time.Second)
	for range timer {
		msg := map[string]interface{}{
			"from": accounts[0],
			"to": accounts[1],
			"value": big.NewInt(10000),
		}
		txid,err := client.EthSendTransaction(context.Background(),msg)
		assert(err)
		fmt.Println("trigger txid: ",txid.Hex())
	}
}

func push_monitor(){
	client,err := ethtool.Dial("ws://localhost:7545")
	assert(err)

	hashes := make(chan *common.Hash)
	sub,err := client.SubscribeNewPendingTransactions(context.Background(),hashes)
	assert(err)

	for {
		select {
		case err := <- sub.Err():
			panic(err)
		case hash := <- hashes:
			fmt.Println("captured tx id: ",hash.Hex())
		}
	}
}

func assert(err error){
	if err != nil {
		panic(err)
	}
}