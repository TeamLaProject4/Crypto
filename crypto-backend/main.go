package main

import (
	"bufio"
	"cryptomunt/utils"
	"fmt"
	"os"
)

func main() {
	utils.InitLogger()
	cryptoNode := CreateCryptoNode()

	go func() {
		stdReader := bufio.NewReader(os.Stdin)
		for {
			sendData, err := stdReader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading from stdin")
				panic(err)
			}
			fmt.Println("writing to topic...")
			cryptoNode.WriteToTopic(sendData, BLOCKCHAIN)
		}
	}()

	select {}
}
