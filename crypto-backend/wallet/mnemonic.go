package wallet

import (
	"github.com/tyler-smith/go-bip39"
)

func GenerateMnemonic() string {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	return mnemonic
}
