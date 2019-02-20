package client

import (
	"time"

	app "github.com/confio/credible-chain/app"
	"github.com/confio/credible-chain/x/votes"
)

// BuildVoteTx will create an unsigned tx to place a vote
func BuildVoteTx(identifier, smsCode, transactionID string, vote *votes.Vote) *app.Tx {
	stamp := time.Now().UTC()
	return &app.Tx{
		Sum: &app.Tx_RecordVoteMsg{&votes.VoteRecord{
			Vote:          vote,
			Identitifer:   identifier,
			SmsCode:       smsCode,
			TransactionId: transactionID,
			VotedAt:       &stamp,
		}},
	}
}

type Tally = votes.Tally

func (b *BnsClient) GetTally(option string) (Tally, error) {
	var out Tally

	resp, err := b.AbciQuery("/tally", []byte(option))
	if err != nil {
		return out, err
	}
	if len(resp.Models) == 0 { // empty list or nil
		return out, nil
	}
	// assume only one result
	model := resp.Models[0]
	err = out.Unmarshal(model.Value)
	return out, err
}

func (b *BnsClient) GetAllTallies(option string) ([]Tally, error) {
	resp, err := b.AbciQuery("/tally?prefix", nil)
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

func (b *BnsClient) GetVote(identifier string) (*VoteRecord, error) {
	resp, err := b.AbciQuery("/vote", []byte(identifier))
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
// func (b *BnsClient) GetWallet(addr weave.Address) (*WalletResponse, error) {
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
