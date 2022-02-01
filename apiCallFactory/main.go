package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

type factoryFlags struct {
	transactions int
	ip           string
}

func ReadFileHex(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return hex.EncodeToString(content)
}

func makeTransaction(ip string) {
	//recieverPublicKey, amount	transactionType(transfer)
	params := url.Values{}
	index := rand.Intn(9) + 1
	pubReceiverKey := ReadFileHex(fmt.Sprintf("../crypto-backend/keys/demo-keys/wallet-pubkey-%d.txt", index))

	params.Add("recieverPublicKey", pubReceiverKey)
	params.Add("amount", strconv.Itoa(rand.Intn(100)))
	params.Add("transactionType", "transfert")

	reader := io.Reader()
	_, err := http.Post("http://"+ip+"/frontend/transaction", "url", reader)
	if err != nil {
		fmt.Println(err)
	}
}

//func loginToWallet(ip string) {
//	//pubkey: 0409c7592e1ad0b738fc7ae71388502c0880580368b50bfae7334ee84e8b8b898f1c55ce10b7a7e33baf1fe31db4268288ef1df21ad8780c90a66e8040dddb87a1
//	mnemonic := "tenant ostrich nation lift screen inside whisper replace foam correct tree cool little announce correct excess slogan term actor crystal scout innocent viable fix"
//	params := url.Values{}
//	params.Add("Mnemonic", mnemonic)
//	http.post
//	_, err := http.PostForm("http://"+ip+"/frontend/confirmMnemonic", params)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
func loginToWallet(ip string) {
	//pubkey: 0409c7592e1ad0b738fc7ae71388502c0880580368b50bfae7334ee84e8b8b898f1c55ce10b7a7e33baf1fe31db4268288ef1df21ad8780c90a66e8040dddb87a1
	//mnemonic := "tenant ostrich nation lift screen inside whisper replace foam correct tree cool little announce correct excess slogan term actor crystal scout innocent viable fix"
	httpposturl := "http://" + ip + "/frontend/confirmMnemonic"

	var jsonData = []byte(`{
		"mnemonic": "tenant ostrich nation lift screen inside whisper replace foam correct tree cool little announce correct excess slogan term actor crystal scout innocent viable fix",
	}`)
	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
}

func main() {
	config := parseFlags()
	loginToWallet(config.ip)
	for i := 0; i < config.transactions; i++ {
		makeTransaction(config.ip)
	}
}

func parseFlags() factoryFlags {
	config := factoryFlags{}
	flag.IntVar(&config.transactions, "amount", 0, "amount of transactions to make")
	flag.StringVar(&config.ip, "ip", "", "Ip address with port, eg 10.2.3.1:8347")
	flag.Parse()
	return config
}
