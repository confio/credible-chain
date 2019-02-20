package client

import (
	tmnode "github.com/tendermint/tendermint/node"

	wc "github.com/confio/credible-chain/weaveclient"
)

// CredibleClient provides app-specific queries
type CredibleClient struct {
	*wc.WeaveClient
}

func NewLocalClient(n *tmnode.Node) *CredibleClient {
	conn := wc.NewLocalConnection(n)
	weave := wc.NewClient(conn)
	return &CredibleClient{weave}
}

func NewRemoteClient(url string) *CredibleClient {
	conn := wc.NewHTTPConnection(url)
	weave := wc.NewClient(conn)
	return &CredibleClient{weave}
}
