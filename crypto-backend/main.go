package main

import (
	"bufio"
	"cryptomunt/blockchain"
	"cryptomunt/networking"
	"cryptomunt/utils"
	"fmt"
	"os"
)

func main() {
	utils.InitLogger()

	node := networking.CreateCryptoNode()

	//getBlockChainFromNetwork()
	node.Blockchain = blockchain.CreateBlockchain()

	fmt.Println(node.Libp2pNode.Addrs())
	peerstore := node.Libp2pNode.Peerstore()

	//utils.Logger.Error("peerstore", peerstore.Peers())
	peers := peerstore.PeersWithAddrs()
	//utils.Logger.Error("peers", peers)
	utils.Logger.Error("0th peer", peers[0])
	//

	//get  ipaddr from peer info
	ipADRESS := peerstore.PeerInfo(peers[2])
	utils.Logger.Error("ipaddr", ipADRESS)
	//then make api call for blockchain...

	//go tempWriteToTopic(node)

	select {}
}

func tempWriteToTopic(node networking.CryptoNode) {
	stdReader := bufio.NewReader(os.Stdin)
	for {
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin")
			panic(err)
		}
		fmt.Println("writing to topic...")
		node.WriteToTopic(sendData, networking.BLOCKCHAIN)
	}
}
