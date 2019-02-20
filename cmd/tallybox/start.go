package main

import (
	"flag"
	"fmt"
	"strings"

	client "github.com/confio/credible-chain/client"
	wc "github.com/confio/credible-chain/weaveclient"
)

const (
	flagPort   = "port"
	flagRemote = "remote"
)

func parseStartFlags(args []string) (string, int, error) {
	// parse flagBind and return the result
	var remote string
	var port int
	startFlags := flag.NewFlagSet("start", flag.ExitOnError)
	startFlags.IntVar(&port, flagPort, 5005, "address server listens on")
	startFlags.StringVar(&remote, flagRemote, "", "location of blockchain node, ex. http://localhost:26657")
	err := startFlags.Parse(args)
	return remote, port, err
}

// StartCmd runs the api server on given port.
// Accepts votes on POST /vote
// Shows tally on GET /tally
func StartCmd(home string, args []string) error {
	remote, _, err := parseStartFlags(args)
	if err != nil {
		return err
	}
	if remote == "" || !(strings.HasPrefix(remote, "http://") || strings.HasPrefix(remote, "https://")) {
		return fmt.Errorf("Usage: tallybox start -remote=<url> [-port=<port>]")
	}

	// load keys
	filename := notaryPath(home)
	// privkey, err := wc.LoadPrivateKey(filename)
	_, err = wc.LoadPrivateKey(filename)
	if err != nil {
		return fmt.Errorf("Cannot load keys from %s", filename)
	}

	// connect to remote chain
	cc := client.NewRemoteClient(remote)
	h, err := cc.Height()
	if err != nil {
		return err
	}
	fmt.Printf("Height: %d\n", h)
	chain, err := cc.ChainID()
	if err != nil {
		return err
	}
	fmt.Printf("ChainId: %s\n", chain)

	// start server
	return nil
}
