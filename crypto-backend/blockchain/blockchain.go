package main

type Blockchain struct {
	blocks       []Block
	accountModel AccountModel
	proofOfStake interface{}
}

func NewBlockchain() (self *Blockchain) {
	self = new(Blockchain)
	self.blocks = []interface{}{Block.genesis()}
	self.accountModel = AccountModel()
	self.pos = ProofOfStake()
	return
}

func (self *Blockchain) to_json() map[interface{}]interface{} {
	json_dict := map[string][]interface{}{}
	json_dict["blocks"] = func() (elts []interface{}) {
		for _, block := range self.blocks {
			elts = append(elts, block.to_json())
		}
		return
	}()
	return json_dict
}

func (self *Blockchain) latest_block_height() int {
	return self.blocks[len(self.blocks)-1].height
}

func (self *Blockchain) add_block(block Block) {
	if self.is_valid_block_count(block) && self.is_valid_previous_block_hash(block) {
		self.execute_transactions(block.transactions)
		self.blocks = append(self.blocks, block)
	}
}

func (self *Blockchain) execute_transactions(transactions []Transaction) {
	for _, tx := range transactions {
		self.execute_transaction(tx)
	}
}

func (self *Blockchain) execute_transaction(transaction Transaction) {
	if transaction.tx_type == TxType.STAKE {
		sender := transaction.sender_public_key
		receiver := transaction.receiver_public_key
		if sender == receiver {
			amount := transaction.amount
			self.pos.update_stake(sender, amount)
			self.accountModel.update_balance(sender, -amount)
		}
	} else {
		sender := transaction.sender_public_key
		receiver := transaction.receiver_public_key
		amount := transaction.amount
		self.accountModel.update_balance(sender, -amount)
		self.accountModel.update_balance(receiver, amount)
	}
}

func (self *Blockchain) is_valid_block_count(block Block) bool {
	return self.blocks[len(self.blocks)-1].height == block.height-1
}

func (self *Blockchain) is_valid_previous_block_hash(block Block) bool {
	return self.latest_previous_hash() == block.previous_hash
}

func (self *Blockchain) is_valid_forger(block Block) bool {
	return block.forger == self.next_forger()
}

func (self *Blockchain) is_block_transactions_valid(block Block) bool {
	transactions := block.transactions
	covered_transactions := self.get_covered_transactions(transactions)
	return len(transactions) == len(covered_transactions)
}

func (self *Blockchain) latest_previous_hash() string {
	return BlockchainUtils.hash(self.blocks[len(self.blocks)-1].payload()).hexdigest()
}

func (self *Blockchain) get_covered_transactions(transactions []Transaction) []Transaction {
	return func() (elts []Transaction) {
		for _, tx := range transactions {
			if self.is_transaction_covered(tx) {
				elts = append(elts, tx)
			}
		}
		return
	}()
}

func (self *Blockchain) is_transaction_covered(transaction Transaction) bool {
	if transaction.tx_type == TxType.EXCHANGE {
		return true
	}
	sender_balance := self.accountModel.get_balance(transaction.sender_public_key)
	return sender_balance >= transaction.amount
}

func (self *Blockchain) get_account_balance(public_key_string string) int {
	return self.accountModel.get_balance(public_key_string)
}

func (self *Blockchain) next_forger() string {
	prev_block_hash := self.latest_previous_hash()
	return self.pos.pick_forger(prev_block_hash)
}

func (self *Blockchain) create_block(transactions []Transaction, forger_wallet interface{}) Block {
	covered_transactions := self.get_covered_transactions(transactions)
	self.execute_transactions(covered_transactions)
	block := forger_wallet.create_block(
		covered_transactions,
		self.latest_previous_hash(),
		self.latest_block_height()+1,
	)
	self.add_block(block)
	return block
}

func (self *Blockchain) is_transaction_in_blockchain(transaction Transaction) bool {
	for _, block := range self.blocks {
		for _, block_tx := range block.transactions {
			if block_tx == transaction {
				return true
			}
		}
	}
	return false
}
