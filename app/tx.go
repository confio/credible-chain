package credchain

import (
	"fmt"
	"reflect"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/x/multisig"
	"github.com/iov-one/weave/x/sigs"
)

//-------------------------------
// copied from weave/app verbatim
//
// any cleaner way to extend a tx with functionality?

// TxDecoder creates a Tx and unmarshals bytes into it
func TxDecoder(bz []byte) (weave.Tx, error) {
	tx := new(Tx)
	err := tx.Unmarshal(bz)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// make sure tx fulfills all interfaces
var _ weave.Tx = (*Tx)(nil)
var _ sigs.SignedTx = (*Tx)(nil)
var _ multisig.MultiSigTx = (*Tx)(nil)

// GetMsg switches over all types defined in the protobuf file
func (tx *Tx) GetMsg() (weave.Msg, error) {
	return ExtractMsgFromSum(tx.GetSum())
}

// GetSignBytes returns the bytes to sign...
func (tx *Tx) GetSignBytes() ([]byte, error) {
	// temporarily unset the signatures, as the sign bytes
	// should only come from the data itself, not previous signatures
	sigs := tx.Signatures
	tx.Signatures = nil

	bz, err := tx.Marshal()

	// reset the signatures after calculating the bytes
	tx.Signatures = sigs
	return bz, err
}

// ExtractMsgFromSum will find a weave message from a tx sum type if it exists.
// Assuming you define your Tx with protobuf, this will help you implement GetMsg()
//
//   ExtractMsgFromSum(tx.GetSum())
//
// To work, this requires sum to be a pointer to a struct with one field,
// and that field can be cast to a Msg.
// Returns an error if it cannot succeed.
func ExtractMsgFromSum(sum interface{}) (weave.Msg, error) {
	// TODO: add better error messages here with new refactor
	if sum == nil {
		return nil, errors.ErrInternal("message container is <nil>")
	}
	pval := reflect.ValueOf(sum)
	if pval.Kind() != reflect.Ptr || pval.Elem().Kind() != reflect.Struct {
		return nil, errors.ErrUnknownTxType(sum)
	}
	val := pval.Elem()
	if val.NumField() != 1 {
		return nil, errors.ErrInternal(fmt.Sprintf("Unexpected message container field count: %d", val.NumField()))
	}
	field := val.Field(0)
	if field.IsNil() {
		return nil, errors.ErrInternal("message is <nil>")
	}
	res, ok := field.Interface().(weave.Msg)
	if !ok {
		return nil, errors.ErrUnknownTxType(field.Interface())
	}
	return res, nil
}
