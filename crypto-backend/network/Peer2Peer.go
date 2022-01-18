package main

import (
	"bufio"
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/multiformats/go-multiaddr"
	"sync"
)

func initKDHT(ctx context.Context, host host.Host, bootstrapPeers []multiaddr.Multiaddr) (*dht.IpfsDHT, error) {
	var options []dht.Option

	if len(bootstrapPeers) == 0 {
		options = append(options, dht.Mode(dht.ModeServer))
	}

	kademliaDHT, err := dht.New(ctx, host, options...)
	if err != nil {
		return nil, err
	}

	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		return nil, err
	}

	//blocking code using WaitGroup
	var wg sync.WaitGroup
	for _, peerAddr := range bootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := host.Connect(ctx, *peerinfo); err != nil {
				logger.Errorf("Error while connecting to node %q: %-v", peerinfo, err)
			} else {
				logger.Infof("Connection established with bootstrap node: %q", *peerinfo)
			}
		}()
	}
	wg.Wait()

	return kademliaDHT, nil
}

//init routing Discovery and connect to peers in the network
func initRoutingDiscovery(ctx context.Context, kademliaDHT *dht.IpfsDHT, node host.Host, readMessages chan string, writeMessages chan string) {
	//Announce node on the network with randevous point every 3 hours
	logger.Info("Announcing ourselves...")
	routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)
	discovery.Advertise(ctx, routingDiscovery, RANDEVOUS_STRING)
	logger.Debug("Successfully announced!")

	//Search for nodes on the network with randevous point
	logger.Debug("Searching for other peers...")
	peerChan, err := routingDiscovery.FindPeers(ctx, RANDEVOUS_STRING)
	if err != nil {
		panic(err)
	}

	connectToPeers(ctx, peerChan, node, readMessages, writeMessages)
}

func connectToPeers(ctx context.Context, peerChan <-chan peer.AddrInfo, node host.Host, readMessages chan string, writeMessages chan string) {
	for peerNode := range peerChan {
		if peerNode.ID == node.ID() {
			continue
		}
		logger.Debug("Found peerNode:", peerNode)
		logger.Debug("Connecting to:", peerNode)

		stream, err := node.NewStream(ctx, peerNode.ID, protocol.ID(PROTOCOL_ID))

		if err != nil {
			logger.Warn("Connection failed:", err)
			continue
		} else {
			rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

			go writeData(rw, writeMessages)
			go readData(rw, readMessages)
		}

		logger.Info("Connected to:", peerNode)
	}
}

func initHost(ctx context.Context, bootstrapPeers []multiaddr.Multiaddr, readMessages chan string, writeMessages chan string) host.Host {
	node, err := libp2p.New()
	if err != nil {
		panic(err)
	}
	logger.Info("Host created. We are:", node.ID())
	logger.Info("address: ", node.Addrs())

	//set streamhandler with unique protocol id
	node.SetStreamHandler(protocol.ID(PROTOCOL_ID), func(stream network.Stream) {
		handleStream(stream, readMessages, writeMessages)
	})

	//init dht
	kademliaDHT, initDHTErr := initKDHT(ctx, node, bootstrapPeers)
	if initDHTErr != nil {
		logger.Error("dht error")
		return nil
	}

	initRoutingDiscovery(ctx, kademliaDHT, node, readMessages, writeMessages)

	return node
}
