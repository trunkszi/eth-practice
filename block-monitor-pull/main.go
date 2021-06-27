package main

import (
	"context"
	"ethtool"
	"fmt"
	"math/big"
	"time"
)

func main(){
	fmt.Println("block monitor demo")
	go trigger()
	pull_monitor()
}

func trigger(){
	client,err := ethtool.Dial("http://localhost:7545")
	assert(err)

	ctx := context.Background()

	accounts,err := client.EthAccounts(ctx)
	assert(err)

	timer := time.Tick(5 * time.Second)
	for range timer {
		msg := map[string]interface{}{
			"from": accounts[0],
			"to": accounts[1],
			"value": big.NewInt(1000),
		}
		txid,err := client.EthSendTransaction(context.Background(),msg)
		assert(err)
		fmt.Println("trigger txid: ",txid.Hex())
	}
}

func pull_monitor(){
	client,err := ethtool.Dial("http://localhost:7545")
	assert(err)

	ctx := context.Background()

	fid,err := client.EthNewBlockFilter(ctx)
	assert(err)
	fmt.Println("filter id: ",fid)

	ticker := time.Tick(2 * time.Second)
	for range ticker {
		hashes,err := client.EthGetFilterChanges(context.Background(),fid)
		assert(err)
		for _, hash := range hashes {
			fmt.Println("captured block hash: ", hash.Hex())
		}
	}

}

func assert(err error){
	if err != nil {
		panic(err)
	}
}