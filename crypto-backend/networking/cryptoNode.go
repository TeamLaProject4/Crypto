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
	TRANSACTION     TopicType = "TRANSACTION" //transaction
	BLOCK_FORGED              = "BLOCK_FORGED"
	CONSENSUS_ERROR           = "CONSENSUS_ERROR"
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
	MemoryPool    blockchain.MemoryPool
	Wallet        blockchain.Wallet
	//sub map[TopicType]Subscription
}

func CreateCryptoNode() CryptoNode {
	utils.Logger.Info("Starting network")
	ctx := context.Background()
	config := parseFlags()

	//p2p
	node := initHost(ctx, config.BootNodes)
	//init
	cryptoNode := CryptoNode{
		Libp2pNode: node,
		ctx:        ctx,
	}

	//pubsub
	cryptoNode.logNodeAddr()
	cryptoNode.subscriptions = cryptoNode.subscribeToTopics()
	cryptoNode.readSubscriptions()

	//blockchain
	cryptoNode.Blockchain = blockchain.CreateBlockchain()
	cryptoNode.MemoryPool = blockchain.CreateMemoryPool()
	cryptoNode.Wallet = blockchain.CreateWallet()

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

func (cryptoNode *CryptoNode) readSubscriptions() {
	for _, subscription := range cryptoNode.subscriptions {
		go cryptoNode.readSubscription(subscription)
	}
}

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
	blockForgedSub, err := subscribeToTopic(cryptoNode.ctx, gossipPubSub, cryptoNode.Libp2pNode.ID(), BLOCK_FORGED)

	if err != nil {
		panic(err)
	}

	//go readSubscription(transactionSub)
	//go readSubscription(blockChainSub)

	return []Subscription{*transactionSub, *blockForgedSub}
}

func (cryptoNode *CryptoNode) readSubscription(sub Subscription) {
	for message := range sub.Messages {
		utils.Logger.Info(message.Message, sub.TopicName)

		topicType := TopicType(sub.TopicName)

		switch topicType {
		case TRANSACTION:
			utils.Logger.Info("Transaction received from the network")
			var transaction blockchain.Transaction
			transaction = utils.GetStructFromJson(message.Message, transaction).(blockchain.Transaction)
			cryptoNode.handleTransaction(transaction)

		case BLOCK_FORGED:
			utils.Logger.Info("Forged block received from the network")
			var block blockchain.Block
			block = utils.GetStructFromJson(message.Message, block).(blockchain.Block)
			//handleBlock

		}
	}
}

//TODO: getBlockchainMetaData, getMissingBLocks, getEntireBlockchain
//TODO: consensus over the network \w bad actor
//TODO: timing, new forged block transaction not in memory pool then wait a few seconds

func (cryptoNode *CryptoNode) handleBlockForged() {
	//block forged on other node
}

func (cryptoNode *CryptoNode) handleTransaction(transaction blockchain.Transaction) {
	payload := transaction.Payload()
	signature := transaction.Signature
	senderPublicKey := transaction.SenderPublicKey

	transactionInMemoryPool := cryptoNode.MemoryPool.IsTransactionInPool(transaction)
	signatureValid := blockchain.IsValidSignature(payload, signature, senderPublicKey)
	transactionInBlockchain := cryptoNode.Blockchain.IsTransactionInBlockchain(transaction)

	if !transactionInMemoryPool && signatureValid && !transactionInBlockchain {
		cryptoNode.MemoryPool.AddTransaction(transaction)
		utils.Logger.Info("Transaction added to memory pool")
	}
}
