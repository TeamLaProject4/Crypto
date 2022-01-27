package main

import (
	"context"
	"cryptomunt/utils"
	"flag"
	"github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-multiaddr"
	"strings"
)

const RANDEVOUS_STRING = "cryptomunt-randevous"

var Logger = log.Logger("cryptomunt")

func CreateNetwork() {
	utils.Logger.Info("Starting network")

	config := Config{}
	ctx := context.Background()
	flag.Var(&config.BootNodes, "peer", "Peer multiaddress for peer discovery")
	flag.Parse()

	node := initHost(ctx, config.BootNodes)
	utils.Logger.Infof("Host ID: %s", node.ID().Pretty())
	utils.Logger.Infof("Connect to me on:")
	for _, addr := range node.Addrs() {
		utils.Logger.Infof("  %s/p2p/%s", addr, node.ID().Pretty())
	}

}

type addrList []multiaddr.Multiaddr

func (al *addrList) String() string {
	strs := make([]string, len(*al))
	for i, addr := range *al {
		strs[i] = addr.String()
	}
	return strings.Join(strs, ",")
}

func (al *addrList) Set(value string) error {
	addr, err := multiaddr.NewMultiaddr(value)
	if err != nil {
		return err
	}
	*al = append(*al, addr)
	return nil
}

type Config struct {
	//Port           int
	//ProtocolID     string
	//Rendezvous     string
	//Seed           int64
	BootNodes addrList
}
