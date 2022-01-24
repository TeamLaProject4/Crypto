package proofOfStake

type ProofOfStake struct {
	stakers map[string]int
}

var proofOfStake = new(ProofOfStake)

func NewProofOfStake() {
	//self.set_genesis_node_stake()
}

// TODO: unit test
//TODO: remove hardcoded path and add more generic solution
// TODO: after others stake, then genesis can no longer be forger, if everyone unstaked, then genesis can be forger
//func set_genesis_node_stake() {
//	func() {
//		f := func() *os.File {
//			f, err := os.OpenFile("keys/genesisPublicKey.pem", os.O_RDONLY, 0o777)
//			if err != nil {
//				panic(err)
//			}
//			return f
//		}()
//		defer func() {
//			if err := f.Close(); err != nil {
//				panic(err)
//			}
//		}()
//		genesisPublicKey := func() string {
//			content, err := ioutil.ReadAll(f)
//			if err != nil {
//				panic(err)
//			}
//			return string(content)
//		}()
//		proofOfStake.stakers[genesisPublicKey] = 1
//	}()
//}
//
//func isAccountInStakers(publicKey string) bool {
//	_, accountInStakers := proofOfStake.stakers[publicKey]
//	return accountInStakers
//}
//
//func addAccountToStakers(publicKey string) {
//	if !isAccountInStakers(publicKey) {
//		proofOfStake.stakers[publicKey] = 0
//	}
//}
//
//func updateStake(publicKey string, stake int) {
//	addAccountToStakers(publicKey)
//	proofOfStake.stakers[publicKey] += stake
//}
//
//func getStake(publicKey string) int {
//	//TODO: addAccount to stakers not logical here, remove?
//	addAccountToStakers(publicKey)
//	return proofOfStake.stakers[publicKey]
//}
//
//func generateLots(seed string) []Lot {
//	generateStakerLots := func(publicKeyString string) []Lot {
//		stakeAmount := getStake(publicKeyString)
//		return func() (elts []interface{}) {
//			for stake := 0; stake < stakeAmount; stake++ {
//				elts = append(elts, Lot(publicKeyString, stake+1, seed))
//			}
//			return
//		}()
//	}
//	lots := []interface{}{}
//	for _, staker := range self.stakers {
//		lots = append(lots, generateStakerLots(staker)...)
//	}
//	return lots
//}
//
//func (self *ProofOfStake) __pick_winner(lots []Lot, seed string) Lot {
//	is_offset_smaller := func(lot Lot) bool {
//		offset := get_offset(lot)
//		return offset < least_offset
//	}
//	get_offset := func(lot Lot) int {
//		lot_int := int(lot.hash(), 16)
//		return math.Abs(lot_int - seed_int)
//	}
//	winner := lots[0]
//	least_offset := func() float64 {
//		i, err := strconv.ParseFloat("INF", 64)
//		if err != nil {
//			panic(err)
//		}
//		return i
//	}()
//	seed_int := int(utils.BlockchainUtils.hash(seed).hexdigest(), 16)
//	for _, lot := range lots {
//		if is_offset_smaller(lot) {
//			least_offset = get_offset(lot)
//			winner = lot
//		}
//	}
//	return winner
//}
//
//func (self *ProofOfStake) pick_forger(previous_block_hash string) string {
//	lots := self.__generate_lots(previous_block_hash)
//	winner := self.__pick_winner(lots, previous_block_hash)
//	return winner.publicKey
//}
