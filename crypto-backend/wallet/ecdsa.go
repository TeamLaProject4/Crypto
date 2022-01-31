package wallet

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"github.com/jbenet/go-base58"
)

func ecdsaVerify(publicKey *ecdsa.PublicKey, data string, signature string) bool {
	hash := sha256.Sum256([]byte(data))

	sigBytes := base58.Decode(signature)

	return ecdsa.VerifyASN1(publicKey, hash[:], sigBytes)
}

func EncodePrivateKey(privateKey *ecdsa.PrivateKey) string {
	x509EncodedPriv, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncodedPriv := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509EncodedPriv})
	return string(pemEncodedPriv)
}

func EncodePublicKey(publicKey *ecdsa.PublicKey) string {
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	return string(pemEncodedPub)
}

func DecodePrivateKey(pemEncodedPriv string) *ecdsa.PrivateKey {
	blockPriv, _ := pem.Decode([]byte(pemEncodedPriv))
	x509EncodedPriv := blockPriv.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509EncodedPriv)

	return privateKey
}

func DecodePublicKey(pemEncodedPub string) *ecdsa.PublicKey {
	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return publicKey
}
