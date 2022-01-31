package utils

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"os"
)

func GetHexadecimalHash(valueToHash string) string {
	sha256Hasher := sha256.New()
	sha256Hasher.Write([]byte(valueToHash))
	hashBytes := sha256Hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

func GetBigIntHash(valueToHash string) big.Int {
	return HexToBigInt(GetHexadecimalHash(valueToHash))
}

func HexToBigInt(hex string) big.Int {
	number := new(big.Int)
	number, ok := number.SetString(hex, 16)
	if !ok {
		panic("error hexTobigInt")
	}
	return *number
}

func GetAbsolutBigInt(number big.Int) big.Int {
	var newNumber = new(big.Int)
	return *newNumber.Abs(&number)
}

func WriteFile(filePath string, content string) {
	err := os.WriteFile(filePath, []byte(content), 0o777)
	if err != nil {
		panic("Cannot write file")
	}
}

func ReadFileBytes(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		Logger.Info(err)
	}
	return content
}

func GetPublicKeyFromHex(hexValue string) (rsa.PublicKey, error) {
	var publicKeyEmpty rsa.PublicKey
	pubPem, _ := hex.DecodeString(hexValue)
	block, _ := pem.Decode(pubPem)

	if block == nil {
		return publicKeyEmpty, nil
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return *publicKey, err
	}

	return *publicKey, nil
}
