package main

import (
	"github.com/multiformats/go-multiaddr"
	"strings"
)

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
	DiscoveryPeers addrList
}
