package utils

import (
	"bufio"
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
		return ""
	}
	fileBytes, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return ""
	}

	return string(fileBytes)
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

func WriteRsaKeyToFile(key rsa.PrivateKey) {
	// Extract public component.
	publicKey := key.Public()

	// Encode private key to PKCS#1 ASN.1 PEM.
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(&key),
		},
	)

	// Encode public key to PKCS#1 ASN.1 PEM.
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publicKey.(*rsa.PublicKey)),
		},
	)

	// Write private key to file.
	if err := ioutil.WriteFile("../keys/wallet.rsa", keyPEM, 0755); err != nil {
		panic(err)
	}

	// Write public key to file.
	if err := ioutil.WriteFile("../keys/wallet.rsa.publicKey", pubPEM, 0755); err != nil {
		panic(err)
	}
}

func ReadRsaKeyFile(filePath string) rsa.PrivateKey {
	//Logger.Info(os.Getwd())

	privateKeyFile, err := os.Open(filePath)

	pemFileInfo, err1 := privateKeyFile.Stat()
	if err1 != nil {
		Logger.Error("rsa error", err)
		return rsa.PrivateKey{}
	}
	var size = pemFileInfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode(pembytes)
	err = privateKeyFile.Close()
	if err != nil {
		Logger.Error("rsa error", err)
		return rsa.PrivateKey{}
	}

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		Logger.Error("rsa error", err)
		return rsa.PrivateKey{}
	}

	return *privateKeyImported
}

func GetRsaPublicKeyHexValue(pubkey *rsa.PublicKey) string {
	pubkeyBytes := x509.MarshalPKCS1PublicKey(pubkey)

	pubkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkeyBytes,
		},
	)

	return hex.EncodeToString(pubkeyPem)
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

//
//func GetStructFromJson(jsonData string, structData interface{}) interface{} {
//	err := json.Unmarshal([]byte(jsonData), &structData)
//	if err != nil {
//		Logger.Error("unmarshal error ", err)
//		return nil
//	}
//	return structData
//}
