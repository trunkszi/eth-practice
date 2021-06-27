package main

import (
	"ethtool"
	"fmt"
	"math/big"
)

func main(){
	weiValue := big.NewInt(2200000000)
	fmt.Println("value in wei: ", weiValue)
	fmt.Println("-> in gwei: ",ethtool.FromWei(weiValue,ethtool.Gwei))
	fmt.Println("-> in ether: ",ethtool.FromWei(weiValue,ethtool.Ether))

	etherValue := big.NewFloat(23.456)
	fmt.Println("value in ether: ",etherValue)
	fmt.Println("-> in wei: ",ethtool.ToWei(etherValue,ethtool.Ether))

	gweiValue := big.NewFloat(2)
	fmt.Println("value in gwei: ",gweiValue)
	fmt.Println("-> in wei: ",ethtool.ToWei(gweiValue,ethtool.Gwei))
}

