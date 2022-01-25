package blockchain

import (
	"fmt"
	"encoding/json"
	"testing"
	"time"
	"github.com/google/uuid"
)

const SENDER_PUBLIC_KEY = "alice"
const RECEIVER_PUBLIC_KEY = "bob"
const AMOUNT = 10
const TX_TYPE = TRANSFER
const SIGNATURE = "VeRyN1CeSiGn4TuRe"

// want := Transaction{
// 	SenderPublicKey:	"alice"
// 	ReceiverPublicKey:	"bob"
// 	Amount:				10
// 	TransactionType:  	TRANSFER
// 	Id:					0
// 	Timestamp:			0
// 	Signature:			"VeRyN1CeSiGn4TuRe"
// }

func constructTransaction() Transaction {
	return Transaction{
		SenderPublicKey:	SENDER_PUBLIC_KEY,
		ReceiverPublicKey:	RECEIVER_PUBLIC_KEY,
		Amount:				AMOUNT,
		TransactionType:  	TX_TYPE,
	}
}

func TestWhenTransactionConstructedThenSignatureIsEmpty(t *testing.T) {
	tx := constructTransaction()
	
	got := tx.Signature
	want := ""
    if want != got {
        t.Errorf("Expected '%s', but got '%s'", want, got)
    }
}

func TestWhenDuplicateTransactionsConstructedThenTransactionsAreEqual(t *testing.T) {
	id := uuid.New().String()
	timestamp := time.Now().Unix()
	
	tx1 := constructTransaction()
	tx2 := constructTransaction()
	tx1.Id = id
	tx2.Id = id
	tx1.Timestamp = timestamp
	tx2.Timestamp = timestamp

    if tx1 != tx2 {
        t.Errorf("Expected '%+v', but got '%+v'", tx1, tx2)
    }
}

func TestWhenTransactionSignedThenSignatureIsSet(t *testing.T) {
	tx := constructTransaction()
	tx.sign(SIGNATURE)

	got := tx.Signature
	want := SIGNATURE
    if got != want {
        t.Errorf("Expected '%s', but got '%s'", want, got)
    }
}

func TestWhenTransactionSignedThenPayloadSignatureStaysEmpty(t *testing.T) {
	tx := constructTransaction()
	tx.sign(SIGNATURE)
	payload := tx.payload()

	var result map[string]interface{}
	json.Unmarshal([]byte(payload), &result)

	got := result["Signature"]
	want := ""
    if got != want {
        t.Errorf("Expected '%v', but got '%v'", want, got)
    }
}
