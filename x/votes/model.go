package votes

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
)

// We persist the same vote data to the store, as we have in the transaction
// Maybe we could optimize later, but seems good for auditing
var _ orm.CloneableData = (*VoteRecord)(nil)
var _ orm.CloneableData = (*Tally)(nil)

// Copy returns a shallow copy of this VoteRecord.
// TODO: improve
func (v *VoteRecord) Copy() orm.CloneableData {
	vote := v.Vote
	if vote != nil {
		vote = vote.Copy().(*Vote)
	}
	return &VoteRecord{
		Vote:       vote,
		Identifier: v.Identifier,
		SmsCode:    v.SmsCode,
		VotedAt:    v.VotedAt,
	}
}

// Copy returns a shallow copy of this Vote.
func (v Vote) Copy() orm.CloneableData {
	return &v
}

// Copy returns a shallow copy of this Tally.
func (t Tally) Copy() orm.CloneableData {
	return &t
}

// Validate ensures the payment channel is valid.
func (t *Tally) Validate() error {
	if t == nil {
		return errors.ErrInternal("Tally is <nil>")
	}
	if len(t.Option) < 1 || !isAlphaNumeric.MatchString(t.Option) {
		return errors.ErrInternal("Option must be an alphanumeric string")
	}
	if t.Total < 0 {
		return errors.ErrInternal("Tally.total must be positive")
	}
	return nil
}

type VoteBucket struct {
	orm.Bucket
}

// NewVoteBucket returns a bucket for storing Vote state.
func NewVoteBucket() VoteBucket {
	b := orm.NewBucket("votes", orm.NewSimpleObj(nil, &VoteRecord{Vote: &Vote{}}))
	return VoteBucket{
		Bucket: b,
	}
}

func (b *VoteBucket) Create(db weave.KVStore, v *VoteRecord) (orm.Object, error) {
	key := []byte(v.Identifier)
	obj := orm.NewSimpleObj(key, v)
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj, b.Bucket.Save(db, obj)
}

// Save updates the state of given Vote entity in the store.
func (b *VoteBucket) Save(db weave.KVStore, obj orm.Object) error {
	if _, ok := obj.Value().(*VoteRecord); !ok {
		return orm.ErrInvalidObject(obj.Value())
	}
	if err := obj.Validate(); err != nil {
		return err
	}
	return b.Bucket.Save(db, obj)
}

// GetVote returns a vote with this key/identifier or returns nil
func (b *VoteBucket) GetVote(db weave.KVStore, identifier string) (*VoteRecord, error) {
	obj, err := b.Get(db, []byte(identifier))
	if err != nil {
		return nil, err
	}
	if obj == nil || obj.Value() == nil {
		return nil, nil
	}
	v, ok := obj.Value().(*VoteRecord)
	if !ok {
		return nil, orm.ErrInvalidObject(obj.Value())
	}
	return v, nil
}

type TallyBucket struct {
	orm.Bucket
}

// NewTallyBucket returns a bucket for storing Tally state.
func NewTallyBucket() TallyBucket {
	b := orm.NewBucket("tally", orm.NewSimpleObj(nil, &Tally{}))
	return TallyBucket{
		Bucket: b,
	}
}

func (b *TallyBucket) Create(db weave.KVStore, v *Tally) (orm.Object, error) {
	key := []byte(v.Option)
	obj := orm.NewSimpleObj(key, v)
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj, b.Bucket.Save(db, obj)
}

// Save updates the state of given Tally entity in the store.
func (b *TallyBucket) Save(db weave.KVStore, obj orm.Object) error {
	if _, ok := obj.Value().(*Tally); !ok {
		return orm.ErrInvalidObject(obj.Value())
	}
	if err := obj.Validate(); err != nil {
		return err
	}
	return b.Bucket.Save(db, obj)
}

// GetTally returns a Tally with this key/identifier or returns nil
func (b *TallyBucket) GetTally(db weave.KVStore, identifier string) (*Tally, error) {
	obj, err := b.Get(db, []byte(identifier))
	if err != nil {
		return nil, err
	}
	if obj == nil || obj.Value() == nil {
		return nil, nil
	}
	v, ok := obj.Value().(*Tally)
	if !ok {
		return nil, orm.ErrInvalidObject(obj.Value())
	}
	return v, nil
}
