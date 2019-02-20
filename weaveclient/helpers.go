package client

import (
	"encoding/json"
	"time"

	"github.com/tendermint/tendermint/rpc/client"
)

// ToString is a generic stringer which outputs
// a struct in its equivalent (indented) json representation
func ToString(d interface{}) string {
	s, err := json.MarshalIndent(d, "", "	")
	if err != nil {
		return err.Error()
	}
	return string(s)
}

// blocks go by fast, no need to wait seconds....
func FastWaiter(delta int64) (abort error) {
	delay := time.Duration(delta) * 5 * time.Millisecond
	time.Sleep(delay)
	return nil
}

var _ client.Waiter = FastWaiter
