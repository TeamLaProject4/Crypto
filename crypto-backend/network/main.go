package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"os"
)

var logger = log.Logger("cryptomunt")

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

//create new host
func initHost() host.Host {
	//host, err := libp2p.New(libp2p.ListenAddrs([]multiaddr.Multiaddr(config.ListenAddresses)...))
	host, err := libp2p.New()
	if err != nil {
		panic(err)
	}

	logger.Info("Host created. We are:", host.ID())
	logger.Info("address: ", host.Addrs())

	// Set a function as stream handler. This function is called when a peer
	// initiates a connection and starts a stream with this peer.
	//protocol id is a unique string on which hosts can agree how to communicate(protocol)?
	host.SetStreamHandler("/cryptomunt/1.0.0", handleStream)

	initDHT(host)
	return host
}
func main() {
	initLogger()
	node := initHost()
	logger.Info("test ", node.ID())

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
