package main

import (
	"context"
	"flag"
	"github.com/ipfs/go-log/v2"
)

var logger = log.Logger("cryptomunt")

const RANDEVOUS_STRING = "cryptomunt-randevous"
const PROTOCOL_ID = "/cryptomunt/1.0.0"

func main() {
	config := Config{}
	ctx := context.Background()
	messages := make(chan string, 100)
	initLogger()

	flag.Var(&config.DiscoveryPeers, "peer", "Peer multiaddress for peer discovery")
	flag.Parse()

	node := initHost(ctx, config.DiscoveryPeers, messages)

	logger.Infof("Host ID: %s", node.ID().Pretty())
	logger.Infof("Connect to me on:")
	for _, addr := range node.Addrs() {
		logger.Infof("  %s/p2p/%s", addr, node.ID().Pretty())
	}

	//TODO: handle messages in blockchain implementation, maybe return this channel?
	for message := range messages {
		logger.Info("CHANNEL MESSAGE: " + message)
	}

	//send into channel?

	//sleep forever
	select {}
}

func initLogger() {
	log.SetAllLoggers(log.LevelWarn)
	err := log.SetLogLevel("cryptomunt", "info")
	if err != nil {
		return
	}
}
