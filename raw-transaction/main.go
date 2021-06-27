package main

import (
	"bytes"
	"context"
	"ethtool"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func main(){
	credential,err := ethtool.HexToCredential("0xdd655526371984f6ae5715ba9ad0ad0a1d190da49d5cc5ad53840257a7b45ff9")
	assert(err)
	fmt.Println("from address: ",credential.Address.Hex())

	to := common.HexToAddress("0x36FeC36DB2f4266457277a6FB3864AC1C83a80bc")
	fmt.Println("to address: ", to.Hex())

	client,err := ethtool.Dial("http://localhost:7545")
	assert(err)

	ctx := context.Background()

	chainid,err := client.NetVersion(ctx)
	assert(err)
	fmt.Println("chainid: ",chainid)

	nonce,err := client.EthGetTransactionCount(ctx,credential.Address,"pending")
	assert(err)
	fmt.Println("nonce: ",nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &to,
		Data: nil,
	})

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &to,
		Value:    big.NewInt(1000000000000000000),
		Data:     nil,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP2930Signer(chainid), credential.PrivateKey)
	if err != nil {
		panic(err)
	}
	assert(err)
	fmt.Println("signedTx: ",signedTx)

	/*
	  err = client.SendTransaction(ctx,signedTx)
	  assert(err)
	  fmt.Println("raw tx id: ",signedTx.Hash().Hex())
	*/

	buf := new(bytes.Buffer)

	err = signedTx.EncodeRLP(buf)
	assert(err)

	txid,err := client.EthSendRawTransaction(ctx,buf.Bytes())
	assert(err)
	fmt.Println("raw tx id: ",txid.Hex())

	balance,err := client.EthGetBalance(ctx,to,"latest")
	assert(err)
	fmt.Println("balance received: ",balance)
}


func assert(err error) {
	if err != nil {
		panic(err)
	}
}

