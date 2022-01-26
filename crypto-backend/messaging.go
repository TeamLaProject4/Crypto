package main

import (
	"bufio"
	//"cryptomunt/utils"
	"fmt"
	"github.com/libp2p/go-libp2p-core/network"
)

func handleStream(stream network.Stream, readMessages chan string, writeMessages chan string) {
	utils.Logger.Info("Got a new stream!")

	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	go readData(rw, readMessages)
	go writeData(rw, writeMessages)
}

//Read data received from peers and put it in the messages channel
func readData(rw *bufio.ReadWriter, messages chan string) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from buffer")
			panic(err)
		}

		if str == "" {
			return
		}
		if str != "\n" {
			// Green console colour: 	\x1b[32m
			// Reset console colour: 	\x1b[0m
			fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
			messages <- str
		}

	}
}

//Write data to peers when the writeMessages channel receives new values
func writeData(rw *bufio.ReadWriter, writeMessages chan string) {
	//utils.Logger.Info("in writedata function")
	for sendData := range writeMessages {
		utils.Logger.Info("sending data")
		_, err := rw.WriteString(fmt.Sprintf("%s\n", sendData))
		utils.Logger.Info("data sent")

		if err != nil {
			fmt.Println("Error writing to buffer")
			panic(err)
		}
		err = rw.Flush()
		if err != nil {
			fmt.Println("Error flushing buffer")
			panic(err)
		}
	}

}
