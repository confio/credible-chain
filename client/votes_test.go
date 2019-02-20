package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/x/sigs"

	wc "github.com/confio/credible-chain/weaveclient"
	"github.com/confio/credible-chain/x/votes"
)

func TestVoteTx(t *testing.T) {
	admin := wc.GenPrivateKey()
	identifier := "xxx123456xxx"

	// TODO: make the validation less strict
	vote := &votes.Vote{
		MainVote:  "A",
		RepVote:   "BRDX1",
		Charity:   "ABC",
		PostCode:  "SW16",
		BirthYear: 1980,
		Donation:  100,
	}
	tx := BuildVoteTx(identifier, "BA383SKD 10", "tx id here", vote)
	// if we sign with 0, we can validate against an empty db
	chainID := "ding-dong"
	SignTx(tx, admin, chainID, 0)

	// make sure the tx has a sig
	require.Equal(t, 1, len(tx.GetSignatures()))

	// make sure this validates
	db := store.MemStore()
	conds, err := sigs.VerifyTxSignatures(db, tx, chainID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(conds))
	assert.EqualValues(t, admin.PublicKey().Condition(), conds[0])

	// make sure other chain doesn't validate
	db = store.MemStore()
	_, err = sigs.VerifyTxSignatures(db, tx, "foobar")
	assert.Error(t, err)

	// parse tx and verify we have the proper fields
	data, err := tx.Marshal()
	require.NoError(t, err)
	parsed, err := ParseBcpTx(data)
	require.NoError(t, err)
	msg, err := parsed.GetMsg()
	require.NoError(t, err)
	voteRecord, ok := msg.(*votes.VoteRecord)
	require.True(t, ok)
	err = voteRecord.Validate()
	require.NoError(t, err)

	assert.Equal(t, identifier, voteRecord.Identitifer)
	assert.EqualValues(t, vote, voteRecord.Vote)
}

// func TestWalletQuery(t *testing.T) {
// 	conn := NewLocalConnection(node)
// 	bcp := NewClient(conn)
// 	client.WaitForHeight(conn, 5, fastWaiter)

// 	// bad address returns error
// 	_, err := bcp.GetWallet([]byte{1, 2, 3, 4})
// 	assert.Error(t, err)

// 	// missing account returns nothing
// 	missing := GenPrivateKey().PublicKey().Address()
// 	wallet, err := bcp.GetWallet(missing)
// 	assert.NoError(t, err)
// 	assert.Nil(t, wallet)

// 	// genesis account returns something
// 	address := faucet.PublicKey().Address()
// 	wallet, err = bcp.GetWallet(address)
// 	assert.NoError(t, err)
// 	require.NotNil(t, wallet)
// 	// make sure we get some reasonable height
// 	assert.True(t, wallet.Height > 4)
// 	// ensure the key matches
// 	assert.EqualValues(t, address, wallet.Address)
// 	// check the wallet
// 	require.Equal(t, 1, len(wallet.Wallet.Coins))
// 	coin := wallet.Wallet.Coins[0]
// 	assert.Equal(t, initBalance.Whole, coin.Whole)
// 	assert.Equal(t, initBalance.Ticker, coin.Ticker)
// }

// func TestSendMoney(t *testing.T) {
// 	conn := NewLocalConnection(node)
// 	bcp := NewClient(conn)

// 	rcpt := GenPrivateKey().PublicKey().Address()
// 	src := faucet.PublicKey().Address()

// 	nonce := NewNonce(bcp, src)
// 	chainID := getChainID()

// 	// build the tx
// 	amount := x.Coin{Whole: 1000, Ticker: initBalance.Ticker}
// 	tx := BuildSendTx(src, rcpt, amount, "Send 1")
// 	n, err := nonce.Query()
// 	require.NoError(t, err)
// 	SignTx(tx, faucet, chainID, n)

// 	// now post it
// 	res := bcp.BroadcastTx(tx)
// 	require.NoError(t, res.IsError())

// 	// verify nonce incremented on chain
// 	n2, err := nonce.Query()
// 	require.NoError(t, err)
// 	assert.Equal(t, n+1, n2)

// 	// verify wallet has cash
// 	wallet, err := bcp.GetWallet(rcpt)
// 	assert.NoError(t, err)
// 	require.NotNil(t, wallet)
// 	// check the wallet
// 	require.Equal(t, 1, len(wallet.Wallet.Coins))
// 	coin := wallet.Wallet.Coins[0]
// 	assert.Equal(t, int64(1000), coin.Whole)
// 	assert.Equal(t, initBalance.Ticker, coin.Ticker)
// }
