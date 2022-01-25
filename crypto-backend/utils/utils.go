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

func BigIntToHex(number big.Int) string {
	bytes := number.Bytes()
	return hex.EncodeToString(bytes)
}

func GetAbsolutBigInt(number big.Int) big.Int {
	var newNumber = new(big.Int)
	return *newNumber.Abs(&number)
}

func GetFileContents(filePath string) string {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0o777)
	if err != nil {
		panic(err)
	}
	fileBytes, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		panic(err2)
	}

	return string(fileBytes)
}
