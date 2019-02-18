package votes

import (
	"bytes"
	"context"
	"testing"
	time "time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/app"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/x"
)

const defaultHeight int64 = 10

var helper x.TestHelpers

func randString(len int) string {
	return "abcdefghi"
}

func mustBuildVote(mainVote, repVote, charity, postCode string, birth, donation int32, id string) *VoteRecord {
	res, err := buildVote(mainVote, repVote, charity, postCode, birth, donation, id)
	if err != nil {
		panic(err)
	}
	return res
}

func buildVote(mainVote, repVote, charity, postCode string, birth, donation int32, id string) (*VoteRecord, error) {
	vote := &Vote{
		MainVote:  mainVote,
		RepVote:   repVote,
		Charity:   charity,
		PostCode:  postCode,
		BirthYear: birth,
		Donation:  donation,
	}
	if id == "" {
		id = randString(16)
	}
	now := time.Now().UTC()
	res := &VoteRecord{
		Identitifer:   id,
		SmsCode:       vote.EncodeToSms(),
		TransactionId: randString(20),
		VotedAt:       &now,
		Vote:          vote,
	}
	return res, res.Validate()
}

func TestRecordVoteHandler(t *testing.T) {
	auth := helper.CtxAuth("auth")
	voteBucket := NewVoteBucket()
	// tallyBucket := NewTallyBucket()

	rt := app.NewRouter()
	RegisterRoutes(rt, auth)

	qr := weave.NewQueryRouter()
	RegisterQuery(qr)

	_, src := helper.MakeKey()

	vote1 := mustBuildVote("A", "SOME1", "HLP", "SW87", 1991, 100, "")
	id1 := vote1.Identitifer

	cases := map[string]struct {
		actions []action
		dbtests []querycheck
	}{
		"one vote is tallied": {
			actions: []action{
				{
					conditions: []weave.Condition{src},
					msg:        vote1,
				},
			},
			dbtests: []querycheck{
				{
					path:   "/vote",
					data:   []byte(id1),
					bucket: voteBucket.Bucket,
					wantRes: []orm.Object{
						orm.NewSimpleObj([]byte(id1), vote1),
					},
				},
			},
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			db := store.MemStore()

			for i, a := range tc.actions {
				cache := db.CacheWrap()
				_, err := rt.Check(a.ctx(), cache, a.tx())
				if a.wantCheckErr {
					assert.Error(t, err, i)
					continue
				}
				assert.NoError(t, err, i)
				cache.Discard()

				_, err = rt.Deliver(a.ctx(), db, a.tx())
				if a.wantDeliverErr {
					assert.Error(t, err, i)
					continue
				}
				assert.NoError(t, err, i)
			}
			for _, tt := range tc.dbtests {
				tt.test(t, db, qr)
			}
		})
	}

}

// action represents a single request call that is handled by a handler.
type action struct {
	conditions     []weave.Condition
	msg            weave.Msg
	height         int64
	wantCheckErr   bool
	wantDeliverErr bool
}

func (a *action) tx() weave.Tx {
	return helper.MockTx(a.msg)
}

func (a *action) ctx() weave.Context {
	height := a.height
	if height == 0 {
		height = defaultHeight
	}
	ctx := weave.WithHeight(context.Background(), height)
	ctx = weave.WithChainID(ctx, "testchain-123")
	return helper.CtxAuth("auth").SetConditions(ctx, a.conditions...)
}

// querycheck is a declaration of a query result. For given path and data
// executed within a bucket, ensure that the result is as expected.
// Make sure to register the query router.
type querycheck struct {
	path    string
	data    []byte
	bucket  orm.Bucket
	wantRes []orm.Object
}

// test ensure that querycheck declaration is the same as the database state.
func (qc *querycheck) test(t *testing.T, db weave.ReadOnlyKVStore, qr weave.QueryRouter) {
	t.Helper()

	result, err := qr.Handler(qc.path).Query(db, "", qc.data)
	require.NoError(t, err)
	require.Equal(t, len(qc.wantRes), len(result))
	for i, wres := range qc.wantRes {
		assert.True(t, bytes.HasSuffix(result[i].Key, wres.Key()))

		got, err := qc.bucket.Parse(nil, result[i].Value)
		require.NoError(t, err)
		wvr := wres.Value().(*VoteRecord)
		gvr := got.Value().(*VoteRecord)
		assert.Equal(t, wvr.Vote, gvr.Vote)
		assert.Equal(t, wvr.VotedAt, gvr.VotedAt)
	}
}
