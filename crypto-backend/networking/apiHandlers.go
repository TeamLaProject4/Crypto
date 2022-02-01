package networking

import (
	"cryptomunt/structs"
	"cryptomunt/utils"
	"encoding/json"
	"strconv"
)

func (cryptoNode *CryptoNode) handleGetBlocks(message string, apiResponse chan structs.ApiCallMessage) {
	var params structs.ParamsGetBlocks
	err := json.Unmarshal([]byte(message), &params)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	startInt, _ := strconv.Atoi(params.Start)
	endInt, _ := strconv.Atoi(params.End)

	blocks := cryptoNode.Blockchain.GetBlocksFromRange(startInt, endInt)
	blocksJson, err := json.Marshal(blocks)
	if err != nil {
		utils.Logger.Error(err)
		return

	}
	apiResponse <- structs.ApiCallMessage{
		CallType: structs.GET_BLOCKS,
		Message:  string(blocksJson),
	}

}
func handleGetBlockLength() {

}
func handleGetBalance() {}

func handleGetTransactions() {}
func handleGetMnemonic()     {}
func handleConfirmMnemonic() {}
func handlePostTransaction() {}
