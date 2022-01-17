package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/multiformats/go-multiaddr"
	"os"
	"strings"
	"sync"
)

var logger = log.Logger("cryptomunt")

const RANDEVOUS_STRING = "cryptomunt-randevous"
const PROTOCOL_ID = "/cryptomunt/1.0.0"

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

func main() {
	config := Config{}
	ctx, _ := context.WithCancel(context.Background())
	initLogger()

	flag.Var(&config.DiscoveryPeers, "peer", "Peer multiaddress for peer discovery")
	flag.Parse()

	node := initHost(ctx, config.DiscoveryPeers)

	logger.Infof("Host ID: %s", node.ID().Pretty())
	logger.Infof("Connect to me on:")
	for _, addr := range node.Addrs() {
		logger.Infof("  %s/p2p/%s", addr, node.ID().Pretty())
	}

	//select so that the streams can be handled
	select {}
}

func initLogger() {
	log.SetAllLoggers(log.LevelWarn)
	err := log.SetLogLevel("cryptomunt", "info")
	if err != nil {
		return
	}
}

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
func initRoutingDiscovery(ctx context.Context, kademliaDHT *dht.IpfsDHT, node host.Host) {
	//Announce node on the network with randevous point
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

	connectToPeers(ctx, peerChan, node)
}

func connectToPeers(ctx context.Context, peerChan <-chan peer.AddrInfo, node host.Host) {
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

			go writeData(rw)
			go readData(rw)
		}

		logger.Info("Connected to:", peerNode)
	}
}

func initHost(ctx context.Context, bootstrapPeers []multiaddr.Multiaddr) host.Host {
	node, err := libp2p.New()
	if err != nil {
		panic(err)
	}
	logger.Info("Host created. We are:", node.ID())
	logger.Info("address: ", node.Addrs())

	//set streamhandler with unique protocol id
	node.SetStreamHandler(protocol.ID(PROTOCOL_ID), handleStream)

	//init dht
	kademliaDHT, initDHTErr := initKDHT(ctx, node, bootstrapPeers)
	if initDHTErr != nil {
		logger.Error("dht error")
		return nil
	}

	initRoutingDiscovery(ctx, kademliaDHT, node)

	return node
}

func handleStream(stream network.Stream) {
	logger.Info("Got a new stream!")

	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	go readData(rw)
	go writeData(rw)
}

func readData(rw *bufio.ReadWriter) {
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
		}

	}
}

func writeData(rw *bufio.ReadWriter) {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin")
			panic(err)
		}

		_, err = rw.WriteString(fmt.Sprintf("%s\n", sendData))
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
}
