package votes

import (
	"bytes"
	"context"
	"encoding/hex"
	"testing"
	time "time"

	"github.com/iov-one/weave/gconf"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/libs/common"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/app"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/x"
)

const defaultHeight int64 = 10

var helper x.TestHelpers

func randString(len int) string {
	return common.RandStr(len)
}

func mustBuildVote(mainVote int32, repVote, charity, postCode string, birth, donation int32, id string) *VoteRecord {
	res := buildVote(mainVote, repVote, charity, postCode, birth, donation, id)
	if err := res.Validate(); err != nil {
		panic(err)
	}
	return res
}

func buildVote(mainVote int32, repVote, charity, postCode string, birth, donation int32, id string) *VoteRecord {
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
	return res
}

func buildTally(option string, total int64) orm.Object {
	return orm.NewSimpleObj([]byte(option), &Tally{
		Option: option,
		Total:  total,
	})
}

func TestRecordVoteHandler(t *testing.T) {
	auth := helper.CtxAuth("auth")
	voteBucket := NewVoteBucket()
	tallyBucket := NewTallyBucket()

	rt := app.NewRouter()
	RegisterRoutes(rt, auth)

	qr := weave.NewQueryRouter()
	RegisterQuery(qr)

	_, us := helper.MakeKey()
	_, them := helper.MakeKey()

	rep1 := "SOM"
	rep2 := "ELS"

	tallyOpt := func(i int32) []byte { return []byte(VoteToString(i)) }

	// this is for over-writing
	vote1 := mustBuildVote(1, rep1, "HP", "SW87", 1991, 100, "")
	id1 := vote1.Identitifer
	vote1b := mustBuildVote(2, rep2, "F7", "SW87", 1991, 100, id1)

	// this is for addition
	vote2 := mustBuildVote(1, rep2, "MR", "B18", 1980, 500, "")
	id2 := vote2.Identitifer

	cases := map[string]struct {
		actions []action
		dbtests []querycheck
	}{
		"requires valid signature": {
			actions: []action{
				{
					conditions:     []weave.Condition{them},
					msg:            vote1,
					wantCheckErr:   true,
					wantDeliverErr: true,
				},
			},
		},
		"rejects invalid ": {
			actions: []action{
				{
					conditions:     []weave.Condition{us},
					msg:            buildVote(5, rep1, "HLP", "SW87", 1991, 100, ""),
					wantCheckErr:   true,
					wantDeliverErr: true,
				},
				{
					conditions:     []weave.Condition{us},
					msg:            buildVote(1, "NOONE", "HLP", "SW87", 1991, 100, ""),
					wantCheckErr:   true,
					wantDeliverErr: true,
				},
				{
					conditions:     []weave.Condition{us},
					msg:            buildVote(2, rep1, "HLP", "SW87", 91, 100, ""),
					wantCheckErr:   true,
					wantDeliverErr: true,
				},
				{
					conditions:     []weave.Condition{us},
					msg:            buildVote(3, rep1, "HLP", "SW87", 2010, 100, ""),
					wantCheckErr:   true,
					wantDeliverErr: true,
				},
				{
					conditions:     []weave.Condition{us},
					msg:            buildVote(1, rep1, "HLP", "SW87 K23", 1986, 100, ""),
					wantCheckErr:   true,
					wantDeliverErr: true,
				},
			},
		},
		"one vote is tallied": {
			actions: []action{
				{
					conditions: []weave.Condition{us},
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
				{
					path:   "/tally",
					data:   tallyOpt(1),
					bucket: tallyBucket.Bucket,
					wantRes: []orm.Object{
						buildTally(VoteToString(1), 1),
					},
				},
				{
					path:   "/tally",
					data:   []byte(rep1),
					bucket: tallyBucket.Bucket,
					wantRes: []orm.Object{
						buildTally(rep1, 1),
					},
				},
				{
					path:   "/tally",
					mod:    "prefix",
					bucket: tallyBucket.Bucket,
					wantRes: []orm.Object{
						buildTally(rep1, 1),
						buildTally(VoteToString(1), 1),
					},
				},
			},
		},
		"can update vote": {
			actions: []action{
				{
					conditions: []weave.Condition{us},
					msg:        vote1,
				},
				{
					conditions: []weave.Condition{us},
					msg:        vote1b,
				},
			},
			dbtests: []querycheck{
				{
					path:   "/vote",
					data:   []byte(id1),
					bucket: voteBucket.Bucket,
					wantRes: []orm.Object{
						orm.NewSimpleObj([]byte(id1), vote1b),
					},
				},
				{
					path:   "/tally",
					data:   tallyOpt(2),
					bucket: tallyBucket.Bucket,
					wantRes: []orm.Object{
						buildTally(VoteToString(2), 1),
					},
				},
				{
					path:   "/tally",
					data:   []byte(rep2),
					bucket: tallyBucket.Bucket,
					wantRes: []orm.Object{
						buildTally(rep2, 1),
					},
				},
				{
					path:   "/tally",
					mod:    "prefix",
					bucket: tallyBucket.Bucket,
					wantRes: []orm.Object{
						buildTally(rep2, 1),
						buildTally(rep1, 0),
						buildTally(VoteToString(1), 0),
						buildTally(VoteToString(2), 1),
					},
				},
			},
		},
		"can combine vote": {
			actions: []action{
				{
					conditions: []weave.Condition{us},
					msg:        vote1,
				},
				{
					conditions: []weave.Condition{us},
					msg:        vote2,
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
				{
					path:   "/vote",
					data:   []byte(id2),
					bucket: voteBucket.Bucket,
					wantRes: []orm.Object{
						orm.NewSimpleObj([]byte(id2), vote2),
					},
				},
				{
					path:   "/tally",
					data:   tallyOpt(1),
					bucket: tallyBucket.Bucket,
					wantRes: []orm.Object{
						buildTally(VoteToString(1), 2),
					},
				},
				{
					path:   "/tally",
					mod:    "prefix",
					bucket: tallyBucket.Bucket,
					wantRes: []orm.Object{
						buildTally(rep2, 1),
						buildTally(rep1, 1),
						buildTally(VoteToString(1), 2),
					},
				},
			},
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			db := store.MemStore()
			gconf.SetValue(db, gconfNotary, hex.EncodeToString(us.Address()))

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
	mod     string
	data    []byte
	bucket  orm.Bucket
	wantRes []orm.Object
}

// test ensure that querycheck declaration is the same as the database state.
func (qc *querycheck) test(t *testing.T, db weave.ReadOnlyKVStore, qr weave.QueryRouter) {
	t.Helper()

	result, err := qr.Handler(qc.path).Query(db, qc.mod, qc.data)
	require.NoError(t, err)
	require.Equal(t, len(qc.wantRes), len(result))
	for i, wres := range qc.wantRes {
		assert.True(t, bytes.HasSuffix(result[i].Key, wres.Key()))

		got, err := qc.bucket.Parse(wres.Key(), result[i].Value)
		require.NoError(t, err)
		assert.Equal(t, wres, got)
	}
}
