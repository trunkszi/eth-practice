package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main() {
	fmt.Println("ethclient demo")
	blockByNumber()
	gasPrice()
}

func blockByNumber() {
	client, err := ethclient.Dial("http://localhost:7545")
	assert(err)
	blockNumber := big.NewInt(0)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	assert(err)
	fmt.Println("hash: ", block.Hash().Hex())
	fmt.Println("coinbase: ", block.Coinbase().Hex())
	fmt.Println("num of transactions: ", block.Transactions().Len())
}

func gasPrice() {
	client, err := ethclient.Dial("http://localhost:7545")
	assert(err)
	price, err := client.SuggestGasPrice(context.Background())
	assert(err)
	fmt.Println("gas price: ", price)
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}
