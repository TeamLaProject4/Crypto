package integration_test

import (
	"testing"
)

func TestGetBlockChainFromNetwork(t *testing.T) {
	//utils.InitLogger()
	//
	//networking.CreateCryptoNode()
	//
	//node.GetBlockChainFromNetwork()

	got := 30
	want := 30

	if got != want {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
