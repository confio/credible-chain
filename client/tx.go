package client

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/crypto"
	"github.com/iov-one/weave/x/sigs"

	app "github.com/confio/credible-chain/app"
)

// Tx is all the interfaces we need rolled into one
type Tx interface {
	weave.Tx
	sigs.SignedTx
}

// SignTx modifies the tx in-place, adding signatures
func SignTx(tx *app.Tx, signer *crypto.PrivateKey, chainID string, nonce int64) error {
	sig, err := sigs.SignTx(signer, tx, chainID, nonce)
	if err != nil {
		return err
	}
	tx.Signatures = append(tx.Signatures, sig)
	return nil
}

// ParseBcpTx will load a serialize tx into a format we can read
func ParseBcpTx(data []byte) (*app.Tx, error) {
	var tx app.Tx
	err := tx.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}
