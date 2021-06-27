package main

import (
	"context"
	"ethtool"
	"fmt"
	"math/big"
	"time"
)

func main(){
	fmt.Println("pending tx monitor demo")
	go trigger()
	filter_monitor()
}

func trigger(){
	client,err := ethtool.Dial("http://localhost:7545")
	assert(err)

	ctx := context.Background()

	accounts,err := client.EthAccounts(ctx)
	assert(err)

	ticker := time.Tick(5 * time.Second)
	for range ticker {
		msg := map[string]interface{}{
			"from": accounts[0],
			"to": accounts[1],
			"value": big.NewInt(1000),
		}
		txid,err := client.EthSendTransaction(ctx,msg)
		assert(err)
		fmt.Println("trigger txid: ",txid.Hex())
	}
}

func filter_monitor(){
	client,err := ethtool.Dial("ws://localhost:7545")
	assert(err)

	ctx := context.Background()

	fid,err := client.EthNewPendingTransactionFilter(ctx)
	assert(err)

	timer := time.Tick(2 * time.Second)
	for range timer {
		hashes,err := client.EthGetFilterChanges(ctx,fid)
		assert(err)
		for _,hash := range hashes{
			fmt.Println("monitored pending txid: ", hash.Hex())
		}
	}

}

func assert(err error){
	if err != nil {
		panic(err)
	}
}