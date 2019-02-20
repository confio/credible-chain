package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/rpc/client"
	rpctest "github.com/tendermint/tendermint/rpc/test"
)

// blocks go by fast, no need to wait seconds....
func fastWaiter(delta int64) (abort error) {
	delay := time.Duration(delta) * 5 * time.Millisecond
	time.Sleep(delay)
	return nil
}

var _ client.Waiter = fastWaiter

func TestMainSetup(t *testing.T) {
	config := rpctest.GetConfig()
	assert.Equal(t, "SetInTestMain", config.Moniker)

	conn := client.NewLocal(node)
	status, err := conn.Status()
	require.NoError(t, err)
	assert.Equal(t, "SetInTestMain", status.NodeInfo.Moniker)

	// wait for some blocks to be produced....
	client.WaitForHeight(conn, 5, fastWaiter)
	status, err = conn.Status()
	require.NoError(t, err)
	assert.True(t, status.SyncInfo.LatestBlockHeight > 4)
}

func TestNonce(t *testing.T) {
	addr := GenPrivateKey().PublicKey().Address()
	conn := NewLocalConnection(node)
	bcp := NewClient(conn)

	nonce := NewNonce(bcp, addr)
	n, err := nonce.Next()
	require.NoError(t, err)
	assert.Equal(t, int64(0), n)

	n, err = nonce.Next()
	require.NoError(t, err)
	assert.Equal(t, int64(1), n)

	n, err = nonce.Next()
	require.NoError(t, err)
	assert.Equal(t, int64(2), n)

	n, err = nonce.Query()
	require.NoError(t, err)
	assert.Equal(t, int64(0), n)
}

func TestSubscribeHeaders(t *testing.T) {
	conn := NewLocalConnection(node)
	bcp := NewClient(conn)

	headers := make(chan *Header, 4)
	cancel, err := bcp.SubscribeHeaders(headers)
	require.NoError(t, err)

	// get two headers and cancel
	h := <-headers
	h2 := <-headers
	cancel()

	assert.NotNil(t, h)
	assert.NotNil(t, h2)
	assert.NotEmpty(t, h.ChainID)
	assert.NotEmpty(t, h.Height)
	assert.Equal(t, h.ChainID, h2.ChainID)
	assert.Equal(t, h.Height+1, h2.Height)

	// nothing else should be produced, let's wait 100ms to be sure
	timer := time.After(100 * time.Millisecond)
	select {
	case evt := <-headers:
		require.Nil(t, evt, "This must be nil from a closed channel")
	case <-timer:
		// we want this to fire
	}
}

// func TestSendMultipleTx(t *testing.T) {
// 	conn := NewLocalConnection(node)
// 	bcp := NewClient(conn)

// 	friend := GenPrivateKey()
// 	rcpt := friend.PublicKey().Address()
// 	src := faucet.PublicKey().Address()

// 	nonce := NewNonce(bcp, src)
// 	chainID, err := bcp.ChainID()
// 	amount := x.Coin{Whole: 1000, Ticker: initBalance.Ticker}
// 	require.NoError(t, err)

// 	// a prep transaction, so the recipient has something to send
// 	prep := BuildSendTx(src, rcpt, amount, "Send 1")
// 	n, err := nonce.Next()
// 	require.NoError(t, err)
// 	SignTx(prep, faucet, chainID, n)

// 	// from sender with a different nonce
// 	tx := BuildSendTx(src, rcpt, amount, "Send 2")
// 	n, err = nonce.Next()
// 	require.NoError(t, err)
// 	SignTx(tx, faucet, chainID, n)

// 	// and a third one to return from rcpt to sender
// 	// nonce must be 0
// 	tx2 := BuildSendTx(rcpt, src, amount, "Return")
// 	SignTx(tx2, friend, chainID, 0)

// 	// first, we send the one transaction so the next two will succeed
// 	prepResp := bcp.BroadcastTx(prep)
// 	require.NoError(t, prepResp.IsError())
// 	prepH := prepResp.Response.Height

// 	txResp := make(chan BroadcastTxResponse, 2)
// 	headers := make(chan interface{}, 1)
// 	cancel, err := bcp.Subscribe(QueryNewBlockHeader, headers)
// 	require.NoError(t, err)

// 	// to avoid race conditions, wait for a new header
// 	// event, then immediately send off the two tx
// 	var ready, start sync.WaitGroup
// 	ready.Add(2)
// 	start.Add(1)

// 	go func() {
// 		ready.Done()
// 		start.Wait()
// 		bcp.BroadcastTxAsync(tx, txResp)
// 	}()
// 	go func() {
// 		ready.Done()
// 		start.Wait()
// 		bcp.BroadcastTxAsync(tx2, txResp)
// 	}()

// 	ready.Wait()
// 	<-headers
// 	start.Done()
// 	cancel()

// 	// both succeed
// 	resp := <-txResp
// 	resp2 := <-txResp
// 	require.NoError(t, resp.IsError())
// 	require.NoError(t, resp2.IsError())
// 	assert.True(t, resp.Response.Height > prepH+1)
// 	assert.True(t, resp2.Response.Height > prepH+1)
// }
