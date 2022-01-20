package main

import (
	"errors"
	"time"
)

var (
	SENDER_PUBLIC_KEY   = "alice"
	RECEIVER_PUBLIC_KEY = "bob"
	AMOUNT              = 10
	TX_TYPE             = TxType.TRANSFER
	SIGNATURE           = "VeRyN1CeSiGn4TuRe"
)

func create_transaction() Transaction {
	return Transaction(SENDER_PUBLIC_KEY, RECEIVER_PUBLIC_KEY, AMOUNT, TX_TYPE)
}

func test_when_new_transaction_constructed_then_signature_is_empty() {
	tx := create_transaction()
	if !(tx.signature == "") {
		panic(errors.New("AssertionError"))
	}
}

func test_when_duplicate_transactions_constructed_then_transactions_are_equal() {
	id := uuid.uuid4().hex
	timestamp := int(float64(time.Now().UnixNano()) / 1000000000.0)
	tx1 := Transaction(SENDER_PUBLIC_KEY, RECEIVER_PUBLIC_KEY, AMOUNT, TX_TYPE, id, timestamp)
	tx2 := Transaction(SENDER_PUBLIC_KEY, RECEIVER_PUBLIC_KEY, AMOUNT, TX_TYPE, id, timestamp)
	if !(tx1 == tx2) {
		panic(errors.New("AssertionError"))
	}
}

func test_when_transaction_signed_then_signature_is_set() {
	tx := create_transaction()
	tx.sign(SIGNATURE)
	if !(tx.signature == SIGNATURE) {
		panic(errors.New("AssertionError"))
	}
}

func test_when_transaction_signed_then_payload_signature_stays_empty() {
	tx := create_transaction()
	tx.sign(SIGNATURE)
	payload := tx.payload()
	if !(payload["signature"] == "") {
		panic(errors.New("AssertionError"))
	}
}
