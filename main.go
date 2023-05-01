package main

import (
	"fmt"
	"main/function"
	"main/str"
)

func main() {
	wallet2 := function.CreateWallet("iDissapoint")
	str.User = append(str.User, *wallet2)
	fmt.Println("", wallet2.PrivateKey)
	Data2 := &str.Wallet{WalletName: wallet2.WalletName, PrivateKey: wallet2.PrivateKey, Publickey: wallet2.Publickey, Balance: wallet2.Balance, Address: wallet2.Address, Password: wallet2.Password}
	Data2.GetPirvateKey("hi")
}
