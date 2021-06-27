package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main(){
	fmt.Println("import key demo")

	prvKeyHex := "0x787230eb7e49b0b457820a8e1f90882d573cc4eea922ab0c255f20cb783ea836"

	prvKey,err := crypto.HexToECDSA(prvKeyHex[2:])
	assert(err == nil, err)
	fmt.Println("private key: ",prvKeyHex)

	pubKey := prvKey.Public()
	pubKeyECDSA,ok := pubKey.(*ecdsa.PublicKey)
	assert(ok, "type cast failed")
	pubKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)
	fmt.Println("public key: ",hexutil.Encode(pubKeyBytes))

	address := crypto.PubkeyToAddress(*pubKeyECDSA)
	fmt.Println("addrss: ",address.Hex())
}

func assert(expr bool, msg interface{}){
	if !expr {
		panic(msg)
	}
}