package queue

import (
	"time"

	"github.com/iov-one/weave/crypto"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/confio/credible-chain/client"
	wc "github.com/confio/credible-chain/weaveclient"
)

const NumConcurrentRequests = 10
const CommitTimeout = 60 * time.Second

type Pipeline struct {
	client        *client.CredibleClient
	key           *crypto.PrivateKey
	chainID       string
	nonce         *wc.Nonce
	awaitResponse chan *Task
}

func NewPipeline(client *client.CredibleClient, key *crypto.PrivateKey) (*Pipeline, error) {
	chainID, err := client.ChainID()
	if err != nil {
		return nil, err
	}
	nonce := wc.NewNonce(client, key.PublicKey().Address())
	p := &Pipeline{
		client:        client,
		key:           key,
		chainID:       chainID,
		nonce:         nonce,
		awaitResponse: make(chan *Task, NumConcurrentRequests),
	}
	return p, nil
}

// Run will initialize a number of go routines then returns the output channel
// only way to stop is to close the input channel
func (p *Pipeline) Run(tasks <-chan *Task) <-chan *Task {
	out := make(chan *Task, 1)

	go signTx(p.client, p.key, p.nonce, p.chainID, tasks, p.awaitResponse)
	go getResponse(p.awaitResponse, out)

	return out
}

func signTx(cc *client.CredibleClient, key *crypto.PrivateKey, nonce *wc.Nonce, chainID string, in <-chan *Task, out chan<- *Task) {
	for {
		task, more := <-in
		// end when input is closed
		if !more {
			close(out)
			return
		}
		// error in, error out
		if task.Error != nil {
			out <- task
			continue
		}
		// process
		task = doSignAndSendTx(cc, key, nonce, chainID, task)
		out <- task
	}
}

func doSignAndSendTx(cc *client.CredibleClient, key *crypto.PrivateKey, nonce *wc.Nonce, chainID string, task *Task) *Task {
	vr := task.Vote
	tx, err := client.BuildVoteTx(vr.Identifier, vr.SmsCode, vr.TransactionId, vr.Vote)
	if err != nil {
		return task.WithError(err)
	}
	n, err := nonce.Next()
	if err != nil {
		return task.WithError(err)
	}
	err = client.SignTx(tx, key, chainID, n)
	if err != nil {
		return task.WithError(err)
	}
	task.Tx = tx
	task.Response = make(chan wc.BroadcastTxResponse)
	err = cc.BroadcastAsyncWithCheck(task.Tx, CommitTimeout, task.Response)
	if err != nil {
		// this means we didn't update the nonce, force a reset from the chain
		nonce.ClearCache()
		// wait for next block
		waitForNextBlock(cc)
		return task.WithError(err)
	}
	return task
}

func getResponse(in <-chan *Task, out chan<- *Task) {
	for {
		task, more := <-in
		// end when input is closed
		if !more {
			close(out)
			return
		}
		// error in, error out
		if task.Error != nil {
			out <- task
			continue
		}
		// we see if we get an error or success on the commit
		result := <-task.Response
		close(task.Response)
		if result.Error != nil {
			task.Error = result.Error
		}
		// no more info needed for success
		out <- task
	}
}

// waitForNextBlock will wait until a new header comes in, so we can get proper nonce queries afterwards
func waitForNextBlock(cc *client.CredibleClient) error {
	// subscribe to block headers
	headers := make(chan *tmtypes.Header, 2)
	cancel, err := cc.SubscribeHeaders(headers)
	if err != nil {
		return err
	}
	// get one header and cancel
	<-headers
	cancel()
	return nil
}
