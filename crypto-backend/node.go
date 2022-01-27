package main

import (
	"bufio"
	"cryptomunt/blockchain"
	"cryptomunt/utils"
	"fmt"
	"os"
)

type Node struct {
	networkChannels NetworkChannels
	blockchain      blockchain.Blockchain
	config          Config
}

func CreateNode(config Config) Node {
	networkChannels := CreateNetwork(config)
	nodeBlockchain := blockchain.CreateBlockchain()

	node := Node{
		networkChannels: networkChannels,
		blockchain:      nodeBlockchain,
		config:          config,
	}

	return node
}

func (node *Node) StartP2pNetwork() NetworkChannels {
	utils.InitLogger()
	networkChannels := CreateNetwork(node.config)
	go node.writeDataToPeers()

	return networkChannels
	//keep running forever
	//select {}
}

func (node *Node) writeDataToPeers() {
	stdReader := bufio.NewReader(os.Stdin)
	for {
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin")
			panic(err)
		}
		node.networkChannels.SendDataToPeers(sendData)
	}
}
