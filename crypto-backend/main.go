package main

import (
	"cryptomunt/utils"
	"fmt"
	"github.com/libp2p/go-libp2p-core/peer"
	"os"
)

func main() {
	utils.InitLogger()
	CreateNetwork()

	select {}
}

// printErr is like fmt.Printf, but writes to stderr.
func printErr(m string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, m, args...)
}

// shortID returns the last 8 chars of a base58-encoded peer id.
func shortID(p peer.ID) string {
	pretty := p.Pretty()
	return pretty[len(pretty)-8:]
}
