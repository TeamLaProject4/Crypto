package networking

import (
	"cryptomunt/blockchain"
	"cryptomunt/utils"
	"encoding/json"
	"github.com/libp2p/go-libp2p-core/peer"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func (cryptoNode *CryptoNode) getMemoryPoolFromPeer(peerIp string) blockchain.MemoryPool {
	response, err := http.Get("http://" + peerIp + "/blockchain/memory-pool")
	if err != nil {
		utils.Logger.Error(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		utils.Logger.Warn(err)
	}
	memPoolJson := string(body)

	var memPool blockchain.MemoryPool
	err = json.Unmarshal([]byte(memPoolJson), &memPool)
	if err != nil {
		utils.Logger.Error("unmarshal error ", err)
	}
	return memPool
}

//get blockchain blocks from directly connected peers
func (cryptoNode *CryptoNode) GetAllBlocksFromNetwork() []blockchain.Block {
	blocks := *new([]blockchain.Block)
	blocksFromPeersChan := make(chan []blockchain.Block)
	peerIps, peerIds := cryptoNode.getIpAddrsFromConnectedPeers()

	blockHeight := cryptoNode.getBlockHeightFromPeer(peerIps[0])
	if blockHeight == -1 {
		utils.Logger.Error("No ")
		cryptoNode.Libp2pNode.Network().ClosePeer(peerIds[0])
	}

	numOfGoRoutines := cryptoNode.getBlocksFromPeers(blockHeight, peerIps, blocksFromPeersChan)
	blocks = cryptoNode.combineBlocksFromPeers(blocksFromPeersChan, blocks, numOfGoRoutines)

	return blocks
}

func (cryptoNode *CryptoNode) combineBlocksFromPeers(blocksFromPeersChan chan []blockchain.Block, blocks []blockchain.Block, numOfGoRoutines int) []blockchain.Block {
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

func (cryptoNode *CryptoNode) getBlocksFromPeers(blockHeight int, peerIps []string, blocksFromPeersChan chan []blockchain.Block) int {
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
	wg.Wait()
	return numOfGoRoutines
}

func (cryptoNode *CryptoNode) getBlocksFromPeer(peerIp string, start int, end int, blocksChan chan []blockchain.Block) {
	response, err := http.Get("http://" + peerIp + "/blockchain/blocks?start=" + strconv.Itoa(start) + "&end=" + strconv.Itoa(end))
	if err != nil {
		utils.Logger.Error("GetBlocksFromNetwork", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		utils.Logger.Warn(err)
	}
	blockJson := string(body)
	//utils.Logger.Info(blockJson)
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
			multiAddrIp4 := strings.Split(addr.String(), "/")
			port, _ := strconv.Atoi(multiAddrIp4[4])
			port = port - 1
			return multiAddrIp4[2] + ":" + strconv.Itoa(port)
		}
	}
	return ""
}
