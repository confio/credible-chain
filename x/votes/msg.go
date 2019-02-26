package votes

import (
	fmt "fmt"
	"regexp"
	time "time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
)

var _ weave.Msg = (*VoteRecord)(nil)

const (
	pathRecordVote = "votes/record"
)

var (
	isAlpha        = regexp.MustCompile("^[a-zA-Z]+$")
	isNumeric      = regexp.MustCompile("^[0-9]+$")
	isAlphaNumeric = regexp.MustCompile("^[a-zA-Z0-9]+$")
)

// Path is used for Handler routing
func (VoteRecord) Path() string { return pathRecordVote }

// Validate enforces desired rules for content of the VoteRecord
func (m *VoteRecord) Validate() error {
	if m == nil {
		return errors.ErrInternal("VoteRecord is <nil>")
	}
	if len(m.Identitifer) == 0 {
		return errors.ErrInternal("Missing Identifier")
	}
	if len(m.Identitifer) > 100 {
		return errors.ErrInternal("Identifier too long")
	}
	if len(m.SmsCode) == 0 {
		return errors.ErrInternal("Missing sms code")
	}
	if len(m.SmsCode) > 100 {
		return errors.ErrInternal("sms code to long")
	}
	if len(m.TransactionId) == 0 {
		return errors.ErrInternal("Missing transaction id")
	}
	if len(m.TransactionId) > 100 {
		return errors.ErrInternal("transaction id to long")
	}
	if m.VotedAt == nil || m.VotedAt.Before(time.Date(2019, 2, 15, 12, 0, 0, 0, time.UTC)) {
		return errors.ErrInternal("Need reasonable voted_at time")
	}
	return m.Vote.Validate()
}

// Validate enforces desired rules for content of the Vote
func (m *Vote) Validate() error {
	if m == nil {
		return errors.ErrInternal("Vote is <nil>")
	}
	if m.MainVote != 1 && m.MainVote != 2 && m.MainVote != 3 {
		return errors.ErrInternal("MainVote must be 1 character")
	}
	if len(m.RepVote) != 3 || !isAlphaNumeric.MatchString(m.RepVote) {
		return errors.ErrInternal(fmt.Sprintf("RepVote must be 3 characters: %s", m.RepVote))
	}
	if len(m.Charity) != 2 || !isAlphaNumeric.MatchString(m.Charity) {
		return errors.ErrInternal("Charity must be 2 characters")
	}
	if len(m.PostCode) > 4 || len(m.PostCode) < 2 || !isAlphaNumeric.MatchString(m.PostCode) {
		return errors.ErrInternal("Post Code must be 2-4 characters")
	}
	if m.BirthYear < 1900 || m.BirthYear > 2002 {
		return errors.ErrInternal("Must include full birth year, between 1900 and 2002")
	}
	if m.Donation < 0 || m.Donation > 10000 {
		return errors.ErrInternal("Donation must be between 0 and 100 pounds")
	}
	return nil
}

// EncodeToSms will produce the code for the sms to send...
func (m *Vote) EncodeToSms() string {
	return "SMScodeNOTimplemented"
}
