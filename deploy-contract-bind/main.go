package main

import (
	"context"
	"ethtool"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"io/ioutil"
	"math/big"
	"practices/contract/warpper/eztoken"
)

func main() {
	fmt.Println("deploy contract demo")

	rpcClient, err := rpc.Dial("http://localhost:7545")
	assert(err)

	client := ethclient.NewClient(rpcClient)

	credential, err := ethtool.HexToCredential("0x57dc2c69b9de918b4c42f2c64abe3e1228602f03816d8e728571a4d341ecaad0")
	assert(err)

	chainId, err := client.ChainID(context.Background())
	assert(err)

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	assert(err)

	nonce, err := client.PendingNonceAt(context.Background(), credential.Address)
	assert(err)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: credential.Address,
		To:   nil,
		Data: nil,
		Value: value,
		GasPrice: gasPrice,
	})

	fmt.Println("chainId: ", chainId)
	txOpts, err := bind.NewKeyedTransactorWithChainID(credential.PrivateKey, chainId)
	txOpts.Context = context.Background()
	txOpts.GasPrice = gasPrice
	txOpts.Value = value
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.GasLimit = gasLimit * 10

	assert(err)

	tokenSupply := big.NewInt(1000000)
	tokenName := "HAPPY TOKEN"
	tokenDecimals := uint8(0)
	tokenSymbol := "HAPY"

	address, tx, inst, err := eztoken.DeployEztoken(txOpts, client, tokenSupply, tokenName, tokenDecimals, tokenSymbol)
	assert(err)
	fmt.Println("deployed at: ", address.Hex())
	fmt.Println("txid: ", tx.Hash().Hex())
	_ = inst

	fmt.Println("save deployed address... : ", address.Hex())
	err = ioutil.WriteFile("../contract/build/EzToken.addr", []byte(address.Hex()), 0644)
	assert(err)

	fmt.Println("done.")



}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}
