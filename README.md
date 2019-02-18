# Credible Chain

This is a blockchain to store an tally all the votes from
the [credible choice](https://github.com/jpincas/credible-choice) project.
The main point is to provide an immutible audit trail of all events,
to provide more confidence in the veracity of the election results.

## Overview

We run the blockchain using Tendermint in PoA mode. That is we will hardcode
likely 4 validators at the genesis of the chain, which will run the validating nodes.
These nodes should be run by independent entities to prevent collusion, but more
importantly, they provide a signed record that can be synced and verified by anyone
running a full-node. We will provide details on how to connect an "audit" node.

The state machine has one operation - set vote, which will add or update the vote
for that particular person. We use a unique id derived from the mobile phone number
provided from the gateway to prevent double-voting. If we get a second request with
the same id, we change the previous vote, rather than adding a new one.

We allow a fixed set of keys (stored on the api server) to sign a vote transaction
as a witness. We will also include the sms code, as well as a unique transaction id
from the sms gateway in order to allow people to spot-check transactions and verify
that everything is being done legitimately. This is the main purpose of the blockchain -
to provide a fully transparent, public ledger.

## Message Design

We want to embed the sms message into a few human readable words representing the following data:

```golang
type Vote struct {
    MainVote int8 // 0-3 for the main votes
    RepVote int16 // 0-1023 for the rep
    Age int8 // 0-127 for age
    PostCode int16 // 0-3000 for postcode -> use lookup table
    Donation int16 // how many pence is being donated
}
```

Except the donation, this is 2 + 10 + 7 + 12 = 31 bits... which we fit into 3 words of bip39, and the Donation
must be appended as a raw amount on the end for the SMS gateway, so we get something like:
`Hungry fox rebellion 10` to donate 10 pence to a charity and cast a vote that is encoded as `Hungry fox rebellion`

We will provide an encode/decode function from words to vote both in js (for the client) and in go (to write to the db).
We can also provide an online tool to allow people to ensure their code matches their choices.

Note that this requires a lookup table for Reps and PostCodes and MainVote.
This must be stored in the API database, ideally also in the blockchain.

## Blockchain Design

We build the application using the [weave framework](github.com/iov-one/weave),
using their signature and multisig modules, and adding a custom vote extension
for our needs. The data we store is:

```golang
type VoteRecord struct {
    Vote Vote // the important info, see above...
    Identifier string // unique anonymized id from sms gateway
    SMSCode string // the human words eg. "hungry horse fly" that were texted (for auditing)
    TransactionID string // a unique id from the sms gateway to look up this exact message was sent
    VotedAt time.Time // when the vote was cast  
}
```

We define a set of notaries in the genesis file, and any submitted VoteRecord must be signed by at least one of them.
If the Identifier was not present, add this record and add the votes to the tally.
If the Identifer was already present, subtract the votes in previous record and add the new votes to the tally, then update the record.
All history data is stored in the transactions as well, we just need to maintain the last vote and tally in the blockchain db.

## APIs

We must provide a client API in golang to encode, sign, and submit transactions. This will be exported by this
repository with the goal of being imported by the API server, to be used with a private key stored there.

If desired, we can extend iov-core to provide a TypeScript REPL allowing exploration and ad-hoc queries of
the current blockchain state (as well as transaction history). Or any other solution to audit that is desired.
