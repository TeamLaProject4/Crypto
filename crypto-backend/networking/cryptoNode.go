package networking

import (
	"context"
	"cryptomunt/blockchain"
	"cryptomunt/utils"
	"flag"
	"github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const RANDEVOUS_STRING = "cryptomunt-randevous"

type TopicType string

const (
	TRANSACTION TopicType = "TRANSACTION" //transaction
	BLOCKCHAIN            = "BLOCKCHAIN"
	ETC                   = "ETC"
)

type Config struct {
	//Port           int
	//ProtocolID     string
	//Rendezvous     string
	//Seed           int64
	BootNodes addrList
}

type CryptoNode struct {
	Libp2pNode    host.Host
	ctx           context.Context
	subscriptions []Subscription
	Blockchain    blockchain.Blockchain
	//sub map[TopicType]Subscription
}

func CreateCryptoNode() CryptoNode {
	utils.Logger.Info("Starting network")
	ctx := context.Background()
	config := parseFlags()

	node := initHost(ctx, config.BootNodes)

	cryptoNode := CryptoNode{
		Libp2pNode: node,
		ctx:        ctx,
	}

	cryptoNode.logNodeAddr()
	cryptoNode.subscriptions = cryptoNode.subscribeToTopics()
	//cryptoNode.readSubscriptions() //log incomin messages

	return cryptoNode
}

func (cryptoNode *CryptoNode) GetBlockChainFromNetwork() {
	peerstore := cryptoNode.Libp2pNode.Peerstore()

	peers := peerstore.PeersWithAddrs()
	utils.Logger.Error("0th peer", peers[0])

	//get  ipaddr from peer info
	ipADRESS := peerstore.PeerInfo(peers[2])
	utils.Logger.Error("ipaddr 1", ipADRESS)
	utils.Logger.Error("ipaddr 2t ", ipADRESS.Addrs)
}

func (cryptoNode *CryptoNode) WriteToTopic(data string, topicType TopicType) {
	for _, subscription := range cryptoNode.subscriptions {
		if subscription.TopicName == string(topicType) {
			err := subscription.Publish(data)
			if err != nil {
				utils.Logger.Error(err)
			}
		}
	}
}

//func (cryptoNode *CryptoNode) readSubscriptions() {
//	for _, subscription := range cryptoNode.subscriptions {
//		go readSubscription(&subscription)
//	}
//}

func parseFlags() Config {
	config := Config{}
	flag.Var(&config.BootNodes, "peer", "Peer multiaddress for peer discovery")
	flag.Parse()
	return config
}

func (cryptoNode *CryptoNode) logNodeAddr() {
	utils.Logger.Infof("Host ID: %s", cryptoNode.Libp2pNode.ID().Pretty())
	utils.Logger.Infof("Connect to me on:")
	for _, addr := range cryptoNode.Libp2pNode.Addrs() {
		utils.Logger.Infof("  %s/p2p/%s", addr, cryptoNode.Libp2pNode.ID().Pretty())
	}
}

func (cryptoNode *CryptoNode) subscribeToTopics() []Subscription {
	//main pub/sub object
	gossipPubSub, err := pubsub.NewGossipSub(cryptoNode.ctx, cryptoNode.Libp2pNode)

	//subscribtion topics
	transactionSub, err := subscribeToTopic(cryptoNode.ctx, gossipPubSub, cryptoNode.Libp2pNode.ID(), TRANSACTION)
	blockChainSub, err := subscribeToTopic(cryptoNode.ctx, gossipPubSub, cryptoNode.Libp2pNode.ID(), BLOCKCHAIN)

	if err != nil {
		panic(err)
	}

	go readSubscription(transactionSub)
	go readSubscription(blockChainSub)

	return []Subscription{*transactionSub, *blockChainSub}
}

func readSubscription(sub *Subscription) {
	for message := range sub.Messages {
		utils.Logger.Info(message.Message, sub.TopicName)
	}
}
