package main

import (
	"ethtool"
	"fmt"
)

func main(){
	credential,err := ethtool.NewCredential()
	assert(err)
	fmt.Println("private key: ",credential.PrivateKeyHex())
	fmt.Println("public key: ",credential.PublicKeyHex())
	fmt.Println("address: ",credential.AddressHex())
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}
