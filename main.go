package main

import (
	"fmt"
	"main/function"
	"main/str"
)

func main() {
	//get := function.GetValidatorID(0)
	//fmt.Println("", get)
	//function.CreateValidator("0xc55c16b98129b480bbbc609c5b130a", 1000000, "fd596985a7a014bb929f3ed27132d3ec1c5d898ce00da034815cf8ccc45a8cd7", 1, "2023 05 06 12:00", 10, 11, 99)
	send1 := function.CreateWallet("iDissapoint").Address
	To1toSen2 := function.CreateWallet("iDissapoint2").Address
	to2 := function.CreateWallet("iDissapoint3").Address
	gen := function.CreateGenblock()
	tx := &str.TransactionDetails{
		Sender:   send1,
		Receiver: To1toSen2,
		Amount:   10,
		Fees:     0.01,
		Status:   true,
	}
	new := function.CreateNewBlock(gen, tx)
	fmt.Println("\n", new.BlockHeight)
	fmt.Println("\n", new.PreviousHash)
	fmt.Println("\n", new.CurrentHash)
	fmt.Println("\n", new.TimeStamp)
	fmt.Println("\n", new.Transaction.Sender)
	fmt.Println("\n", new.Transaction.Receiver)
	fmt.Println("\n", new.Transaction.Amount)
	fmt.Println("\n", new.Transaction.Fees)
	fmt.Println("", function.CalculateBlockReward(int64(gen.BlockHeight)))
	tx2 := &str.TransactionDetails{
		Sender:   To1toSen2,
		Receiver: to2,
		Amount:   10,
		Fees:     0.01,
		Status:   true,
	}
	new2 := function.CreateNewBlock(new, tx2)
	fmt.Println("\n", new2.BlockHeight)
	fmt.Println("\n", new2.PreviousHash)
	fmt.Println("\n", new2.CurrentHash)
	fmt.Println("\n", new2.TimeStamp)
	fmt.Println("\n", new2.Transaction.Sender)
	fmt.Println("\n", new2.Transaction.Receiver)
	fmt.Println("\n", new2.Transaction.Amount)
	fmt.Println("\n", new2.Transaction.Fees)
	fmt.Println("", function.CalculateBlockReward(int64(new.BlockHeight)))
	tx3 := &str.TransactionDetails{
		Sender:   To1toSen2,
		Receiver: to2,
		Amount:   10,
		Fees:     0.01,
		Status:   true,
	}
	new3 := function.CreateNewBlock(new2, tx3)
	fmt.Println("\n", new3.BlockHeight)
	fmt.Println("\n", new3.PreviousHash)
	fmt.Println("\n", new3.CurrentHash)
	fmt.Println("\n", new3.TimeStamp)
	fmt.Println("\n", new3.Transaction.Sender)
	fmt.Println("\n", new3.Transaction.Receiver)
	fmt.Println("\n", new3.Transaction.Amount)
	fmt.Println("\n", new3.Transaction.Fees)
	fmt.Println("", function.CalculateBlockReward(int64(new2.BlockHeight)))
	//function.CreateWallet("iDissapoint")
	//function.CreateWallet("iDissapoint2")
	//function.CreateWallet("iDissapoint3")
	//fmt.Println("", function.GetValidator())
	//str.User = append(str.User, wallet2, wallet, wallet3)
	/*for i := range str.User {
		fmt.Println("\n\n", str.User[i])
	}*/

}
