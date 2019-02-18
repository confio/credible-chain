package votes

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/x"
	"github.com/iov-one/weave/x/cash"
)

const (
	recordVoteCost int64 = 50
)

// RegisterQuery registers payment channel bucket under /paychans.
func RegisterQuery(qr weave.QueryRouter) {
	NewVoteBucket().Register("vote", qr)
	NewTallyBucket().Register("tally", qr)
}

// RegisterRouters registers payment channel message handelers in given registry.
func RegisterRoutes(r weave.Registry, auth x.Authenticator, cash cash.Controller) {
	bucket := NewVoteBucket()
	tallies := NewTallyBucket()
	r.Handle(pathRecordVote, &recordVoteHandler{auth: auth, bucket: bucket, tallies: tallies})
}

type recordVoteHandler struct {
	auth    x.Authenticator
	bucket  VoteBucket
	tallies TallyBucket
}

var _ weave.Handler = (*recordVoteHandler)(nil)

func (h *recordVoteHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (weave.CheckResult, error) {
	var res weave.CheckResult
	if _, err := h.validate(ctx, db, tx); err != nil {
		return res, err
	}

	res.GasAllocated += recordVoteCost
	return res, nil
}

func (h *recordVoteHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*VoteRecord, error) {
	rmsg, err := tx.GetMsg()
	if err != nil {
		return nil, err
	}
	msg, ok := rmsg.(*VoteRecord)
	if !ok {
		return nil, errors.ErrUnknownTxType(rmsg)
	}

	if err := msg.Validate(); err != nil {
		return msg, err
	}

	// TODO: only allow authorized addresses (from genesis/init) to write
	// if !h.auth.HasAddress(ctx, msg.Src) {
	// 	return msg, errors.UnauthorizedErr.New("invalid address")
	// }

	return msg, nil
}

func (h *recordVoteHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (weave.DeliverResult, error) {
	var res weave.DeliverResult
	msg, err := h.validate(ctx, db, tx)
	if err != nil {
		return res, err
	}

	// existing is non-nil if we voted before... to properly update talies
	existing, err := h.bucket.GetVote(db, msg.Identitifer)
	if err != nil {
		return res, err
	}

	// let's update the tallies
	if existing != nil {
		err = h.subtractTally(db, existing.Vote.MainVote)
		if err != nil {
			return res, err
		}
		err = h.subtractTally(db, existing.Vote.RepVote)
		if err != nil {
			return res, err
		}
	}
	err = h.addTally(db, msg.Vote.MainVote)
	if err != nil {
		return res, err
	}
	err = h.addTally(db, msg.Vote.RepVote)
	if err != nil {
		return res, err
	}

	// We just overwrite the original data if there was any...
	// Previous value was only needed for updating tallies
	_, err = h.bucket.Create(db, msg)
	if err != nil {
		return res, err
	}

	// update tallies
	return res, nil
}

func (h *recordVoteHandler) addTally(db weave.KVStore, option string) error {
	tally, err := h.tallies.GetTally(db, option)
	if err != nil {
		return err
	}
	if tally == nil {
		tally = &Tally{
			Option: option,
			Total:  0,
		}
	}
	tally.Total++
	_, err = h.tallies.Create(db, tally)
	return err
}

func (h *recordVoteHandler) subtractTally(db weave.KVStore, option string) error {
	tally, err := h.tallies.GetTally(db, option)
	if err != nil {
		return err
	}
	if tally == nil {
		return errors.ErrInternal("Cannot subtract from non-existent tally")
	}
	tally.Total--
	_, err = h.tallies.Create(db, tally)
	return err
}
