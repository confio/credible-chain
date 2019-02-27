package client

import (
	"time"

	app "github.com/confio/credible-chain/app"
	"github.com/confio/credible-chain/x/votes"
)

// BuildVoteTx will create an unsigned tx to place a vote
func BuildVoteTx(identifier, smsCode, transactionID string, vote *votes.Vote) (*app.Tx, error) {
	stamp := time.Now().UTC()
	msg := &votes.VoteRecord{
		Vote:          vote,
		Identifier:   identifier,
		SmsCode:       smsCode,
		TransactionId: transactionID,
		VotedAt:       &stamp,
	}
	if err := msg.Validate(); err != nil {
		return nil, err
	}
	return &app.Tx{Sum: &app.Tx_RecordVoteMsg{msg}}, nil
}

type Tally = votes.Tally

func (c *CredibleClient) GetTally(option string) (*Tally, error) {
	resp, err := c.AbciQuery("/tally", []byte(option))
	if err != nil {
		return nil, err
	}
	if len(resp.Models) == 0 { // empty list or nil
		return nil, nil
	}

	// assume only one result
	var out Tally
	err = out.Unmarshal(resp.Models[0].Value)
	return &out, err
}

func (c *CredibleClient) GetAllTallies() ([]Tally, error) {
	resp, err := c.AbciQuery("/tally?prefix", nil)
	if err != nil {
		return nil, err
	}

	var out = make([]Tally, len(resp.Models))
	for i := range out {
		err = out[i].Unmarshal(resp.Models[i].Value)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

type VoteRecord = votes.VoteRecord

func (c *CredibleClient) GetVote(identifier string) (*VoteRecord, error) {
	resp, err := c.AbciQuery("/vote", []byte(identifier))
	if err != nil {
		return nil, err
	}
	if len(resp.Models) == 0 { // empty list or nil
		return nil, nil
	}

	// assume only one result
	var out VoteRecord
	err = out.Unmarshal(resp.Models[0].Value)
	return &out, err
}

// // GetWallet will return a wallet given an address
// // If non wallet is present, it will return (nil, nil)
// // Error codes are used when the query failed on the server
// func (b *CredibleClient) GetWallet(addr weave.Address) (*WalletResponse, error) {
// 	// make sure we send a valid address to the server
// 	err := addr.Validate()
// 	if err != nil {
// 		return nil, errors.WithMessage(err, "Invalid Address")
// 	}

// 	resp, err := b.AbciQuery("/wallets", addr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(resp.Models) == 0 { // empty list or nil
// 		return nil, nil
// 	}
// 	// assume only one result
// 	model := resp.Models[0]
// 	// make sure the return value is expected
// 	acct := walletKeyToAddr(model.Key)
// 	if !addr.Equals(acct) {
// 		return nil, errors.Errorf("Mismatch. Queried %s, returned %s", addr, acct)
// 	}
// 	out := WalletResponse{
// 		Address: acct,
// 		Height:  resp.Height,
// 	}

// 	// parse the value as wallet bytes
// 	err = out.Wallet.Unmarshal(model.Value)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &out, nil
// }
