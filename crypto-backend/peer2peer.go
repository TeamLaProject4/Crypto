package main

import (
	"context"
	"cryptomunt/utils"
	//"cryptomunt/utils"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
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
				utils.Logger.Errorf("Error while connecting to node %q: %-v", peerinfo, err)
			} else {
				utils.Logger.Infof("Connection established with bootstrap node: %q", *peerinfo)
			}
		}()
	}
	wg.Wait()

	return kademliaDHT, nil
}

//init routing Discovery and connect to peers in the network
func initRoutingDiscovery(ctx context.Context, kademliaDHT *dht.IpfsDHT, node host.Host) {
	//Announce node on the network with randevous point every 3 hours
	utils.Logger.Info("Announcing ourselves...")
	routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)
	discovery.Advertise(ctx, routingDiscovery, RANDEVOUS_STRING)
	utils.Logger.Debug("Successfully announced!")

	//Search for nodes on the network with randevous point
	utils.Logger.Debug("Searching for other peers...")
	peerChan, err := routingDiscovery.FindPeers(ctx, RANDEVOUS_STRING)
	if err != nil {
		panic(err)
	}

	connectToPeers(ctx, peerChan, node)
}

func connectToPeers(ctx context.Context, peerChan <-chan peer.AddrInfo, node host.Host) {
	for peerNode := range peerChan {
		//filter out yourself
		if peerNode.ID == node.ID() {
			continue
		}

		utils.Logger.Debug("Found peerNode:", peerNode)
		utils.Logger.Debug("Connecting to:", peerNode)

		//connect to peer
		err := node.Connect(ctx, peerNode)

		if err != nil {
			utils.Logger.Warn("Connection failed:", err)
			continue
		}

		utils.Logger.Info("Connected to:", peerNode)
	}
}

func initHost(ctx context.Context, bootstrapPeers []multiaddr.Multiaddr) host.Host {
	node, err := libp2p.New()
	if err != nil {
		panic(err)
	}

	utils.Logger.Info("Host created. We are:", node.ID())
	utils.Logger.Info("address: ", node.Addrs())

	gossipPubSub, err := pubsub.NewGossipSub(ctx, node)

	//init dht
	kademliaDHT, initDHTErr := initKDHT(ctx, node, bootstrapPeers)
	if initDHTErr != nil {
		utils.Logger.Error("dht error")
		panic(initDHTErr)
		return nil
	}

	initRoutingDiscovery(ctx, kademliaDHT, node)

	// join the chat room

	//sub to multiple topics
	//methode send to topic
	//extra: unsub from topic

	cr, err := JoinChatRoom(ctx, gossipPubSub, node.ID(), "", "ROOM")
	if err != nil {
		panic(err)
	}

	// draw the UI
	ui := NewChatUI(cr)
	if err = ui.Run(); err != nil {
		printErr("error running text UI: %s", err)
	}

	return node
}
