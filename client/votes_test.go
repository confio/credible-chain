package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/rpc/client"

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
		MainVote:  1,
		RepVote:   "BR1",
		Charity:   "AB",
		PostCode:  "SW16",
		BirthYear: 1980,
		Donation:  100,
	}
	tx, err := BuildVoteTx(identifier, "BA383SKD 10", vote)
	require.NoError(t, err)
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

	assert.Equal(t, identifier, voteRecord.Identifier)
	assert.EqualValues(t, vote, voteRecord.Vote)
}

func TestOneVote(t *testing.T) {
	cc := NewLocalClient(node)
	client.WaitForHeight(cc, 5, wc.FastWaiter)
	chainID, err := cc.ChainID()
	require.NoError(t, err)

	// let's submit one vote
	var main int32 = 3
	rep := "FOO"

	// assert both are empty now
	mainTally := votes.VoteToString(main)
	imain, err := cc.GetTally(mainTally)
	assert.NoError(t, err)
	assert.Nil(t, imain)
	irep, err := cc.GetTally(rep)
	assert.NoError(t, err)
	assert.Nil(t, irep)

	vote := &votes.Vote{
		MainVote:  main,
		RepVote:   rep,
		Charity:   "AL",
		PostCode:  "NE11",
		BirthYear: 1987,
		Donation:  200,
	}
	identifier := "abcdef12345"

	// prepare step
	tx, err := BuildVoteTx(identifier, "BA383SKD 10", vote)
	require.NoError(t, err)
	nonce := wc.NewNonce(cc, notary.PublicKey().Address())
	n, err := nonce.Next()
	require.NoError(t, err)
	SignTx(tx, notary, chainID, n)

	// send step
	res := cc.BroadcastTx(tx)
	require.NoError(t, res.IsError())

	// verify nonce incremented on chain
	n2, err := nonce.Query()
	require.NoError(t, err)
	assert.Equal(t, n2, n+1)

	// let's go check the tally was updated
	fmain, err := cc.GetTally(mainTally)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, fmain.Total)
	assert.Equal(t, mainTally, fmain.Option)
	frep, err := cc.GetTally(rep)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, frep.Total)
	assert.Equal(t, rep, frep.Option)

	// and we should be able to find the vote by id
	voted, err := cc.GetVote(identifier)
	require.NoError(t, err)
	assert.EqualValues(t, vote, voted.Vote)
}
