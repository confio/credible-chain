package client

import (
	"testing"

	"github.com/confio/credible-chain/x/votes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/x/sigs"
)

func TestVoteTx(t *testing.T) {
	admin := GenPrivateKey()
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
