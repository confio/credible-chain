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
	if len(m.Identitifer) < 5 || len(m.Identitifer) > 100 {
		return errors.ErrInternal("Unexpected size for Identifier")
	}
	if len(m.SmsCode) < 5 || len(m.SmsCode) > 100 {
		return errors.ErrInternal("Unexpected size for sms code")
	}
	if len(m.TransactionId) < 5 || len(m.TransactionId) > 100 {
		return errors.ErrInternal("Unexpected size for transaction id")
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
	if len(m.MainVote) != 1 || !isAlpha.MatchString(m.MainVote) {
		return errors.ErrInternal("MainVote must be 1 character")
	}
	if len(m.RepVote) != 5 || !isAlpha.MatchString(m.RepVote[:4]) || !isNumeric.MatchString(m.RepVote[4:]) {
		return errors.ErrInternal(fmt.Sprintf("RepVote must be 4 letters (%s) and 1 digit (%s)", m.RepVote[:4], m.RepVote[4:]))
	}
	if len(m.Charity) != 3 || !isAlpha.MatchString(m.Charity) {
		return errors.ErrInternal("Charity must be 3 letters")
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
