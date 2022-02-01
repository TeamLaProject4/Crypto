package structs

type ApiCallType int64

type ApiCallMessage struct {
	CallType ApiCallType
	Message  string
}

const (
	GET_BLOCKS ApiCallType = iota
	GET_BLOCK_LENGTH
	GET_MEMORY_POOL
	GET_BALANCE
	GET_TRANSACTIONS
	GET_MNEMONIC
	CONFIRM_MNEMONIC
	POST_TRANSACTION
)

type ParamsGetBlocks struct {
	Start string
	End   string
}
