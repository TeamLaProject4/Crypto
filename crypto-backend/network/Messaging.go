package main

import (
	"bufio"
	"fmt"
	"github.com/libp2p/go-libp2p-core/network"
)

func handleStream(stream network.Stream, readMessages chan string, writeMessages chan string) {
	logger.Info("Got a new stream!")

	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	go readData(rw, readMessages)
	go writeData(rw, writeMessages)
}

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

func writeData(rw *bufio.ReadWriter, writeMessages chan string) {
	//for {
	logger.Info("in writedata function")
	for sendData := range writeMessages {
		logger.Info("sending data")
		_, err := rw.WriteString(fmt.Sprintf("%s\n", sendData))
		logger.Info("data sent")
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
	//}

	//stdReader := bufio.NewReader(os.Stdin)
	//for {
	//	sendData := <-writeMessages
	//	logger.Info("sending data")
	//	_, err := rw.WriteString(fmt.Sprintf("%s\n", sendData))
	//	logger.Info("data sent")
	//	if err != nil {
	//		fmt.Println("Error writing to buffer")
	//		panic(err)
	//	}
	//	err = rw.Flush()
	//	if err != nil {
	//		fmt.Println("Error flushing buffer")
	//		panic(err)
	//	}
	//	//fmt.Print("> ")
	//	//sendData, err := stdReader.ReadString('\n')
	//	//if err != nil {
	//	//	fmt.Println("Error reading from stdin")
	//	//	panic(err)
	//	//}
	//	//
	//	//_, err = rw.WriteString(fmt.Sprintf("%s\n", sendData))
	//	//if err != nil {
	//	//	fmt.Println("Error writing to buffer")
	//	//	panic(err)
	//	//}
	//	//err = rw.Flush()
	//	//if err != nil {
	//	//	fmt.Println("Error flushing buffer")
	//	//	panic(err)
	//	//}
	//}
}
