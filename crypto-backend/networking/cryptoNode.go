package networking

import (
	"context"
	"cryptomunt/blockchain"
	"cryptomunt/proofOfStake"
	"cryptomunt/structs"
	"cryptomunt/utils"
	"cryptomunt/wallet"
	"encoding/json"
	"github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const RANDEVOUS_STRING = "cryptomunt-randevous"

type TopicType string

const (
	TRANSACTION     TopicType = "TRANSACTION" //transaction
	BLOCK_FORGED              = "BLOCK_FORGED"
	CONSENSUS_ERROR           = "CONSENSUS_ERROR"
	NODE_
)

type Config struct {
	NodesToBoot int
	BootNodes   AddrList
}

type CryptoNode struct {
	Libp2pNode    host.Host
	ctx           context.Context
	subscriptions []Subscription
	Blockchain    blockchain.Blockchain
	MemoryPool    blockchain.MemoryPool
	Wallet        wallet.Wallet
	//sub map[TopicType]Subscription
}

func CreateAndInitCryptoNode(config Config, apiRequest chan structs.ApiCallMessage, apiResponse chan structs.ApiCallMessage) *CryptoNode {
	utils.Logger.Info("Starting network")
	ctx := context.Background()

	node := initHost(ctx, config.BootNodes)
	//init
	cryptoNode := CryptoNode{
		Libp2pNode: node,
		ctx:        ctx,
	}

	//pubsub
	initPubSub(&cryptoNode)
	//blockchain, memPool, wallet
	initBlockchain(config, &cryptoNode)

	return &cryptoNode
}

func (cryptoNode *CryptoNode) HandleApiCalls(apiRequest chan structs.ApiCallMessage, apiResponse chan structs.ApiCallMessage) {
	for requestMessage := range apiRequest {
		utils.Logger.Info("GOTTEN API CALL: ", requestMessage.Message, requestMessage.CallType)

		switch requestMessage.CallType {
		case structs.GET_BLOCKS:
			cryptoNode.handleGetBlocks(requestMessage.Message, apiResponse)

		}

	}
}

func initBlockchain(config Config, cryptoNode *CryptoNode) {
	//if bootnode then initialise, else set using network
	if len(config.BootNodes) > 0 {
		cryptoNode.SetBlockchainUsingNetwork()
	} else {
		cryptoNode.Blockchain = blockchain.CreateBlockchain()
		cryptoNode.MemoryPool = blockchain.CreateMemoryPool()
		cryptoNode.Wallet = wallet.CreateWalletFromKeyFile()
	}
}

func initPubSub(cryptoNode *CryptoNode) {
	cryptoNode.logNodeAddr()
	cryptoNode.subscriptions = cryptoNode.subscribeToTopics()
	cryptoNode.readSubscriptions()
}

func (cryptoNode *CryptoNode) SetBlockchainUsingNetwork() {
	//set blocks
	blocks := cryptoNode.GetAllBlocksFromNetwork()
	cryptoNode.Blockchain.Blocks = blocks

	pos := proofOfStake.NewProofOfStake()
	cryptoNode.Blockchain.ProofOfStake = &pos

	//set memorypool
	peerIps, _ := cryptoNode.getIpAddrsFromConnectedPeers()
	cryptoNode.MemoryPool = cryptoNode.getMemoryPoolFromPeer(peerIps[0])

	//calculate and set account balances
	am := blockchain.CreateAccountModel()
	cryptoNode.Blockchain.AccountModel = &am
	cryptoNode.Blockchain.AccountModel.SetBalancesFromBlockChain(cryptoNode.Blockchain)
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
			err := json.Unmarshal([]byte(message.Message), &transaction)
			if err != nil {
				utils.Logger.Error("unmarshal error ", err)
			}
			cryptoNode.handleTransaction(transaction)

		case BLOCK_FORGED:
			utils.Logger.Info("Forged block received from the network")
			//var block blockchain.Block
			//block = utils.GetStructFromJson(message.Message, block).(blockchain.Block)
			//handleBlock

		}
	}
}

//TODO: getMissingBLocks
//TODO: consensus over the network \w bad actor
//TODO: timing, new forged block transaction not in memory pool then wait a few seconds

//block forged on other node
func (cryptoNode *CryptoNode) handleBlockForged(block blockchain.Block) {

	//TODO: cryptoNode.request_missing_blocks()
	//if !blockCountValid {
	//	cryptoNode.request_missing_blocks()
	//}
	if cryptoNode.isForgedBlockValid(block) {
		cryptoNode.Blockchain.AddBlock(block)
		cryptoNode.MemoryPool.RemoveTransactions(block.Transactions)
	}
}

func (cryptoNode *CryptoNode) isForgedBlockValid(block blockchain.Block) bool {
	payload := block.Payload()
	signature := block.Signature
	forgerPublicKey := block.Forger

	blockCountValid := cryptoNode.Blockchain.IsValidBlockHeight(block)
	previousBlockHashValid := cryptoNode.Blockchain.IsValidPreviousBlockHash(block)
	signatureValid := wallet.IsValidSignature(payload, signature, forgerPublicKey)
	forgerValid := cryptoNode.Blockchain.IsValidForger(block)
	blockTransactionsValid := cryptoNode.Blockchain.IsBlockTransactionsValid(block)

	return blockTransactionsValid && forgerValid && signatureValid && previousBlockHashValid && blockCountValid
}

func (cryptoNode *CryptoNode) handleTransaction(transaction blockchain.Transaction) {
	if cryptoNode.IsTransactionValid(transaction) {
		cryptoNode.MemoryPool.AddTransaction(transaction)
		utils.Logger.Info("Transaction added to memory pool")
	}
}

func (cryptoNode *CryptoNode) IsTransactionValid(transaction blockchain.Transaction) bool {
	payload := transaction.Payload()
	signature := transaction.Signature
	senderPublicKey := transaction.SenderPublicKeyString

	transactionInMemoryPool := cryptoNode.MemoryPool.IsTransactionInPool(transaction)
	signatureValid := wallet.IsValidSignature(payload, signature, senderPublicKey)
	transactionInBlockchain := cryptoNode.Blockchain.IsTransactionInBlockchain(transaction)
	balanceNegative := cryptoNode.balanceNegativeAfterTransaction(transaction)

	return !transactionInMemoryPool && signatureValid && !transactionInBlockchain && !balanceNegative
}

func (cryptoNode *CryptoNode) balanceNegativeAfterTransaction(transaction blockchain.Transaction) bool {
	balance := cryptoNode.Blockchain.AccountModel.GetBalance(cryptoNode.Wallet.GetPublicKeyString())
	return transaction.Amount > balance
}
