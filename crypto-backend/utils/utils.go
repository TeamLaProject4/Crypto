package utils

import (
	"crypto/sha256"
	"encoding/hex"
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
