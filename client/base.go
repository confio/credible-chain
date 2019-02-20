package client

import (
	wc "github.com/confio/credible-chain/weaveclient"
)

// CredibleClient provides app-specific queries
type CredibleClient struct {
	wc.WeaveClient
}
