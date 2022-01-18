package main

import (
	"bufio"
	"context"
	"flag"
	"github.com/ipfs/go-log/v2"
	"os"
)

var logger = log.Logger("cryptomunt")

const RANDEVOUS_STRING = "cryptomunt-randevous"
const PROTOCOL_ID = "/cryptomunt/1.0.0"

func main() {
	config := Config{}
	ctx := context.Background()

	//TODO: figure out channel buffer size
	readMessages := make(chan string, 100)
	writeMessages := make(chan string, 100)

	initLogger()

	flag.Var(&config.DiscoveryPeers, "peer", "Peer multiaddress for peer discovery")
	flag.Parse()

	node := initHost(ctx, config.DiscoveryPeers, readMessages, writeMessages)

	logger.Infof("Host ID: %s", node.ID().Pretty())
	logger.Infof("Connect to me on:")
	for _, addr := range node.Addrs() {
		logger.Infof("  %s/p2p/%s", addr, node.ID().Pretty())
	}

	//use channels to communicate with goroutines for each peer
	go printDataFromPeers(readMessages)
	go sendDataToPeers(writeMessages)

	//sleep forever
	select {}
}

//temporary
func printDataFromPeers(readMessages chan string) {
	//TODO: handle messages in blockchain implementation, maybe return this channel?
	for message := range readMessages {
		logger.Info("CHANNEL READ MESSAGE: " + message)
	}
}

//temporary
func sendDataToPeers(writeMessages chan string) {
	//TODO: input json data from blockchain
	stdReader := bufio.NewReader(os.Stdin)
	for {
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			logger.Warn("Read input error")
		}
		writeMessages <- sendData
	}
}

func initLogger() {
	log.SetAllLoggers(log.LevelWarn)
	err := log.SetLogLevel("cryptomunt", "info")
	if err != nil {
		return
	}
}
