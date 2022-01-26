package main

import (
	"context"
	"flag"
	"github.com/ipfs/go-log/v2"
)

const RANDEVOUS_STRING = "cryptomunt-randevous"
const PROTOCOL_ID = "/cryptomunt/1.0.0"

type NetworkModel struct {
	ReadMessages  chan string
	WriteMessages chan string
}
type Utils struct {
	Logger *log.ZapEventLogger
}

var Logger = log.Logger("cryptomunt")
var utils = Utils{Logger: Logger}

func InitLogger() {
	log.SetAllLoggers(log.LevelWarn)
	err := log.SetLogLevel("cryptomunt", "info")
	if err != nil {
		return
	}
}

func CreateNetwork() {
	config := Config{}
	ctx := context.Background()
	InitLogger()

	//TODO: figure out channel buffer size
	//TODO: dont remove from channel until all peers have received a message?
	readMessages := make(chan string, 1000)
	writeMessages := make(chan string, 1000)

	flag.Var(&config.DiscoveryPeers, "peer", "Peer multiaddress for peer discovery")
	flag.Parse()

	node := initHost(ctx, config.DiscoveryPeers, readMessages, writeMessages)
	utils.Logger.Infof("Host ID: %s", node.ID().Pretty())
	utils.Logger.Infof("Connect to me on:")
	for _, addr := range node.Addrs() {
		utils.Logger.Infof("  %s/p2p/%s", addr, node.ID().Pretty())
	}

	//use channels to communicate with goroutines for each peer
	go printDataFromPeers(readMessages)
	//go sendDataToPeers(writeMessages)

	select {}
	//return NetworkModel{
	//	ReadMessages:  readMessages,
	//	WriteMessages: writeMessages,
	//}
	//sleep forever
	//select {}
	//return
}

//temporary
func printDataFromPeers(readMessages chan string) {
	//TODO: handle messages in blockchain implementation, maybe return this channel?
	for message := range readMessages {
		utils.Logger.Info("CHANNEL READ MESSAGE: " + message)
		//blockchain.createTransaction()
	}
}

func (NetworkModel *NetworkModel) SendDataToPeers(jsonData string) {
	NetworkModel.WriteMessages <- jsonData
}
