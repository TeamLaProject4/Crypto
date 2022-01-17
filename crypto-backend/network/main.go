package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/multiformats/go-multiaddr"
	"io"
	mrand "math/rand"
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

	//run(node, cancel)
	select {}
}

func initLogger() {
	log.SetAllLoggers(log.LevelWarn)
	err := log.SetLogLevel("cryptomunt", "info")
	if err != nil {
		return
	}
}

// Start a DHT, for use in peer discovery. We can't just make a new DHT
// client because we want each peer to maintain its own local copy of the
// DHT, so that the bootstrapping node of the DHT can go down without
// inhibiting future peer discovery.
func initDHT(host host.Host) {
	ctx := context.Background()
	kademliaDHT, err := dht.New(ctx, host)
	if err != nil {
		panic(err)
	}

	// Bootstrap the DHT. In the default configuration, this spawns a Background
	// thread that will refresh the peer table every five minutes.
	logger.Debug("Bootstrapping the DHT")
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		panic(err)
	}
	// Let's connect to the bootstrap nodes first. They will tell us about the
	// other nodes in the network.
	//<blocking code>o
	//peerAddr := multiaddr.StringCast("/ip4/127.0.0.1/tcp/3932")
	//logger.Info("peeraddr: ", peerAddr)
	//peerinfo, err := peer.AddrInfoFromP2pAddr(peerAddr)
	//if err != nil {
	//	logger.Warn("peerInfo is nil ", peerinfo)
	//}
	//logger.Info("peerInfo: ", peerinfo)

	peerAddr := dht.DefaultBootstrapPeers
	peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr[0])
	err = host.Connect(ctx, *peerinfo)
	logger.Info("is there an error", err)

	//var wg sync.WaitGroup
	//for _, peerAddr := range config.BootstrapPeers {
	//	peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		if err := host.Connect(ctx, *peerinfo); err != nil {
	//			logger.Warn(err)
	//		} else {
	//			logger.Info("Connection established with bootstrap node:", *peerinfo)
	//		}
	//	}()
	//}
	//wg.Wait()
	//</blocking code>

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

func NewHost(seed int64, port int) (host.Host, error) {

	// If the seed is zero, use real cryptographic randomness. Otherwise, use a
	// deterministic randomness source to make generated keys stay the same
	// across multiple runs
	var r io.Reader
	if seed == 0 {
		r = rand.Reader
	} else {
		r = mrand.New(mrand.NewSource(seed))
	}

	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, err
	}

	addr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))

	return libp2p.New(
		libp2p.ListenAddrs(addr),
		libp2p.Identity(priv),
	)
}

//create new host
func initHost(ctx context.Context, bootstrapPeers []multiaddr.Multiaddr) host.Host {
	//host, err := libp2p.New(libp2p.ListenAddrs([]multiaddr.Multiaddr(config.ListenAddresses)...))
	node, err := libp2p.New()
	//node, err := NewHost(0, 9784)
	if err != nil {
		panic(err)
	}
	logger.Info("Host created. We are:", node.ID())
	logger.Info("address: ", node.Addrs())

	// Set a function as stream handler. This function is called when a peerNode
	// initiates a connection and starts a stream with this peerNode.
	//protocol id is a unique string on which hosts can agree how to communicate(protocol)?
	node.SetStreamHandler(protocol.ID(PROTOCOL_ID), handleStream)
	//node.SetStreamHandler("/cryptomunt/1.0.0", handleStream)

	//initDHT(host)
	//init dht
	kademliaDHT, initDHTErr := initKDHT(ctx, node, bootstrapPeers)
	if initDHTErr != nil {
		logger.Error("dht error")
		return nil
	}

	//// We use a rendezvous point to announce our location.
	//// This is like telling your friends to meet you at the Eiffel Tower.
	//logger.Info("Announcing ourselves...")
	////routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)
	//go Discover(ctx, node, kademliaDHT, RANDEVOUS_STRING)
	////discovery.Advertise(ctx, routingDiscovery, RANDEVOUS_STRING)
	//logger.Debug("Successfully announced!")

	// We use a rendezvous point "meet me here" to announce our location.
	// This is like telling your friends to meet you at the Eiffel Tower.
	logger.Info("Announcing ourselves...")
	routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)
	discovery.Advertise(ctx, routingDiscovery, RANDEVOUS_STRING)
	logger.Debug("Successfully announced!")

	// Now, look for others who have announced
	// This is like your friend telling you the location to meet you.
	logger.Debug("Searching for other peers...")
	peerChan, err := routingDiscovery.FindPeers(ctx, RANDEVOUS_STRING)
	if err != nil {
		panic(err)
	}

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

	return node
}

func handleStream(stream network.Stream) {
	logger.Info("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	go readData(rw)
	go writeData(rw)

	// 'stream' will stay open until you close it (or the other side closes it).
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
