package queue

import (
	"time"

	"github.com/confio/credible-chain/client"
	wc "github.com/confio/credible-chain/weaveclient"

	"github.com/iov-one/weave/crypto"
)

const NumConcurrentRequests = 10
const CommitTimeout = 60 * time.Second

type Pipeline struct {
	client        *client.CredibleClient
	key           *crypto.PrivateKey
	chainID       string
	nonce         *wc.Nonce
	toSend        chan *Task
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
		toSend:        make(chan *Task, 1),
		awaitResponse: make(chan *Task, NumConcurrentRequests),
	}
	return p, nil
}

// Run will initialize a number of go routines then returns the output channel
// only way to stop is to close the input channel
func (p *Pipeline) Run(tasks <-chan *Task) <-chan *Task {
	out := make(chan *Task, 1)

	go signTx(p.key, p.nonce, p.chainID, tasks, p.toSend)
	go sendTx(p.client, p.nonce, p.toSend, p.awaitResponse)
	go getResponse(p.awaitResponse, out)

	return out
}

func signTx(key *crypto.PrivateKey, nonce *wc.Nonce, chainID string, in <-chan *Task, out chan<- *Task) {
	for {
		task, more := <-in
		// end when input is closed
		if !more {
			close(out)
			return
		}
		// push errors out
		task = doSignTx(key, nonce, chainID, task)
		out <- task
	}
}

func doSignTx(key *crypto.PrivateKey, nonce *wc.Nonce, chainID string, task *Task) *Task {
	vr := task.Vote
	tx, err := client.BuildVoteTx(vr.Identifier, vr.SmsCode, vr.TransactionId, vr.Vote)
	if err != nil {
		return task.WithError(err)
	}
	n, err := nonce.Next()
	if err != nil {
		return task.WithError(err)
	}
	client.SignTx(tx, key, chainID, n)
	task.Tx = tx
	return task
}

func sendTx(client *client.CredibleClient, nonce *wc.Nonce, in <-chan *Task, out chan<- *Task) {
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
		task.Response = make(chan wc.BroadcastTxResponse)
		err := client.BroadcastAsyncWithCheck(task.Tx, CommitTimeout, task.Response)
		if err != nil {
			task.Error = err
			// this means we didn't update the nonce, force a reset from the chain
			// TODO: think about how singing and sending are separated, maybe combine them???
			nonce.ClearCache()
		}
		out <- task
	}
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
