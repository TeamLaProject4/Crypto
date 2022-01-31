package networking

import (
	"context"
	"cryptomunt/blockchain"
	"cryptomunt/proofOfStake"
	"cryptomunt/utils"
	"cryptomunt/wallet"
	"encoding/json"
	"flag"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
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
	//Port           int
	//ProtocolID     string
	//Rendezvous     string
	//Seed           int64
	BootNodes AddrList
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

func CreateAndInitCryptoNode(bootnodes AddrList) CryptoNode {
	utils.Logger.Info("Starting network")
	ctx := context.Background()
	//config := parseFlags()

	//p2p
	node := initHost(ctx, bootnodes)
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
	//cryptoNode.Wallet = wallet.CreateWallet()

	//return cryptoNode
	return cryptoNode
}

func (cryptoNode *CryptoNode) getIpAddrsFromConnectedPeers() ([]string, []peer.ID) {
	peerstore := cryptoNode.Libp2pNode.Peerstore()
	peers := peerstore.PeersWithAddrs()
	peerIpAdresses := make([]string, 1)
	peerIds := make([]peer.ID, 1)
	for _, peer := range peers {
		if peer != cryptoNode.Libp2pNode.ID() {
			peerInfo := peerstore.PeerInfo(peer)
			peerIpAdresses = append(peerIpAdresses, getIpv4AddrFromAddrInfo(peerInfo))
			peerIds = append(peerIds, peer)
		}
	}

	for index, peerIp := range peerIpAdresses {
		if peerIp == "" {
			//remove empty peerIp
			peerIpAdresses = append(peerIpAdresses[:index], peerIpAdresses[index+1:]...)
		}

	}

	utils.Logger.Info("peerIpAdresses", peerIpAdresses)
	return peerIpAdresses, peerIds
}

func getIpv4AddrFromAddrInfo(addrInfo peer.AddrInfo) string {
	for _, addr := range addrInfo.Addrs {
		if strings.Contains(addr.String(), "ip4") && !strings.Contains(addr.String(), "127.0.0") {
			utils.Logger.Info("TEST", addr.String())
			multiAddrIp4 := strings.Split(addr.String(), "/")
			port, _ := strconv.Atoi(multiAddrIp4[4])
			port = port - 1
			return multiAddrIp4[2] + ":" + strconv.Itoa(port)
			//return strings.Split(addr.String(), "/")[2]
		}
	}
	return ""
}

func (cryptoNode *CryptoNode) GetOwnIpAddr() string {
	for _, addr := range cryptoNode.Libp2pNode.Addrs() {
		if strings.Contains(addr.String(), "ip4") && !strings.Contains(addr.String(), "127.0.0") {
			utils.Logger.Info("TEST", addr.String())
			multiAddrIp4 := strings.Split(addr.String(), "/")
			port, _ := strconv.Atoi(multiAddrIp4[4])
			port = port - 1
			return multiAddrIp4[2] + ":" + strconv.Itoa(port)
		}
	}
	return ""
}

func (cryptoNode *CryptoNode) SetBlockchainUsingNetwork() {
	//set blocks
	blocks := cryptoNode.GetAllBlocksFromNetwork()
	cryptoNode.Blockchain.Blocks = blocks

	//TODO: proof of stake? remember stakers?? should it not be removed after stake completed?
	pos := proofOfStake.NewProofOfStake()
	cryptoNode.Blockchain.ProofOfStake = &pos

	//cryptoNode.MemoryPool = creat

	//calculate and set account balances
	am := blockchain.CreateAccountModel()
	cryptoNode.Blockchain.AccountModel = &am
	cryptoNode.Blockchain.AccountModel.SetBalancesFromBlockChain(cryptoNode.Blockchain)
}

//get blockchain blocks from directly connected peers
func (cryptoNode *CryptoNode) GetAllBlocksFromNetwork() []blockchain.Block {
	blocks := *new([]blockchain.Block)
	blocksFromPeersChan := make(chan []blockchain.Block)
	peerIps, peerIds := cryptoNode.getIpAddrsFromConnectedPeers()

	//cryptoNode.Libp2pNode.

	utils.Logger.Info("peerIps", peerIps)

	blockHeight := cryptoNode.getBlockHeightFromPeer(peerIps[0])
	if blockHeight == -1 {
		utils.Logger.Error("No ")
		cryptoNode.Libp2pNode.Network().ClosePeer(peerIds[0])
		//cryptoNode.GetAllBlocksFromNetwork()
		//blockHeight = cryptoNode.getBlockHeightFromPeer(peerIps[1])
	}

	step := blockHeight / len(peerIps)
	step++ //round up to not mis blocks
	start := 0
	end := step

	numOfGoRoutines := 0
	var wg sync.WaitGroup
	for _, peerIp := range peerIps {
		wg.Add(1)
		go func(peerIp string, start int, end int) {
			defer wg.Done()
			go cryptoNode.getBlocksFromPeer(peerIp, start, end, blocksFromPeersChan)
		}(peerIp, start, end)
		numOfGoRoutines++
		start = end //including
		end += step //excluding
	}
	utils.Logger.Info("reached before wg.wait")
	//wg.Done()
	wg.Wait()
	utils.Logger.Info("reached after wg.wait")

	//combine blocks
	blockFromPeersIndex := 0
	for blocksFromPeer := range blocksFromPeersChan {
		utils.Logger.Info("reaced range")
		blocks = append(blocks, blocksFromPeer...)

		if numOfGoRoutines-1 == blockFromPeersIndex {
			close(blocksFromPeersChan)
		}

		blockFromPeersIndex++
	}

	return blocks
}

func (cryptoNode *CryptoNode) getBlocksFromPeer(peerIp string, start int, end int, blocksChan chan []blockchain.Block) {
	response, err := http.Get("http://" + peerIp + "/blockchain/blocks?start=" + strconv.Itoa(start) + "&end=" + strconv.Itoa(end))
	if err != nil {
		utils.Logger.Error("GetBlocksFromNetwork", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			utils.Logger.Warn(err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		utils.Logger.Warn(err)
	}
	blockJson := string(body)
	utils.Logger.Info(blockJson)
	blocksChan <- blockchain.GetBlocksFromJson(blockJson)
}

func (cryptoNode *CryptoNode) getBlockHeightFromPeer(peerIp string) int {
	response, err := http.Get("http://" + peerIp + "/blockchain/block-length")
	if err != nil {
		utils.Logger.Error("GetBlocksFromNetwork", err)
		return -1
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		utils.Logger.Warn(err)
	}

	blockHeightJson := string(body)
	blockHeight, err := strconv.Atoi(blockHeightJson)
	if err != nil {
		utils.Logger.Warn(err)
	}

	return blockHeight
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

//TODO: getBlockchainMetaData, getMissingBLocks, getEntireBlockchain
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
	payload := transaction.Payload()
	signature := transaction.Signature
	senderPublicKey := transaction.SenderPublicKey

	transactionInMemoryPool := cryptoNode.MemoryPool.IsTransactionInPool(transaction)
	signatureValid := wallet.IsValidSignature(payload, signature, senderPublicKey)
	transactionInBlockchain := cryptoNode.Blockchain.IsTransactionInBlockchain(transaction)

	if !transactionInMemoryPool && signatureValid && !transactionInBlockchain {
		cryptoNode.MemoryPool.AddTransaction(transaction)
		utils.Logger.Info("Transaction added to memory pool")
	}
}
