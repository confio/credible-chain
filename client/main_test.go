package client

import (
	"fmt"
	"os"
	"testing"
	"time"

	abci "github.com/tendermint/tendermint/abci/types"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"
	nm "github.com/tendermint/tendermint/node"
	rpctest "github.com/tendermint/tendermint/rpc/test"
	tm "github.com/tendermint/tendermint/types"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/crypto"

	credchain "github.com/confio/credible-chain/app"
	wc "github.com/confio/credible-chain/weaveclient"
)

// adjust this to get debug output
var logger = log.NewNopLogger() // log.NewTMLogger()

// useful values for test cases
var node *nm.Node
var notary *crypto.PrivateKey

func getChainID() string {
	return rpctest.GetConfig().ChainID()
}

func TestMain(m *testing.M) {
	notary = wc.GenPrivateKey()

	// TODO: check out config file...
	config := rpctest.GetConfig()
	config.Moniker = "SetInTestMain"

	// set up our application
	admin := notary.PublicKey().Address()
	app, err := initApp(config, admin)
	if err != nil {
		panic(err) // what else to do???
	}

	// run the app inside a tendermint instance
	node = rpctest.StartTendermint(app)
	time.Sleep(100 * time.Millisecond) // time to setup app context
	code := m.Run()

	// and shut down proper at the end
	node.Stop()
	node.Wait()
	os.Exit(code)
}

func initApp(config *cfg.Config, addr weave.Address) (abci.Application, error) {
	bcp, err := credchain.GenerateApp(config.RootDir, logger, true)
	if err != nil {
		return nil, err
	}

	// generate genesis file...
	err = initGenesis(config.GenesisFile(), addr)
	return bcp, err
}

func initGenesis(filename string, addr weave.Address) error {
	doc, err := tm.GenesisDocFromFile(filename)
	if err != nil {
		return err
	}
	appState, err := credchain.GenInitOptions([]string{addr.String()})
	if err != nil {
		return fmt.Errorf("serialize state: %s", err)
	}
	doc.AppState = appState
	return doc.SaveAs(filename)
}
