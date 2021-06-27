package main

import (
	"context"
	"ethtool"
	"fmt"
)

func main(){
	fmt.Println("transaction demo")

	client,err := ethtool.Dial("http://localhost:7545")
	assert(err)

	ctx := context.Background()

	accounts,err := client.EthAccounts(ctx)
	assert(err)

	fmt.Printf("transfer:  %v -> %v\n",accounts[0].Hex(),accounts[1].Hex())

	var msg = map[string]interface{}{
		"from": accounts[0],
		"to": accounts[1],
		"value": "1000000000000000000",
	}

	txid,err := client.EthSendTransaction(ctx,msg)
	assert(err)
	fmt.Println("txid: ",txid.Hex())

	balance,err := client.EthGetBalance(ctx,accounts[0],"latest")
	assert(err)
	fmt.Println("account#0 balance: ", balance)
}

func assert(err error){
	if err != nil {
		panic(err)
	}
}