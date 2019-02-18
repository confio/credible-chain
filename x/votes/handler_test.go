package votes

import (
	"bytes"
	"context"
	"testing"
	time "time"

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

	_, src := helper.MakeKey()

	rep1 := "SOME1"
	rep2 := "ELSE4"

	// this is for over-writing
	vote1 := mustBuildVote("A", rep1, "HLP", "SW87", 1991, 100, "")
	id1 := vote1.Identitifer
	vote1b := mustBuildVote("B", rep2, "FOO", "SW87", 1991, 100, id1)

	// this is for addition
	vote2 := mustBuildVote("A", rep2, "MRE", "B18", 1980, 500, "")
	id2 := vote2.Identitifer

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
				{
					path:   "/tally",
					data:   []byte("A"),
					bucket: tallyBucket.Bucket,
					wantRes: []orm.Object{
						buildTally("A", 1),
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
						buildTally("A", 1),
						buildTally(rep1, 1),
					},
				},
			},
		},
		"can update vote": {
			actions: []action{
				{
					conditions: []weave.Condition{src},
					msg:        vote1,
				},
				{
					conditions: []weave.Condition{src},
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
					data:   []byte("B"),
					bucket: tallyBucket.Bucket,
					wantRes: []orm.Object{
						buildTally("B", 1),
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
						buildTally("A", 0),
						buildTally("B", 1),
						buildTally(rep2, 1),
						buildTally(rep1, 0),
					},
				},
			},
		},
		"can combine vote": {
			actions: []action{
				{
					conditions: []weave.Condition{src},
					msg:        vote1,
				},
				{
					conditions: []weave.Condition{src},
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
					data:   []byte("A"),
					bucket: tallyBucket.Bucket,
					wantRes: []orm.Object{
						buildTally("A", 2),
					},
				},
				{
					path:   "/tally",
					mod:    "prefix",
					bucket: tallyBucket.Bucket,
					wantRes: []orm.Object{
						buildTally("A", 2),
						buildTally(rep2, 1),
						buildTally(rep1, 1),
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
