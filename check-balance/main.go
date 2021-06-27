package main

import (
	"context"
	"ethtool"
	"fmt"
)

func main(){
	fmt.Println("check balance demo")

	client,err := ethtool.Dial("http://localhost:7545")
	assert(err)

	ctx := context.Background()

	accounts,err := client.EthAccounts(ctx)
	assert(err)

	balance,err := client.EthGetBalance(ctx,accounts[0],"latest")
	assert(err)
	fmt.Println("balance@latest: ", balance)

	balance,err = client.EthGetBalance(ctx,accounts[0],"earliest")
	assert(err)
	fmt.Println("balance@earliest: ", balance)
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}