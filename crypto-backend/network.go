package main

import (
	"context"
	"cryptomunt/utils"
	"github.com/ipfs/go-log/v2"
)

const RANDEVOUS_STRING = "cryptomunt-randevous"
const PROTOCOL_ID = "/cryptomunt/1.0.0"

type NetworkChannels struct {
	ReadMessages  chan string
	WriteMessages chan string
}

var Logger = log.Logger("cryptomunt")

func CreateNetwork(config Config) NetworkChannels {
	utils.Logger.Info("Starting network")

	//config := Config{}
	ctx := context.Background()

	//TODO: figure out channel buffer size
	//TODO: dont remove from channel until all peers have received a message?
	readMessages := make(chan string, 1000)
	writeMessages := make(chan string, 1000)

	//flag.Var(&config.DiscoveryPeers, "peer", "Peer multiaddress for peer discovery")
	//flag.Parse()

	node := initHost(ctx, config.DiscoveryPeers, readMessages, writeMessages)
	utils.Logger.Infof("Host ID: %s", node.ID().Pretty())
	utils.Logger.Infof("Connect to me on:")
	for _, addr := range node.Addrs() {
		utils.Logger.Infof("  %s/p2p/%s", addr, node.ID().Pretty())
	}

	//use channels to communicate with goroutines for each peer
	//go printDataFromPeers(readMessages)
	//go sendDataToPeers(writeMessages)

	return NetworkChannels{
		ReadMessages:  readMessages,
		WriteMessages: writeMessages,
	}
	//sleep forever
	//select {}
}

func (NetworkModel *NetworkChannels) SendDataToPeers(jsonData string) {
	NetworkModel.WriteMessages <- jsonData
}
