package main

import (
	_ "encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

func main(){
	keyHex := "0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d"
	auth := "my secret"
	fmt.Println("original: ",keyHex)

	encoded := encode(keyHex,auth)
	fmt.Println("encoded: ", string(encoded))

	decoded := decode(encoded,auth)
	fmt.Println("decoded: ",decoded)
}

func encode(keyHex,auth string) []byte {
	prvKey,err := crypto.HexToECDSA(keyHex[2:])
	assert(err)
	uuid,err := uuid.NewRandom()
	if err != nil{
		panic(err)
	}
	key := &keystore.Key{
		Id: uuid,
		Address: crypto.PubkeyToAddress(prvKey.PublicKey),
		PrivateKey: prvKey,
	}
	json,err := keystore.EncryptKey(key,auth,keystore.StandardScryptN,keystore.StandardScryptP)
	assert(err)
	return json
}

func decode(json []byte,auth string) string {
	key,err := keystore.DecryptKey(json,auth)
	assert(err)
	return hexutil.Encode(crypto.FromECDSA(key.PrivateKey))
}


func assert(err error) {
	if err != nil {
		panic(err)
	}
}