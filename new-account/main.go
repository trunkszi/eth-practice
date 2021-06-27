package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main(){
	fmt.Println("new account demo")

	prvKey,err := crypto.GenerateKey()
	assert(err == nil,err)
	prvKeyBytes := crypto.FromECDSA(prvKey)
	fmt.Println("private key: ",hexutil.Encode(prvKeyBytes))

	pubKey := prvKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	assert(ok,"cast failed")
	pubKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)
	fmt.Println("public key: ",hexutil.Encode(pubKeyBytes))

	address :=crypto.PubkeyToAddress(*pubKeyECDSA)
	fmt.Println("address: ",address.Hex())
}

func assert(expected bool, msg interface{}) {
	if !expected {
		panic(msg)
	}
}