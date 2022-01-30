package main

import (
	"crypto/ecdsa"
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip32"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	//Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Seed: ", seed)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)

	// encPriv, encPub := encode(masterKey, publicKey)

	// fmt.Println(encPriv)
	// fmt.Println(encPub)

	// priv2, pub2 := decode(encPriv, encPub)

	// if !reflect.DeepEqual(masterKey, priv2) {
	// 	fmt.Println("Private keys do not match.")
	// }
	// if !reflect.DeepEqual(publicKey, pub2) {
	// 	fmt.Println("Public keys do not match.")
	// }

}

func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

// func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
// 	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
// 	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

// 	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
// 	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

// 	return string(pemEncoded), string(pemEncodedPub)
// }

func decode(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return privateKey, publicKey
}

// func test() {
// 	privateKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
// 	publicKey := &privateKey.PublicKey

// 	encPriv, encPub := encode(privateKey, publicKey)

// 	fmt.Println(encPriv)
// 	fmt.Println(encPub)

// 	priv2, pub2 := decode(encPriv, encPub)

// 	if !reflect.DeepEqual(privateKey, priv2) {
// 		fmt.Println("Private keys do not match.")
// 	}
// 	if !reflect.DeepEqual(publicKey, pub2) {
// 		fmt.Println("Public keys do not match.")
// 	}
// }
