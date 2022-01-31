package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	// "crypto/sha256"
	"cryptomunt/wallet"
	// "encoding/hex"
	"fmt"
)

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	wallett := wallet.CreateWallet(*privateKey)
	msg := "hello, world"

	sig := wallett.Sign(msg)

	fmt.Printf("signature: %s\n", sig)
	valid := wallet.IsValidSignature(msg, sig, wallett.GetPublicKeyHex())

	fmt.Println(valid)

	// hash := sha256.Sum256([]byte(msg))

	// sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("signature: %x\n", sig)
	// sigString := hex.EncodeToString(sig)
	// hexBytes, err := hex.DecodeString(sigString)
	// fmt.Printf("signature hex: %x\n", hexBytes)

	// valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sig)
	// fmt.Println("signature verified:", valid)
}