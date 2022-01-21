package main

import "fmt"

func main() {
	NewTransaction(Transaction{
		SenderPublicKey:   "",
		ReceiverPublicKey: "",
		Amount:            0,
		TransactionType:   0,
		Id:                "",
		Timestamp:         0,
		Signature:         "",
	})
	fmt.Println(transaction)
	Hash()

}

//1839806922502695369

//7314111797897617339
//7314111797897617339
