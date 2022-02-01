package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"cryptomunt/wallet"
)

func main2() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	wallett := wallet.CreateWallet(*privateKey)

	pemEncodedPriv := wallet.EncodePrivateKey(&wallett.Key)
	pemEncodedPub := wallet.EncodePublicKey(&wallett.Key.PublicKey)

	wallet.WriteKeyToFile(wallet.PRIVATE_KEY_PATH, pemEncodedPriv)
	wallet.WriteKeyToFile(wallet.PUBLIC_KEY_PATH, pemEncodedPub)

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
