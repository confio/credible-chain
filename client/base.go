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

func NewRemoteClient(url string) (*CredibleClient, error) {
	conn := wc.NewHTTPConnection(url)
	// we must start it to enable websocket subscriptions
	err := conn.Start()
	if err != nil {
		return nil, err
	}
	weave := wc.NewClient(conn)
	return &CredibleClient{weave}, nil
}
