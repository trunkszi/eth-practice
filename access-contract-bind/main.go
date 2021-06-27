package main

import (
	"context"
	"contract/warpper/eztoken"
	"ethtool"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"io/ioutil"
	"math/big"
)

// 需要先部署合约

func main() {
	fmt.Println("access contract demo")

	rpcClient, err := rpc.Dial("http://localhost:7545")
	assert(err)
	client := ethclient.NewClient(rpcClient)

	assert(err)

	addrHexBytes, err := ioutil.ReadFile("../contract/build/EzToken.addr")
	assert(err)
	contractAddress := common.HexToAddress(string(addrHexBytes))
	fmt.Println("contract address: ", contractAddress.Hex())

	inst, err := eztoken.NewEztoken(contractAddress, client)
	assert(err)
	//fmt.Println("inst: ",inst)

	credential, err := ethtool.HexToCredential("0x57dc2c69b9de918b4c42f2c64abe3e1228602f03816d8e728571a4d341ecaad0")
	assert(err)

	chainId, err := client.NetworkID(context.Background())
	assert(err)

	fmt.Println("chainId: ", chainId)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	assert(err)

	nonce, err := client.PendingNonceAt(context.Background(), credential.Address)
	assert(err)
	value := big.NewInt(0) // in wei (0 eth)
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: credential.Address,
		To:   nil,
		Data: nil,
		Value: value,
		GasPrice: gasPrice,
	})

	txOpts,err := bind.NewKeyedTransactorWithChainID(credential.PrivateKey, chainId)
	txOpts.Context = context.Background()
	txOpts.GasPrice = gasPrice
	txOpts.Value = value
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.GasLimit = gasLimit * 10
	assert(err)

	toAddress := common.HexToAddress("0xf74dFb6664815dF0D69B3784Aa2369eb097EA1b7")
	amount := big.NewInt(100)

	tx, err := inst.Transfer(txOpts, toAddress, amount)
	assert(err)
	fmt.Println("txid: ", tx.Hash().Hex())

	callOpts := &bind.CallOpts{
		From: credential.Address,
	}
	balance, err := inst.BalanceOf(callOpts, toAddress)
	assert(err)
	fmt.Println("balance: ", balance)

}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}
