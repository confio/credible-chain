package credchain

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"path/filepath"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/app"
	"github.com/iov-one/weave/gconf"
	"github.com/iov-one/weave/x/multisig"
)

// GenInitOptions will produce empty structure to fill in
func GenInitOptions(args []string) (json.RawMessage, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Usage: init <notary address hex>")
	}
	bz, err := hex.DecodeString(args[0])
	if err != nil {
		return nil, err
	}
	err = weave.Address(bz).Validate()
	if err != nil {
		return nil, err
	}
	opts := fmt.Sprintf(`
          {
			"multisig": [],
			"gconf": {
				"votes:notary": "%s"
			}	
          }
	`, args[0])
	return []byte(opts), nil
}

// GenerateApp is used to create a stub for server/start.go command
func GenerateApp(home string, logger log.Logger, debug bool) (abci.Application, error) {
	// db goes in a subdir, but "" stays "" to use memdb
	var dbPath string
	if home != "" {
		dbPath = filepath.Join(home, "credchain.db")
	}
	stack := Stack()
	application, err := Application("credchain", stack, TxDecoder, dbPath, debug)
	if err != nil {
		return nil, err
	}
	return DecorateApp(application, logger), nil
}

// DecorateApp adds initializers and Logger to an Application
func DecorateApp(application app.BaseApp, logger log.Logger) app.BaseApp {
	application.WithInit(app.ChainInitializers(
		&gconf.Initializer{},
		&multisig.Initializer{},
	))
	application.WithLogger(logger)
	return application
}

// InlineApp will take a previously prepared CommitStore and return a complete Application
func InlineApp(kv weave.CommitKVStore, logger log.Logger, debug bool) abci.Application {
	stack := Stack()
	ctx := context.Background()
	store := app.NewStoreApp("credchain", kv, QueryRouter(), ctx)
	base := app.NewBaseApp(store, TxDecoder, stack, nil, debug)
	return DecorateApp(base, logger)
}
